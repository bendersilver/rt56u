package server

import (
	"bufio"
	"bytes"
	"io"
	"net"
	"net/http"
	"strings"

	"github.com/bendersilver/blog"
)

type replaceFn func(b io.ReadCloser, w *io.PipeWriter)

var cli = http.DefaultClient

func transfer(url string, w http.ResponseWriter, fn replaceFn) error {
	req, err := cli.Get(url)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return err
	}
	defer req.Body.Close()
	copyHeader(w.Header(), req.Header)
	if fn == nil {
		io.Copy(w, req.Body)
	} else {
		reader, writer := io.Pipe()
		defer reader.Close()
		go fn(req.Body, writer)
		io.Copy(w, reader)
	}
	return nil
}

func copyHeader(dst, src http.Header) {
	for k, vv := range src {
		for _, v := range vv {
			dst.Add(k, v)
		}
	}
	dst.Del("Content-Length")
	dst.Del("Connection")
}

func match(args ...string) bool {
	if len(args) < 2 {
		panic("http: invalid pattern")
	}

	for _, pattern := range args[1:] {
		if strings.Index(args[0], pattern) == 0 {
			return true
		}
	}
	return false
}
func replaceURL(host string) replaceFn {
	return func(b io.ReadCloser, w *io.PipeWriter) {
		defer w.Close()
		scanner := bufio.NewScanner(b)
		var line []byte
		for scanner.Scan() {
			line = scanner.Bytes()
			if bytes.HasPrefix(line, []byte("http://")) {
				line = bytes.Replace(line, []byte("http:/"), []byte("http://"+host), 1)
			}
			w.Write(line)
			w.Write([]byte("\n"))
		}
		line = nil
	}
}

var block = newBlockList()

func handler(w http.ResponseWriter, r *http.Request) {
	h, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		blog.Error(err)
	} else {
		if block.Check(h) {
			http.Error(w, "you block", http.StatusUnauthorized)
			return
		}
	}

	switch r.Method {
	case http.MethodPost, http.MethodGet:
		var path string
		if ar := strings.Split(r.URL.Path, "/"); len(ar) > 1 {
			path = ar[1]
		}
		switch {

		case path == "jsonAPI":
			if strings.HasPrefix(h, "192.168") {
				jsonAPI(w, r)
			} else {
				block.Add(h)
			}
		case path == "plst.m3u":
			playList(w, r)
		case match(path, "www", "js", "css", "favicon.ico"):
			if strings.HasPrefix(h, "192.168") {
				fileServ(w, r)
			} else {
				block.Add(h)
			}
		case path == "xml.gz":
			transfer("http://ott.tv.planeta.tc/epg/program.xml?fields=desc&fields=icon", w, nil)
		case match(path, "cache"):
			transfer("http:/"+r.RequestURI, w, nil)
		case path == "playlist":
			transfer("http://playlist.tv.planeta.tc"+r.RequestURI, w, replaceURL(r.Host))
		default:
			http.Error(w, "page not found", http.StatusNotFound)
			block.Add(h)
		}
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		block.Add(h)
	}
}

// New -
func New(port string) (svr *http.Server) {

	go block.unblock()
	svr = new(http.Server)
	svr.Addr = port
	svr.Handler = http.HandlerFunc(handler)
	// svr.ErrorLog =
	return
}
