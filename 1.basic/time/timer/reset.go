package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	log.Println("✔︎ resetBeforeFired")
	resetBeforeFired()
	fmt.Println()

	log.Println("✘ wrongResetAfterFired")
	wrongResetAfterFired()
	fmt.Println()

	log.Println("✔︎ correctResetAfterFired")
	correctResetAfterFired()
	fmt.Println()

	log.Println("✔︎ stop n times")
	stopMore()
	fmt.Println()

	log.Println("✘ stop n times but with drain")
	wrongStopMore()
	fmt.Println()

	log.Println("✘ too many receiving")
	wrongReceiveMore()
}

func resetBeforeFired() {
	timer := time.NewTimer(5 * time.Second)
	b := timer.Stop()
	log.Printf("stop: %t", b)
	timer.Reset(10 * time.Second)
	t := <-timer.C
	log.Printf("fired at %s", t.String())
}

func wrongResetAfterFired() {
	timer := time.NewTimer(5 * time.Millisecond)
	time.Sleep(time.Second)

	b := timer.Stop()
	log.Printf("stop: %t", b)
	timer.Reset(10 * time.Second)
	t := <-timer.C
	log.Printf("fired at %s", t.String())
}

func correctResetAfterFired() {
	timer := time.NewTimer(5 * time.Millisecond)
	time.Sleep(time.Second)

	b := timer.Stop()
	log.Printf("stop: %t", b)
	if !b {
		<-timer.C
	}
	log.Printf("reset")
	timer.Reset(10 * time.Second)
	t := <-timer.C
	log.Printf("fired at %s", t.String())
}

func wrongReceiveMore() {
	timer := time.NewTimer(5 * time.Millisecond)
	t := <-timer.C
	log.Printf("fired at %s", t.String())

	t = <-timer.C
	log.Printf("receive again at %s", t.String())
}

func stopMore() {
	timer := time.NewTimer(5 * time.Millisecond)
	b := timer.Stop()
	log.Printf("stop: %t", b)
	time.Sleep(time.Second)
	b = timer.Stop()
	log.Printf("stop more: %t", b)
}

func wrongStopMore() {
	timer := time.NewTimer(5 * time.Millisecond)
	b := timer.Stop()
	log.Printf("stop: %t", b)
	time.Sleep(time.Second)
	b = timer.Stop()
	if !b {
		<-timer.C
	}
	log.Printf("stop more: %t", b)
}
