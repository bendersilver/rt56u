package handler

import (
	"encoding/gob"
	"fmt"
	"os"
	"path"
	"sort"
	"time"

	"github.com/bendersilver/simplog"
)

var gobFile, plstFile string

func init() {
	gobFile = path.Join(os.Getenv("GOB"), "plst")
	plstFile = path.Join(os.Getenv("DIST"), "plst.m3u")
	updatePlst()
	go func() {
		for range time.Tick(time.Hour) {
			updatePlst()
		}
	}()
}

// M3UItem -
type M3UItem struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Logo  string `json:"img"`
	Group string `json:"group"`
	URL   string `json:"url"`
	Order int    `json:"order"`
	Del   bool   `json:"del"`
	Hide  bool   `json:"hide"`
}

func (m *M3UItem) str() string {
	return fmt.Sprintf("#EXTINF:-1 tvg-id=\"%s\" tvg-name=\"%s\" tvg-logo=\"%s\" group-title=\"%s\",%s\n%s\n",
		m.ID, m.Name, m.Logo, m.Group, m.Name, m.URL)
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

	file, err = os.OpenFile(plstFile+".tmp", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	file.WriteString("#EXTM3U\n")
	for _, line := range *m {
		if line.Del || line.Hide {
			continue
		}
		file.WriteString(line.str())
	}
	file.Close()
	return os.Rename(plstFile+".tmp", plstFile)
}

func updatePlst() {
	var plst M3U
	err := plst.Loads()
	if err != nil {
		simplog.Error(err)
		return
	}

	ch, err := getAll()
	if err != nil {
		simplog.Error(err)
		return
	}
	if len(ch) < 50 {
		simplog.Error("num channels less than 50")
		return
	}
	// hide all
	for _, item := range plst {
		item.Del = true
	}

	// update exists
Parent:
	for i, c := range ch {
		for _, item := range plst {
			if c.URI == item.URL {
				item.URL = c.URI
				item.Name = c.Name
				item.Logo = c.Tags["tvg-logo"]
				item.Group = c.Tags["group-title"]
				item.Del = false
				ch[i] = nil
				continue Parent
			}
		}
		// add new
		if id, ok := c.Tags["tvg-id"]; ok {
			plst = append(plst, &M3UItem{
				ID:    id,
				URL:   c.URI,
				Name:  c.Name,
				Logo:  c.Tags["tvg-logo"],
				Group: c.Tags["group-title"],
				Order: len(plst) + 1,
			})
		}
		ch[i] = nil
	}
	plst.Dumps()
}
