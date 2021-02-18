package main

import (
	"runtime"

	"github.com/bendersilver/blog"
	"github.com/bot/rt56u/server"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	svr := server.New(":33880")
	blog.Fatal(svr.ListenAndServe())

}
