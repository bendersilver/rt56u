package handler

import (
	"bufio"
	"bytes"
	"io"
	"net"
	"os"
	"strings"
)

// transfer -
func transfer(p, h string, w io.Writer) {
	sp := strings.Split(p, "/")
	d, p := sp[0], strings.Join(sp[1:], "/")

	dl, err := net.Dial("tcp", d+":80")
	if err != nil {
		Err500(w, err.Error())
		return
	}
	defer dl.Close()
	_, err = dl.Write([]byte("GET /" + p + " HTTP/1.0\nHosh: " + d + "\nUser-Agent: go-iptv\n\n"))
	if err != nil {
		Err500(w, err.Error())
		return
	}
	var buf = make([]byte, 256)
	var i int
	if strings.HasPrefix(p, "playlist/") {
		var head bool = true
		uri := []byte("http://" + h + "/" + os.Getenv("TOKEN") + "/")
		for {
			if _, err = dl.Read(buf[i : i+1]); err != nil {
				w.Write(buf[:i])
				break
			}
			if head {
				if i == 1 && buf[1] == '\n' {
					head = false
				}
			} else {
				// line startwish http://
				//                     ^^
				if i == 6 && buf[5] == '/' && buf[6] == '/' {

					for ix, b := range uri {
						buf[ix] = b
						i = ix
					}
				}
			}
			if buf[i] == '\n' {
				if !bytes.HasPrefix(buf, []byte("Content-Length")) {
					w.Write(buf[:i+1])
				}
				i = -1
			}
			i++
		}
	} else {
		io.Copy(w, dl)
	}
}

func m3u(w io.Writer, host string) {
	item, ok := static["/plst.m3u"]
	if !ok {
		Err404(w)
		return
	}
	f, err := os.OpenFile(item.Path, os.O_RDONLY, 0644)
	if err != nil {
		Err500(w, err.Error())
	} else {
		defer f.Close()
		Status200(w)
		w.Write([]byte("Content-Encoding: identity\n"))
		w.Write([]byte(item.Type))
		token := os.Getenv("TOKEN")
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			if strings.HasPrefix(scanner.Text(), "http://") {
				io.WriteString(w,
					strings.Replace(scanner.Text(),
						"http:/",
						"http://"+host+"/"+token, 1))
			} else {
				io.WriteString(w, scanner.Text())
			}
			io.WriteString(w, "\n")
		}
	}
}

// Public -
func Public(con *net.TCPConn, path, host string) {
	switch path {
	case "plst.m3u":
		m3u(con, host)
	case "xml.gz":
		transfer("ott.tv.planeta.tc/epg/program.xml.gz", "", con)
	default:
		transfer(path, host, con)
	}
}
