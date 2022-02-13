package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"sort"
	"time"

	"github.com/marusama/cyclicbarrier"
	"golang.org/x/sync/semaphore"
)

// H2O 水分子工厂.
type H2O struct {
	// 氢原子的信号量
	semaH *semaphore.Weighted
	// 氧原子的信号量
	semaO *semaphore.Weighted
	// 等待水分子的产生
	b cyclicbarrier.CyclicBarrier
}

// 创建一个水分子工厂.
func New() *H2O {
	return &H2O{
		semaH: semaphore.NewWeighted(2),
		semaO: semaphore.NewWeighted(1),
		b:     cyclicbarrier.New(3),
	}
}

// 被氢原子goroutine调用，满足条件的时候就会提供一个H原子来产生一个水分子.
func (h2o *H2O) hydrogen(releaseHydrogen func()) {
	// 准备一个H原子填坑
	h2o.semaH.Acquire(context.Background(), 1)
	releaseHydrogen()
	// 等待栅栏(另一个H原子和O原子的坑填好后栅栏开启)
	h2o.b.Await(context.Background())
	// 释放H原子的坑
	h2o.semaH.Release(1)
}

// 被氧原子goroutine调用,满足条件的时候就会一个O原子来产生一个水分子.
func (h2o *H2O) oxygen(releaseOxygen func()) {
	// 准备一个O原子填坑
	h2o.semaO.Acquire(context.Background(), 1)
	releaseOxygen()
	// 等待栅栏(另两个H原子)
	h2o.b.Await(context.Background())
	// 释放O原子的坑
	h2o.semaO.Release(1)
}

func main() {
	var ch chan string
	releaseHydrogen := func() {
		ch <- "H"
		fmt.Print("H")
	}
	releaseOxygen := func() {
		ch <- "O"
		fmt.Print("O")
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
