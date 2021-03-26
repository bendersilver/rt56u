package handler

import (
	"encoding/json"
	"io"
	"net"
	"os"
	"path"
	"path/filepath"

	"github.com/bendersilver/simplog"
)

// PrivatePost -
func PrivatePost(con *net.TCPConn, m *MIMEHeader) {
	var plst M3U
	switch m.Path {
	case "/jsonAPI/toggle":
		plst.Loads()
		defer plst.Dumps()
		var m struct {
			ID   string
			Hide bool
		}
		if err := json.NewDecoder(con).Decode(&m); err == nil {
			for _, v := range plst {
				if v.ID == m.ID {
					v.Hide = m.Hide
					break
				}
			}
			Status200(con)
		} else {
			Err404(con)
			simplog.Error(err)
		}
	case "/jsonAPI/save":
		plst.Loads()
		defer plst.Dumps()
		var m []struct {
			ID    string
			Order int
		}
		if err := json.NewDecoder(con).Decode(&m); err == nil {
			for _, v := range plst {
				for _, i := range m {
					if v.ID == i.ID {
						v.Order = i.Order
					}
				}
			}
			Status200(con)
		} else {
			Err404(con)
			simplog.Error(err)
		}
	default:
		Err404(con)
	}
}

// Private -
func Private(w io.Writer, m *MIMEHeader) {
	var plst M3U
	switch m.Path {
	case "/xml.gz":
		err := transferFull("ott.tv.planeta.tc/epg/program.xml.gz", m, w, false)
		if err != nil {
			Err500(w, err.Error())
			simplog.Error(err)
		}
	case "/jsonAPI/get":
		plst.Loads()
		b, _ := json.Marshal(plst)
		Status200(w)
		w.Write([]byte("Content-Type: application/json; charset=utf-8\r\n\r\n"))
		w.Write(b)
	default:
		f, err := os.OpenFile(path.Join(os.Getenv("DIST"), m.Path), os.O_RDONLY, 0644)
		if err != nil {
			Err500(w, err.Error())
		} else {
			defer f.Close()
			Status200(w)
			w.Write(cType(m.Path))
			io.Copy(w, f)
		}
	}
}

func cType(p string) []byte {
	switch filepath.Ext(p) {
	case ".html":
		return []byte("Content-Type: text/html; charset=utf-8\r\n\r\n")
	case ".css":
		return []byte("Content-Type: text/css; charset=utf-8\r\n\r\n")
	case ".js":
		return []byte("Content-Type: text/javascript; charset=utf-8\r\n\r\n")
	case ".ico":
		return []byte("Content-Type: image/x-icon\r\n\r\n")
	case ".map":
		return []byte("Content-Type: application/json; charset=utf-8\r\n\r\n")
	case ".m3u":
		return []byte("Content-Type: application/x-mpegurl; charset=utf-8\r\n\r\n")
	}
	return []byte("\r\n\r\n")
}
