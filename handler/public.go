package handler

import (
	"bytes"
	"errors"
	"io"
	"net"
	"strconv"
	"strings"
	"time"
)

// transferFull -
func transferFull(uri string, con *net.TCPConn) (err error) {
	sp := strings.Split(uri, "/")
	if len(sp) < 2 {
		err = errors.New("incorrect url")
		return
	}
	var host, path string
	var d net.Conn

	host, path = sp[0], strings.Join(sp[1:], "/")
	d, err = net.DialTimeout("tcp", host+":80", time.Second*5)
	if err != nil {
		return
	}
	defer d.Close()
	_, err = d.Write([]byte("GET /" + path + " HTTP/1.0\r\nHosh: " + host + "\r\n"))
	if err != nil {
		return
	}
	// Копируем заголовки запроса
	buf := make([]byte, 4)
	for {
		copy(buf, buf[1:])
		_, err = con.Read(buf[3:4])
		if err != nil {
			break
		}
		d.Write(buf[3:4])
		if string(buf) == "\r\n\r\n" {
			break
		}
	}
	// Content-Length
	// Копируем заголовки ответа и ищем размерность тела ответа
	buf = make([]byte, len([]byte("Content-Length: ")))
	var ok bool
	var length []byte
	var l int = len(buf)
	for {
		copy(buf, buf[1:])
		_, err = d.Read(buf[l-1 : l])
		if err != nil {
			break
		}
		con.Write(buf[l-1 : l])
		if !ok && string(buf) == "Content-Length: " {
			buf = make([]byte, 1)
		num:
			for {
				_, err = d.Read(buf)
				con.Write(buf)
				if buf[0] >= 48 && buf[0] <= 57 {
					length = append(length, buf...)
				} else {
					break num
				}
			}
			l = 4
			buf = make([]byte, l)
			ok = true
		}
		if ok && string(buf) == "\r\n\r\n" {
			break
		}
	}
	i, err := strconv.Atoi(string(length))
	if err != nil {
		return err
	}
	io.CopyN(con, d, int64(i))
	return
}

// transfer -
func transfer(uri string, con *net.TCPConn) (err error) {
	sp := strings.Split(uri, "/")
	if len(sp) < 2 {
		err = errors.New("incorrect url")
		return
	}
	var host, path string
	var d net.Conn

	host, path = sp[0], strings.Join(sp[1:], "/")
	d, err = net.DialTimeout("tcp", host+":80", time.Second*5)
	if err != nil {
		return
	}
	defer d.Close()
	_, err = d.Write([]byte("GET /" + path + " HTTP/1.0\nHosh: " + host + "\nUser-Agent: go-iptv\n\n"))
	if err != nil {
		return
	}
	var b = make([]byte, 1)
	var buf []byte
	for {
		_, err = con.Read(b)
		if err != nil {
			break
		}
		buf = append(buf, b...)
		if b[0] == '\n' {
			if len(buf) == 1 {
				break
			} else {
				if bytes.Index(buf, []byte("Content-Length")) == -1 {
					con.Write(buf)
				}
				buf = nil
			}
		}
	}
	io.Copy(con, d)
	return
}

func m3u(w *net.TCPConn) {
	// item, ok := static["/plst.m3u"]
	// if !ok {
	// 	Err404(w)
	// 	return
	// }
	// f, err := os.OpenFile(item.Path, os.O_RDONLY, 0644)
	// if err != nil {
	// 	Err500(w, err.Error())
	// } else {
	// 	defer f.Close()
	// 	Status200(w)
	// 	w.Write([]byte("Content-Encoding: identity\n"))
	// 	w.Write([]byte(item.Type))
	// 	token := os.Getenv("TOKEN")
	// 	scanner := bufio.NewScanner(f)
	// 	for scanner.Scan() {
	// 		if strings.HasPrefix(scanner.Text(), "http://") {
	// 			io.WriteString(w,
	// 				strings.Replace(scanner.Text(),
	// 					"http:/",
	// 					"http://"+w.LocalAddr().String()+"/"+token, 1))
	// 		} else {
	// 			io.WriteString(w, scanner.Text())
	// 		}
	// 		io.WriteString(w, "\n")
	// 	}
	// }
}

// Public -
func Public(con *net.TCPConn, path string) {
	switch path {
	case "plst.m3u":
		m3u(con)
	case "xml.gz":
		transferFull("ott.tv.planeta.tc/epg/program.xml.gz", con)
	default:
		// if strings.HasPrefix("playlist.tv.planeta.tc/") {
		// 	transfer(path, con)
		// } else {
		// 	transferFull(path, con)
		// }
		// transfer(path, host, con)
	}
}
