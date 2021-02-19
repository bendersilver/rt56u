package server

import (
	"encoding/gob"
	"os"
	"sync"
	"time"

	"github.com/bendersilver/blog"
)

var gobBlock string
var mx sync.Mutex

func init() {
	if f, ok := os.LookupEnv("ENVBLOCK"); !ok {
		blog.Fatal("set ENVBLOCK")
	} else {
		gobBlock = f
	}
}

type blockIP struct {
	Count int
	Tme   time.Time
}

type blockList map[string]*blockIP

func newBlockList() blockList {
	b := make(blockList)
	b.Loads()
	return b
}

// Loads -
func (l blockList) unblock() {
	var now time.Time
	for range time.Tick(time.Hour) {
		now = time.Now()
		for k, v := range l {
			if v.Count >= 5 && v.Tme.Add(time.Hour*6).Before(now) {
				mx.Lock()
				delete(l, k)
				blog.Info("unblock", k)
				mx.Unlock()
			}
		}
	}
}

// Loads -
func (l blockList) Loads() error {
	file, err := os.OpenFile(gobBlock, os.O_RDONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := gob.NewDecoder(file)
	err = encoder.Decode(&l)
	if err != nil {
		return err
	}
	return nil
}

// Dumps -
func (l blockList) Dumps() error {
	file, err := os.OpenFile(gobBlock, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	encoder := gob.NewEncoder(file)
	err = encoder.Encode(&l)
	if err != nil {
		return err
	}
	return nil
}

// Check -
func (l blockList) Check(ip string) bool {
	mx.Lock()
	defer mx.Unlock()
	if b, ok := l[ip]; ok {
		return b.Count >= 5
	}
	return false
}

// Dumps -
func (l blockList) Add(ip string) {
	mx.Lock()
	defer mx.Unlock()
	var b *blockIP
	var ok bool
	if b, ok = l[ip]; !ok {
		b = new(blockIP)
		l[ip] = b
	}
	b.Count++
	if b.Count >= 5 {
		b.Tme = time.Now()
		blog.Notice("ip block", ip)
		err := l.Dumps()
		if err != nil {
			blog.Error(err)
		}
	}
}
