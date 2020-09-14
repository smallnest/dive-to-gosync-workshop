package main

import (
	"time"
)

type Counter struct {
	Website      string
	Start        time.Time
	PageCounters map[string]int
}

func main() {
	var c Counter
	c.Website = "baidu.com"

	c.PageCounters["/"]++
}
