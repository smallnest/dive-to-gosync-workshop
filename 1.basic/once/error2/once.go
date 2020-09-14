package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"sync"
)

func main() {
	var once sync.Once
	var googleConn net.Conn

	once.Do(func() {
		googleConn, _ = net.Dial("tcp", "google.com:80")
	})

	googleConn.Write([]byte("GET / HTTP/1.1\r\nHost: google.com\r\n Accept: */*\r\n\r\n"))
	io.Copy(os.Stdout, googleConn)
}
