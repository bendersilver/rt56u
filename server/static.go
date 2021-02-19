package server

import (
	"bufio"
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/bendersilver/blog"
)

var static string

func init() {
	if f, ok := os.LookupEnv("DIST"); !ok {
		blog.Fatal("set DIST")
	} else {
		static = f
	}
}

func jsonAPI(w http.ResponseWriter, r *http.Request) {
	var p M3U
	p.Loads()

	if r.URL.Path == "/jsonAPI/get" {
		b, _ := json.Marshal(p)
		w.Write(b)
	} else if strings.HasPrefix(r.URL.Path, "/jsonAPI/save") {
		var list M3U
		json.NewDecoder(r.Body).Decode(&list)
		for _, i := range list {
			for _, v := range p {
				if v.ID == i.ID {
					v.Order = i.Order
				}
			}
		}
		p.Dumps()

	} else if strings.HasPrefix(r.URL.Path, "/jsonAPI/toggle/") {
		d, id := path.Split(r.URL.Path)
		for _, v := range p {
			if v.ID == id {
				v.Hide = path.Base(d) == "true"
				break
			}
		}
		p.Dumps()
	}

}

func fileServ(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path == "/www/" {
		http.ServeFile(w, req, path.Join(static, "index.html"))
	} else {
		http.ServeFile(w, req, path.Join(static, req.URL.Path))
	}
}

func playList(w http.ResponseWriter, req *http.Request) {
	if strings.HasPrefix(req.RemoteAddr, "192.168") {
		http.ServeFile(w, req, path.Join(static, "plst.m3u"))
		return
	}
	file, err := os.OpenFile(path.Join(static, "plst.m3u"), os.O_RDONLY, 0644)
	if err != nil {
		blog.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		defer file.Close()

		reader, writer := io.Pipe()
		defer reader.Close()

		go func() {
			defer writer.Close()
			scanner := bufio.NewScanner(file)
			var line []byte
			for scanner.Scan() {
				line = scanner.Bytes()
				if bytes.HasPrefix(line, []byte("http://")) {
					line = bytes.Replace(line, []byte("playlist.tv.planeta.tc"), []byte(req.Host), 1)
				}
				w.Write(line)
				w.Write([]byte("\n"))
			}
			line = nil
		}()
		w.Header().Set("Content-Type", "application/x-mpegurl; charset=utf-8")
		io.Copy(w, reader)
	}
}
