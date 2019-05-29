package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	timer := time.AfterFunc(time.Second, func() {
		fmt.Println("fired")
	})

	t := <-timer.C                        // nil
	log.Printf("fired at %s", t.String()) // 海枯石烂你也等不来
}
