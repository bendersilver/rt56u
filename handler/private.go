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
func PrivatePost(con *net.TCPConn, p string) {
	var plst M3U
	switch p {
	case "toggle":
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
	case "save":
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
func Private(con *net.TCPConn, p string) {
	var plst M3U
	switch p {
	case "/xml.gz":
		err := transferFull("ott.tv.planeta.tc/epg/program.xml.gz", con)
		if err != nil {
			Err404(con)
			simplog.Error(err)
		}
	case "/jsonAPI/get":
		plst.Loads()
		b, _ := json.Marshal(plst)
		Status200(con)
		con.Write([]byte("Content-Type: application/json; charset=utf-8\r\n\r\n"))
		con.Write(b)
	default:
		f, err := os.OpenFile(path.Join(os.Getenv("DIST"), p), os.O_RDONLY, 0644)
		if err != nil {
			simplog.Error(err)
			Err404(con)
		} else {
			defer f.Close()
			Status200(con)
			con.Write(cType(p))
			io.Copy(con, f)
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
