package handler

import (
	"errors"
	"io"
	"net"
	"strings"
)

// Channel -
// Track represents an m3u track with a Name, Lengh, URI and a set of tags
type Channel struct {
	Name string
	URI  string
	Tags map[string]string
}

// read byte to saparator
func readValue(r io.Reader, endByte byte) (buf []byte, err error) {
	var b = make([]byte, 1)
	// read first quote
	if endByte == '"' {
		r.Read(b)
	}
	for {
		_, err = r.Read(b)
		if err != nil {
			return
		}
		switch b[0] {
		case endByte:
			return
		default:
			buf = append(buf, b...)
		}
	}
}

func readUpToTag(r io.Reader, tag []byte) (err error) {
	buf := make([]byte, len(tag))
	for {
		copy(buf, buf[1:])
		_, err = r.Read(buf[len(tag)-1 : len(tag)])
		if err != nil {
			return
		}
		if string(buf) == string(tag) {
			return
		}
	}
}

func isPlaylist(r io.Reader) (err error) {
	// pass http version
	readValue(r, ' ')
	// status code
	buf, err := readValue(r, ' ')
	if string(buf) != "200" {
		return errors.New(string(buf))
	}
	err = readUpToTag(r, []byte("#EXTM3U"))
	if err != nil {
		return errors.New("invalid m3u file format. Expected #EXTM3U file header")
	}
	return
}

func extF(r io.Reader) (ch *Channel, err error) {
	b := make([]byte, 1)
	ch = new(Channel)
	ch.Tags = make(map[string]string)

	var key, val []byte
Loop:
	for {
		_, err = r.Read(b)
		if err != nil {
			return
		}
		switch b[0] {
		case ' ':
			key, err = readValue(r, '=')
			if err != nil {
				return
			}
			val, err = readValue(r, '"')
			if err != nil {
				return
			}
			ch.Tags[string(key)] = string(val)
		case ',':
			break Loop
		case '\n':
			err = errors.New("missing name")
			return
		}
	}
	// read name
	key, err = readValue(r, '\n')
	if err != nil {
		return
	}
	ch.Name = strings.TrimSpace(string(key))
	// read url
	key, err = readValue(r, '\n')
	if err != nil {
		return
	}
	ch.URI = strings.TrimSpace(string(key))
	return
}

func getAll() (channels []*Channel, err error) {
	var d net.Conn
	d, err = net.Dial("tcp", "ott.tv.planeta.tc:80")
	if err != nil {
		return
	}
	defer d.Close()
	_, err = d.Write([]byte("GET /playlist/channels.m3u?4k&groupChannels=thematic&fields=epg,group&hlsQuality=min&hlsVideoOnly HTTP/1.0\r\nHosh: ott.tv.planeta.tc\r\nnUser-Agent: go-iptv\r\n\r\n"))
	if err != nil {
		return
	}
	err = isPlaylist(d)
	if err != nil {
		return
	}
	for {
		err = readUpToTag(d, []byte("#EXTINF"))
		if err != nil {
			if err == io.EOF {
				err = nil
			}
			break
		}
		ch, err := extF(d)
		if err != nil {
			break
		}
		if ch.Name == "" || ch.URI == "" {
			continue
		}
		channels = append(channels, ch)
	}
	return
}

func GetAllChannels() (channels []*Channel, err error) {
	return getAll()
}
