package main

import (
	"bytes"
	"io"
	"net"
	"os"
	"path"
	"runtime"
	"strings"

	"github.com/bendersilver/blog"
	"github.com/bot/rt56u/handler"
)

var token string

func init() {
	if f, ok := os.LookupEnv("TOKEN"); !ok {
		blog.Fatal("set TOKEN")
	} else {
		token = f
	}
}

func readHeader(conn *net.TCPConn) (method, path, proto string, head map[string]string) {
	head = make(map[string]string)
	buf := make([]byte, 2048)
	var first bool = true
	var vals []string
	var i, last int
	for {
		_, err := conn.Read(buf[i : i+1])
		if i >= 4 && bytes.Equal(buf[i-3:i+1], []byte("\r\n\r\n")) {
			break
		}
		if err != nil {
			break
		}
		if buf[i] == '\n' {
			if first {
				vals = strings.Split(strings.TrimSpace(string(buf[last:i])), " ")
				if len(vals) == 3 {
					method = vals[0]
					path = vals[1]
					proto = vals[2]

				}
				first = false
			} else {
				vals = strings.Split(strings.TrimSpace(string(buf[last:i])), ": ")
				if len(vals) >= 2 {
					head[vals[0]] = strings.Join(vals[1:], ": ")
				}
			}
			last = i + 1
		}
		if i == len(buf) {
			break
		}
		i++
	}
	return
}

// hasPrefixMulti -
func hasPrefixMulti(s string, prefix ...string) bool {
	for _, p := range prefix {
		if strings.HasPrefix(s, p) {
			return true
		}
	}
	return false
}

func removeKey(s string) string {
	return strings.Join(strings.Split(s, "/")[2:], "/")
}

// Func -
func Func(con *net.TCPConn) {

	method, pth, _, head := readHeader(con)
	// blog.Debug(method, pth, proto, head)
	defer con.Close()
	defer con.CloseRead()
	defer con.CloseWrite()

	rd, wr := io.Pipe()
	defer rd.Close()
	switch method {
	case "GET":
		switch {
		case strings.HasPrefix(pth, path.Join("/", token+"/")):
			go handler.Transfer(removeKey(pth), head["Host"], wr)
		case pth == path.Join("/", token, "/xml.gz"):
			go handler.Transfer("ott.tv.planeta.tc/epg/program.xml?fields=desc&fields=icon", "", wr)
		case pth == path.Join("/", token, "/plst.m3u"):
			go handler.M3U(wr, head["Host"])
		case strings.HasPrefix(pth, path.Join("/", token, "/cache")):
			blog.Debug(token)
		case hasPrefixMulti(pth, "/css/", "/js/", "/favicon.ico", "/"):
			go handler.Static(wr, pth)
		default:
			go handler.Err404(wr)
		}
	case "POST":
		go handler.Err404(wr)
	default:
		go handler.Err405(wr)
	}

	io.Copy(con, rd)
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	listener, err := net.ListenTCP("tcp", &net.TCPAddr{Port: 33880})
	if err != nil {
		blog.Fatal(err)
	}
	defer listener.Close()
	for {
		con, err := listener.AcceptTCP()
		if err != nil {
			blog.Error(err)
			continue
		}
		go Func(con)
	}

	// runtime.GOMAXPROCS(runtime.NumCPU())
	// svr := server.New(":33880")
	// blog.Fatal(svr.ListenAndServe())

}
