package main

import (
	"log"
	"math/rand"
	"sort"
	"time"
)

// H2O 水分子工厂.
type H2O struct {
	// 氢原子的信号量
	semaH chan int
	// 氧原子的信号量
	semaO chan int
}

// 创建一个水分子工厂.
func New() *H2O {
	return &H2O{
		semaH: make(chan int, 2),
		semaO: make(chan int, 0),
	}
}

// 被氢原子goroutine调用，满足条件的时候就会提供一个H原子来产生一个水分子.
func (h2o *H2O) hydrogen(releaseHydrogen func()) {
	h2o.semaO <- 0
	releaseHydrogen()

	<-h2o.semaH
}

// 被氧原子goroutine调用,满足条件的时候就会一个O原子来产生一个水分子.
func (h2o *H2O) oxygen(releaseOxygen func()) {
	h2o.semaH <- 0
	h2o.semaH <- 0

	releaseOxygen()

	<-h2o.semaO
	<-h2o.semaO
}

func main() {
	var ch chan string
	releaseHydrogen := func() {
		ch <- "H"
	}
	releaseOxygen := func() {
		ch <- "O"
	}

	// goroutine数
	M := 2
	// 每个goroutine产生的原子数
	N := 10
	ch = make(chan string, M*N*3)
	h2o := New()

	for k := 0; k < M; k++ {
		go func() {
			for i := 0; i < N; i++ {
				time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
				h2o.hydrogen(releaseHydrogen)
			}
		}()
		go func() {
			for i := 0; i < N; i++ {
				time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
				h2o.hydrogen(releaseHydrogen)
			}
		}()
		go func() {
			for i := 0; i < N; i++ {
				time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
				h2o.oxygen(releaseOxygen)
			}
		}()
	}

	time.Sleep(5 * time.Second)

	s := make([]string, 3)
	for i := 0; i < len(ch)/3; i++ {
		s[0] = <-ch
		s[1] = <-ch
		s[2] = <-ch
		sort.Strings(s)
		water := s[0] + s[1] + s[2]
		if water != "HHO" {
			log.Panicf("expect a water molecule but got %s", water)
		}
	}
}
