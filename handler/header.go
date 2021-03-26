package handler

import (
	"bytes"
	"errors"
	"io"
)

type MIMEHeader struct {
	First   []byte
	Method  string
	Path    string
	Proto   string
	Header  map[string]string
	Private bool
}

func readLine(r io.Reader) ([]byte, error) {
	b := make([]byte, 512)
	var err error
	for i := 0; i <= 512; i++ {
		if i > 512 {
			return nil, errors.New("malformed HTTP first line")
		}
		_, err = r.Read(b[i : i+1])
		if err != nil {
			return nil, err
		}
		if b[i] == '\n' {
			return bytes.TrimSpace(b[:i+1]), nil
		}
	}
	return nil, errors.New("too long HTTP header line")
}

func NewMIMEHeader(r io.Reader) (m *MIMEHeader, err error) {
	m = new(MIMEHeader)
	m.Header = make(map[string]string)

	m.First, err = readLine(r)
	split := bytes.Split(m.First, []byte(" "))
	if len(split) < 3 {
		return nil, errors.New("malformed HTTP first line")
	}
	m.Method = string(split[0])
	m.Path = string(split[1])
	m.Proto = string(split[2])

	var b []byte
	for {
		b, err = readLine(r)
		if err != nil || len(b) == 0 {
			return
		}
		split = bytes.SplitN(b, []byte(": "), 2)
		if len(split) != 2 {
			return nil, errors.New("malformed HTTP head")
		}
		m.Header[string(split[0])] = string(split[1])
	}
}
