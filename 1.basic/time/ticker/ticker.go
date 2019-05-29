package main

import (
	"log"
	"time"
)

func main() {
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()
	log.Println("create a ticker")

	t := <-ticker.C
	log.Println("the first tick:", t.String())

	time.Sleep(15 * time.Second)
	t = <-ticker.C
	log.Println("then the second tick:", t.String())

	t = <-ticker.C
	log.Println("then the third tick:", t.String())
}
