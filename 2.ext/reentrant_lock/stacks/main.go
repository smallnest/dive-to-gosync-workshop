//GOTRACEBACK=1
package main

import (
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"
)

func main() {

	go func() {
		for {
			time.Sleep(time.Second)
		}
	}()

	go func() {
		for {
			time.Sleep(time.Minute)
		}
	}()

	time.Sleep(5 * time.Second)
	DumpStacks()

	// panic("")
}

func setupSigusr1Trap() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM)
	go func() {
		for range c {
			DumpStacks()
		}
	}()
}
func DumpStacks() {
	buf := make([]byte, 16384)
	buf = buf[:runtime.Stack(buf, true)]
	fmt.Printf("=== BEGIN goroutine stack dump ===\n%s\n=== END goroutine stack dump ===", buf)
}
