package handler

import (
	"encoding/json"
	"io"
	"net"
	"os"
	"path/filepath"
)

// PrivatePost -
func PrivatePost(con *net.TCPConn, p string) {
	var plst M3U
	switch p {
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
			Err500(con, err.Error())
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
			Err500(con, err.Error())
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
			Err500(con, err.Error())
		}
	case "/jsonAPI/get":
		plst.Loads()
		b, _ := json.Marshal(plst)
		Status200(con)
		con.Write([]byte("Content-Type: application/json; charset=utf-8\r\n\r\n"))
		con.Write(b)
	default:
		item, ok := static[p]
		if ok {
			f, err := os.OpenFile(item.Path, os.O_RDONLY, 0644)
			if err != nil {
				Err500(con, err.Error())
			} else {
				defer f.Close()
				Status200(con)
				con.Write([]byte(item.Type))
				io.Copy(con, f)
			}
		} else {
			Err404(con)
		}
	}
}

type staticFile struct {
	Type string
	Path string
}

var static map[string]*staticFile

// Init -
func Init() {

	static = make(map[string]*staticFile)
	err := filepath.Walk(os.Getenv("DIST"), func(p string, i os.FileInfo, err error) error {
		if !i.IsDir() && err == nil {
			item := new(staticFile)
			item.Path = p

			var path string
			path, err = filepath.Rel(os.Getenv("DIST"), p)
			switch filepath.Ext(p) {
			case ".html":
				item.Type = "Content-Type: text/html; charset=utf-8\r\n\r\n"
			case ".css":
				item.Type = "Content-Type: text/css; charset=utf-8\r\n\r\n"
			case ".js":
				item.Type = "Content-Type: text/javascript; charset=utf-8\r\n\r\n"
			case ".ico":
				item.Type = "Content-Type: image/x-icon\r\n\r\n"
			case ".map":
				item.Type = "Content-Type: application/json; charset=utf-8\r\n\r\n"
			case ".m3u":
				item.Type = "Content-Type: application/x-mpegurl; charset=utf-8\r\n\r\n"
			default:
				return nil
			}
			static["/"+path] = item
		}
		return err
	})
	if err != nil {
		panic(err)
	}
}
