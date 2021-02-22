package handler

import "io"

// Err400 -
func Err400(w io.Writer) {
	w.Write([]byte("HTTP/1.1 400 Bad Request\r\n\r\nBad Request"))
}

// Err404 -
func Err404(w io.Writer) {
	w.Write([]byte("HTTP/1.1 404 Not Found\r\n\r\nNot Found"))
}

// Err405 -
func Err405(w io.Writer) {
	w.Write([]byte("HTTP/1.1 405 Method Not Allowed\r\n\r\nMethod Not Allowed"))
}

// Err500 -
func Err500(w io.Writer, s string) {
	w.Write([]byte("HTTP/1.1 500 Internal Server Error\r\n\r\n"))
	w.Write([]byte(s))
	w.Write([]byte("\r\n"))
}

// Status200 -
func Status200(w io.Writer) {
	w.Write([]byte([]byte("HTTP/1.1 200 OK\n")))
}
