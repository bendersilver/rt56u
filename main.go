package main

import (
	"bufio"
	"bytes"
	"io"
	"net"
	"os"
	"runtime"
	"strings"

	"github.com/bendersilver/simplog"
	"github.com/bot/rt56u/handler"
)

var token string

func init() {
	f, err := os.OpenFile("./.env", os.O_RDONLY, 0644)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		spl := strings.Split(scanner.Text(), "=")
		if len(spl) > 1 {
			os.Setenv(spl[0], strings.Join(spl[1:], "="))
		}
	}
	for _, v := range []string{"TOKEN", "DIST", "GOB"} {
		if _, ok := os.LookupEnv(v); !ok {
			panic(v)
		}
	}
	token = os.Getenv("TOKEN")
}

func readMethodPath(r io.Reader) (method, pth string, err error) {
	var b = make([]byte, 1)
	var buf []byte
	for {
		_, err = r.Read(b)
		if err != nil {
			break
		}
		if b[0] == '\n' {
			break
		}
		if b[0] == ' ' {
			if method == "" {
				method = string(buf)
				buf = nil
				continue
			} else if pth == "" {
				pth = string(buf)
				buf = nil
				continue
			}
		}
		buf = append(buf, b...)
	}
	return
}

func passHead(r io.Reader) (err error) {
	var b = make([]byte, 4)
	for {
		for i := 0; i < 3; i++ {
			b[i] = b[i+1]
		}
		_, err = r.Read(b[3:4])
		if err != nil {
			return err
		}
		if bytes.Equal(b, []byte("\r\n\r\n")) {
			break
		}
	}
	return
}

// Handler -
func Handler(con *net.TCPConn) {
	defer con.CloseWrite()
	defer con.CloseRead()
	m, err := handler.NewMIMEHeader(con)
	if err != nil {
		handler.Err500(con, "«Internal server error»")
		simplog.Error(err)
		return
	}

	m.Private = handler.IsPrivate(con.RemoteAddr().String())
	simplog.Debug(m.Private, m.Path, strings.Index(m.Path, os.Getenv("TOKEN")))
	switch m.Method {
	case "GET":
		// if m.Private {
		// 	if m.Path == "/" || m.Path == "" {
		// 		m.Path = "/index.html"
		// 	}
		// 	handler.Private(con, m)
		// } else
		if strings.Index(m.Path, os.Getenv("TOKEN")) == 1 {
			handler.Public(con, m)
		}
		// if len(splitPath) >= 3 && splitPath[1] == token {
		// 	handler.Public(con, strings.Join(splitPath[2:], "/"))
		// } else if ptivate {
		// 	if m.Path == "/" || m.Path == "" {
		// 		m.Path = "/index.html"
		// 	}
		// 	handler.Private(con, m.Path)
		// } else {
		// 	handler.Err404(con)
		// }
	case "POST":
		// passHead(con)
		// if ptivate && len(splitPath) == 3 && splitPath[1] == "jsonAPI" {
		// 	handler.PrivatePost(con, splitPath[2])
		// } else {
		// 	handler.Err404(con)
		// }
	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	listener, err := net.ListenTCP("tcp", &net.TCPAddr{Port: 33880})
	if err != nil {
		simplog.Fatal(err)
	}
	defer listener.Close()

	for {
		con, err := listener.AcceptTCP()
		if err != nil {
			continue
		}

		Handler(con)
	}
}
