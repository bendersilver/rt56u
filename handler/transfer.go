package handler

import (
	"bufio"
	"io"
	"net"
	"strings"
)

// Transfer -
func Transfer(p, h string, pw *io.PipeWriter) {
	sp := strings.Split(p, "/")
	d, p := sp[0], strings.Join(sp[1:], "/")
	dl, err := net.Dial("tcp", d+":80")
	if err != nil {
		Err404(pw)
		return
	}
	_, err = dl.Write([]byte("GET /" + p + " HTTP/1.0\nHosh: " + d + "\nUser-Agent: go-iptv\n\n"))
	if err != nil {
		Err404(pw)
		return
	}
	if strings.HasPrefix(p, "playlist/") {
		scanner := bufio.NewScanner(dl)
		for scanner.Scan() {
			if strings.HasPrefix(scanner.Text(), "Content-Type:") || strings.HasPrefix(scanner.Text(), "Content-Length:") || strings.HasPrefix(scanner.Text(), "Connection:") {
				continue
			}
			if strings.HasPrefix(scanner.Text(), "http://") {
				io.WriteString(pw, strings.Replace(scanner.Text(), "http:/", "http://"+h+"/"+token, 1))
			} else {
				io.WriteString(pw, scanner.Text())
			}
			io.WriteString(pw, "\n")
		}
	} else {
		io.Copy(pw, dl)
	}
	pw.Close()
}
