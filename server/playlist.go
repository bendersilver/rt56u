package server

import (
	"bufio"
	"encoding/gob"
	"net/url"
	"os"
	"path"
	"regexp"
	"sort"
	"strings"
	"time"

	"github.com/bendersilver/blog"
)

const uriPlst = "http://ott.tv.planeta.tc/playlist/channels.m3u?4k&groupChannels=thematic&fields=epg,group"

var reID = regexp.MustCompile(`tvg-id="([^"]+)"`)
var reName = regexp.MustCompile(`tvg-name="([^"]+)"`)
var reLogo = regexp.MustCompile(`tvg-logo="([^"]+)"`)
var reGroup = regexp.MustCompile(`group-title="([^"]+)"`)
var reRemove = regexp.MustCompile(`tvg-url="([^"]+)" `)

var gobFile string

func init() {
	if f, ok := os.LookupEnv("ENVGOB"); !ok {
		blog.Fatal("set ENVGOB")
	} else {
		gobFile = f
	}
	go updatePlast()
	go func() {
		for range time.Tick(time.Hour * 6) {
			updatePlast()
		}
	}()
}

// M3UItem -
type M3UItem struct {
	ID     string   `json:"id"`
	Name   string   `json:"name"`
	URL    *url.URL `json:"-"`
	Logo   string   `json:"img"`
	Group  string   `json:"group"`
	Order  int      `json:"order"`
	Del    bool     `json:"del"`
	Hide   bool     `json:"hide"`
	Extinf string   `json:"-"`
}

func (m *M3UItem) setFields() {
	m.find(&m.Name, reName)
	m.find(&m.Logo, reLogo)
	m.find(&m.Group, reGroup)
}

func (m *M3UItem) find(f *string, re *regexp.Regexp) {
	sub := re.FindStringSubmatch(m.Extinf)
	if len(sub) == 2 {
		*f = sub[1]
	}
}

// M3U -
type M3U []*M3UItem

func (m M3U) Len() int {
	return len(m)
}

func (m M3U) Less(i, j int) bool {
	return m[i].Order < m[j].Order
}

func (m M3U) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

// Loads -
func (m *M3U) Loads() error {
	file, err := os.OpenFile(gobFile, os.O_RDONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := gob.NewDecoder(file)
	err = encoder.Decode(m)
	if err != nil {
		return err
	}
	sort.Sort(m)
	return nil
}

// Dumps -
func (m *M3U) Dumps() error {
	file, err := os.OpenFile(gobFile, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	sort.Sort(m)
	encoder := gob.NewEncoder(file)
	err = encoder.Encode(m)
	if err != nil {
		return err
	}
	file.Close()

	file, err = os.OpenFile(path.Join(static, "plst.m3u"), os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	file.WriteString("#EXTM3U\n")
	for _, line := range *m {
		if line.Del || line.Hide {
			continue
		}
		file.WriteString(line.Extinf)
	}
	file.Close()
	return nil
}

func updatePlast() {
	rsp, err := cli.Get(uriPlst)
	if err != nil {
		blog.Error(err)
		return
	}
	defer rsp.Body.Close()
	scanner := bufio.NewScanner(rsp.Body)

	plstDump := M3U{}
	plstDump.Loads()
	defer plstDump.Dumps()

	var playList bool
	var line string
	var item *M3UItem

Loop:
	for scanner.Scan() {
		line = strings.TrimSpace(scanner.Text())
		if !playList {
			if strings.HasPrefix(line, "#EXTM3U") {
				playList = true
				continue
			}
			blog.Error("Invalid playlist")
			return
		}
		switch {
		case line == "":
			continue
		case strings.HasPrefix(line, "http") && item != nil:
			item.URL, err = url.Parse(line)
			if err != nil {
				blog.Error(err)
			}
			item.Extinf += "\n" + line + "\n"
			item = nil
		case strings.HasPrefix(line, "#EXTINF:"):
			line = reRemove.ReplaceAllString(line, "")
			sub := reID.FindStringSubmatch(line)
			if len(sub) == 2 {
				for _, i := range plstDump {
					if i.ID == sub[1] {
						item = i
						item.Extinf = line
						item.setFields()
						continue Loop
					}
				}
				item = new(M3UItem)
				item.ID = sub[1]
				item.Order = len(plstDump)
				item.setFields()
				plstDump = append(plstDump, item)
			}
		}
	}
}
