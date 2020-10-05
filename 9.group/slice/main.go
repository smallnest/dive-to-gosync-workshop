package main

import "fmt"

func main() {
	var s []int
	for i := 0; i < 100000; i++ {
		go func(n int) {
			s = append(s, n)
		}(i)
	}
	fmt.Printf("%d\n", s[0])
}
