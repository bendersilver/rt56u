package handler

import (
	"bufio"
	"io"
	"os"
	"path"
	"strings"

	"github.com/bendersilver/blog"
)

var staticPath, token string

func init() {
	if f, ok := os.LookupEnv("DIST"); !ok {
		blog.Fatal("set DIST")
	} else {
		staticPath = f
	}
	token = os.Getenv("TOKEN")
}

// M3U - playlist
func M3U(pw *io.PipeWriter, host string) {
	f, err := os.OpenFile(path.Join(staticPath, "plst.m3u"), os.O_RDONLY, 0644)
	if err != nil {
		Err404(pw)
	} else {
		defer pw.Close()
		defer f.Close()
		pw.Write([]byte("HTTP/1.1 200 OK\n"))
		pw.Write([]byte("Content-Encoding: identity\n"))
		pw.Write([]byte("Content-Type: application/x-mpegurl; charset=utf-8\r\n\r\n"))

		if IsPrivate(host) {
			io.Copy(pw, f)
		} else {
			scanner := bufio.NewScanner(f)
			for scanner.Scan() {
				if strings.HasPrefix(scanner.Text(), "http://") {
					io.WriteString(pw,
						strings.Replace(scanner.Text(),
							"http:/",
							"http://"+host+"/"+token, 1))
				} else {
					io.WriteString(pw, scanner.Text())
				}
				io.WriteString(pw, "\n")
			}
		}
	}
}

var ext = map[string]string{
	".html": "Content-Type: text/html; charset=utf-8\r\n\r\n",
	".css":  "Content-Type: text/css; charset=utf-8\r\n\r\n",
	".js":   "Content-Type: text/javascript; charset=utf-8\r\n\r\n",
	".ico":  "Content-Type: image/x-icon\r\n\r\n",
	".map":  "Content-Type: application/json; charset=utf-8n\r\n\r\n",
}

// Static - static files
func Static(pw *io.PipeWriter, p string) {
	if p == "/" {
		p = "/index.html"
	}
	ct, ok := ext[path.Ext(p)]
	if !ok {
		Err404(pw)
		return
	}
	f, err := os.OpenFile(path.Join(staticPath, p), os.O_RDONLY, 0644)
	if err != nil {
		Err404(pw)
	} else {
		defer pw.Close()
		defer f.Close()
		pw.Write([]byte("HTTP/1.1 200 OK\n"))
		pw.Write([]byte("Accept-Ranges: bytes\n"))
		pw.Write([]byte(ct))
		io.Copy(pw, f)
	}
}
