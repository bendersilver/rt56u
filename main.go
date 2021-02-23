package main

import (
	"bufio"
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
	handler.Init()
}

// Handler -
func Handler(con *net.TCPConn) {
	defer con.Close()
	method, pth, host, err := handler.ParseHeader(con, 248)
	if err != nil {
		handler.Err400(con)
		return
	}
	var ptivate = handler.IsPrivate(con.RemoteAddr().String())
	// if host == "localhost:33880" {
	// 	ptivate = false
	// }
	switch method {
	case "GET":
		if pth == "/" || pth == "" {
			pth = "/index.html"
		}
		if ptivate {
			handler.Private(con, pth)
			return
		}
		splitPath := strings.Split(pth, "/")
		if len(splitPath) < 3 || splitPath[1] != os.Getenv("TOKEN") {
			handler.Err404(con)
			// ban
			return
		}
		handler.Public(con, strings.Join(splitPath[2:], "/"), host)
	case "POST":
		handler.PrivatePost(con, pth)
	default:
		handler.Err405(con)
		return
	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	listener, err := net.ListenTCP("tcp", &net.TCPAddr{Port: 33880})
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	for {
		con, err := listener.AcceptTCP()
		if err != nil {
			continue
		}
		go Handler(con)
	}
}
