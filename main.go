package main

import (
	"bufio"
	"net/http"
	"os"
	"path"
	"runtime"
	"strings"
	"sync"
	"time"

	"github.com/bot/rt56u/handler"
)

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
}

var t <-chan time.Time
var mx sync.Mutex

func waitGC() {
	mx.Lock()
	defer mx.Unlock()
	<-t
	runtime.GC()
	t = nil
}

func setTimer() {
	mx.Lock()
	defer mx.Unlock()
	if t == nil {
		t = time.Tick(time.Second * 10)
		go waitGC()
	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	handler.Init()
	hlnd := func(w http.ResponseWriter, r *http.Request) {
		var err error
		var ptivate = handler.IsPrivate(r.RemoteAddr)
		splitPath := strings.Split(r.RequestURI, "/")

		switch r.Method {
		case http.MethodGet:
			if len(splitPath) >= 3 && splitPath[1] == os.Getenv("TOKEN") {
				err = handler.Public(w, strings.Join(splitPath[2:], "/"), r.Host)
			} else if ptivate {
				switch r.RequestURI {
				case "/jsonAPI/get":
					err = handler.PrivateGetJSON(w)
				case "/xml.gz":
					err = handler.PrivateXML(w)
				default:
					http.ServeFile(w, r, path.Join(os.Getenv("DIST"), r.RequestURI))
				}
			} else {
				http.Error(w, "Not Found", http.StatusNotFound)
			}
		case http.MethodPost:
			if ptivate && len(splitPath) == 3 && splitPath[1] == "jsonAPI" {
				err = handler.PrivatePOST(r.Body, splitPath[2])

			}
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		setTimer()

	}
	// go func() {
	// 	for range time.Tick(time.Second * 5) {
	// 		runtime.GC()
	// 	}
	// }()
	svr := new(http.Server)
	svr.Addr = ":33880"
	svr.Handler = http.HandlerFunc(hlnd)
	svr.ListenAndServe()
}
