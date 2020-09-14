package main

import (
	"errors"
	"io"
)

func main() {
	io.EOF = errors.New("我们自己定义的EOF")
}
