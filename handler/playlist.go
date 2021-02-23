package handler

import (
	"bufio"
	"encoding/gob"
	"net"
	"os"
	"path"
	"regexp"
	"sort"
	"strings"
	"time"

	"github.com/bendersilver/blog"
)

const uriPlst = "http://ott.tv.planeta.tc/playlist/channels.m3u?4k&groupChannels=thematic&fields=epg,group&hlsQuality=min&hlsVideoOnly"

var gobFile, plstFile string

func init() {
	gobFile = path.Join(os.Getenv("GOB"), "plst")
	plstFile = path.Join(os.Getenv("DIST"), "plst.m3u")
	updatePlst()
	go func() {
		for range time.Tick(time.Hour * 6) {
			updatePlst()
		}
	}()
}

var reID = regexp.MustCompile(`tvg-id="([^"]+)"`)
var reName = regexp.MustCompile(`tvg-name="([^"]+)"`)
var reLogo = regexp.MustCompile(`tvg-logo="([^"]+)"`)
var reGroup = regexp.MustCompile(`group-title="([^"]+)"`)
var reRemove = regexp.MustCompile(`tvg-url="([^"]+)" `)

// M3UItem -
type M3UItem struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Logo   string `json:"img"`
	Group  string `json:"group"`
	Order  int    `json:"order"`
	Del    bool   `json:"del"`
	Hide   bool   `json:"hide"`
	Extinf string `json:"-"`
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

	file, err = os.OpenFile(plstFile, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
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

func updatePlst() {
	dl, err := net.Dial("tcp", "ott.tv.planeta.tc:80")
	if err != nil {
		return
	}
	defer dl.Close()
	_, err = dl.Write([]byte("GET /playlist/channels.m3u?4k&groupChannels=thematic&fields=epg,group&hlsQuality=min&hlsVideoOnly HTTP/1.0\nHosh: ott.tv.planeta.tc\nUser-Agent: go-iptv\n\n"))
	if err != nil {
		return
	}
	_, _, _, err = ParseHeader(dl, 256)
	if err != nil {
		return
	}

	plstDump := M3U{}
	plstDump.Loads()
	defer plstDump.Dumps()

	var line string
	var item *M3UItem

	scanner := bufio.NewScanner(dl)
Loop:
	for scanner.Scan() {
		line = scanner.Text()
		switch {
		case line == "":
			continue
		case strings.HasPrefix(line, "http") && item != nil:
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
