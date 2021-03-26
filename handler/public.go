package handler

import (
	"bytes"
	"errors"
	"io"
	"io/ioutil"
	"net"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/bendersilver/simplog"
)

// transferFull -
func transferFull(uri string, m *MIMEHeader, w io.Writer, changeURL bool) (err error) {
	sp := strings.Split(uri, "/")
	if len(sp) < 2 {
		return errors.New("incorrect url")
	}

	host, path := sp[0], strings.Join(sp[1:], "/")
	d, err := net.DialTimeout("tcp", host+":80", time.Second*5)
	if err != nil {
		return err
	}
	defer d.Close()

	_, err = d.Write([]byte("GET /" + path + " " + m.Proto + "\r\n"))
	if err != nil {
		return
	}

	simplog.Notice("Копируем заголовки запроса")
	// Копируем заголовки запроса
	m.Header["Host"] = host
	for k, v := range m.Header {
		d.Write([]byte(k + ": " + v + "\r\n"))
	}
	d.Write([]byte("\r\n"))

	dHead, err := NewMIMEHeader(d)
	if err != nil {
		return err
	}

	// Копируем заголовки ответа
	w.Write([]byte(dHead.Method + " " + dHead.Path + " " + dHead.Proto + "\r\n"))
	for k, v := range dHead.Header {
		w.Write([]byte(k + ": " + v + "\r\n"))
	}
	w.Write([]byte("\r\n"))
	i, err := strconv.Atoi(string(dHead.Header["Content-Length"]))
	if err != nil {
		return err
	}
	io.CopyN(w, d, int64(i))
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

// http://ott.tv.planeta.tc/plst.m3u
func m3u(w io.Writer, m *MIMEHeader) {
	host := m.Header["Host"]

	// копируем заголовки оригинального ответа
	d, err := net.Dial("tcp", "ott.tv.planeta.tc:80")
	if err != nil {
		return
	}
	defer d.Close()
	_, err = d.Write([]byte("GET /plst.m3u HTTP/1.0\r\n"))
	if err != nil {
		return
	}
	m.Header["Host"] = "ott.tv.planeta.tc"
	for k, v := range m.Header {
		d.Write([]byte(k + ": " + v + "\r\n"))
	}
	d.Write([]byte("\r\n"))

	dHead, err := NewMIMEHeader(d)
	if err != nil {
		simplog.Error(err)
	}

	buf, err := ioutil.ReadFile(plstFile)
	if err != nil {
		return
	}
	buf = bytes.ReplaceAll(buf, []byte("\nhttp://"), []byte("\nhttp://"+host+"/"+os.Getenv("TOKEN")+"/"))
	dHead.Header["Content-Length"] = strconv.Itoa(len(buf))
	// Копируем заголовки ответа
	w.Write([]byte(dHead.Method + " " + dHead.Path + " " + dHead.Proto + "\r\n"))
	for k, v := range dHead.Header {
		w.Write([]byte(k + ": " + v + "\r\n"))
	}
	w.Write([]byte("\r\n"))
	w.Write(buf)
}

// Public -
func Public(w io.Writer, m *MIMEHeader) {
	split := strings.Split(m.Path, "/")
	if len(split) < 3 {
		return
	}
	simplog.Debug(split[2])
	switch strings.Join(split[2:], "/") {
	case "plst.m3u":
		m3u(w, m)
	case "xml.gz":
		err := transferFull("ott.tv.planeta.tc/epg/program.xml.gz", m, w, false)
		if err != nil {
			simplog.Error(err)
		}
	default:
		var replace bool
		if split[2] == "playlist.tv.planeta.tc" {
			replace = true
		}
		transferFull(strings.Join(split[2:], "/"), m, w, replace)
		// transfer(path, host, con)
	}
}
