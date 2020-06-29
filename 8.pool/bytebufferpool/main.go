package main

import (
	"fmt"

	"github.com/valyala/bytebufferpool"
)

func main() {
	bb := bytebufferpool.Get()

	bb.WriteString("first line\n")
	bb.Write([]byte("second line\n"))
	bb.B = append(bb.B, "third line\n"...)

	fmt.Printf("bytebuffer contents=%q", bb.B)

	bytebufferpool.Put(bb)
}
