package handler

import (
	"io"
	"net/http"
	"os"
	"strings"
)

func changeBody(w io.Writer, r io.Reader, u string) {
	var err error
	var buf = make([]byte, 512)
	var i int
	uri := []byte(u)
	for {
		if _, err = r.Read(buf[i : i+1]); err != nil {
			w.Write(buf[:i])
			break
		}
		// line startwish http://
		//                     ^^
		if i == 6 && buf[5] == '/' && buf[6] == '/' {

			for ix, b := range uri {
				buf[ix] = b
				i = ix
			}
		}
		if buf[i] == '\n' {
			w.Write(buf[:i+1])
			i = -1
		}
		i++
	}
	uri = nil
	buf = nil
}

// transfer -
func transfer(p, h string, w http.ResponseWriter) error {
	res, err := cli.Get("http://" + p)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	for k, v := range res.Header {
		if len(v) > 0 && k != "Content-Length" {
			w.Header().Add(k, v[0])
		}
	}
	if strings.HasPrefix(p, "playlist.tv.planeta.tc/") {
		changeBody(w, res.Body, "http://"+h+"/"+os.Getenv("TOKEN")+"/")
	} else {
		io.Copy(w, res.Body)
	}
	return nil
}

func m3u(w http.ResponseWriter, host string) error {
	f, err := os.OpenFile(plstFile, os.O_RDONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	w.Header().Add("Content-Type", "application/x-mpegurl; charset=utf-8")
	changeBody(w, f, "http://"+host+"/"+os.Getenv("TOKEN")+"/")
	return nil
}

// Public -
func Public(w http.ResponseWriter, path, host string) error {
	switch path {
	case "plst.m3u":
		return m3u(w, host)
	case "xml.gz":
		return transfer("ott.tv.planeta.tc/epg/program.xml.gz", "", w)
	default:
		return transfer(path, host, w)
	}
}
