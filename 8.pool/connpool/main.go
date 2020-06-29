package main

import (
	"fmt"
	"net"

	"github.com/fatih/pool"
)

func main() {
	factory := func() (net.Conn, error) { return net.Dial("tcp", "baidu.com:80") }

	p, err := pool.NewChannelPool(5, 30, factory)
	if err != nil {
		panic(err)
	}

	conn, err := p.Get()
	conn.Close()

	fmt.Println(p.Len())
}
