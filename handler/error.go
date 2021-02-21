package handler

import "io"

// Err404 -
var Err404 = func(pw *io.PipeWriter) {
	defer pw.Close()
	pw.Write([]byte("HTTP/1.1 404 Not Found\r\n\r\nNot Found"))
}

// Err405 -
var Err405 = func(pw *io.PipeWriter) {
	defer pw.Close()
	pw.Write([]byte("HTTP/1.1 405 Method Not Allowed\r\n\r\nMethod Not Allowed"))
}
