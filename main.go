package main

import (
	"bufio"
	"io"
	"net"
	"os"
	"runtime"
	"strings"

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
	handler.Init()
}

func readMethodPath(r io.Reader) (method, pth string, err error) {
	var b = make([]byte, 1)
	var buf []byte
	for {
		_, err = r.Read(b)
		if err != nil {
			break
		}
		if b[0] == ' ' {
			if method == "" {
				method = string(buf)
				buf = nil
				continue
			} else {
				pth = string(buf)
				break
			}
		}
		buf = append(buf, b...)
	}
	return
}

func passHead(r io.Reader) (err error) {
	var b = make([]byte, 1)
	var buf []byte
	for {
		_, err = r.Read(b)
		if err != nil {
			break
		}
		if b[0] == '\r' {
			continue
		}
		buf = append(buf, b...)
		if b[0] == '\n' {
			if len(buf) == 1 {
				break
			} else {
				buf = nil
			}
		}
	}
	return
}

// Handler -
func Handler(con *net.TCPConn) {
	defer con.CloseWrite()
	defer con.CloseRead()
	var err error
	var method, pth string

	method, pth, err = readMethodPath(con)
	if err != nil {
		handler.Err500(con, err.Error())
		return
	}

	err = passHead(con)
	if err != nil {
		handler.Err500(con, err.Error())
		return
	}

	ptivate := handler.IsPrivate(con.LocalAddr().String())
	splitPath := strings.Split(pth, "/")

	switch method {
	case "GET":
		if len(splitPath) >= 3 && splitPath[1] == token {
			// handler.Public(con, strings.Join(splitPath[2:], "/"), con.LocalAddr().String())
		} else if ptivate {
			if pth == "/" || pth == "" {
				pth = "/index.html"
			}
			handler.Private(con, pth)
		} else {
			handler.Err404(con)
		}
		// handler.PrivatePost(con, pth)
	default:
		handler.Err405(con)
		return
	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	// ch, err := handler.GetAllChannels()
	// if err != nil {
	// 	simplog.Fatal(err)
	// }
	// for _, c := range ch {
	// 	simplog.Notice(c)
	// }
	// simplog.Fatal(handler.GetAllChannels())
	// listener, err := net.ListenTCP("tcp", &net.TCPAddr{Port: 33880})
	// if err != nil {
	// 	panic(err)
	// }
	// defer listener.Close()

	// for {
	// 	con, err := listener.AcceptTCP()
	// 	if err != nil {
	// 		continue
	// 	}
	// 	Handler(con)
	// }
}
