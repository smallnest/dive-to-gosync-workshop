package main

import (
	"log"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Semaphore chan struct{}

func (s Semaphore) Acquire() {
	s <- struct{}{}
}

func (s Semaphore) TryAcquire() bool {
	select {
	case s <- struct{}{}: // 还有空位子
		return true
	default: // 没有空位子了,离开
		return false
	}
}

func (s Semaphore) Release() {
	<-s
}

var seats = make(Semaphore, 10)

func main() {
	// 托尼、凯文、艾伦理发师三巨头
	go barber("Tony")
	go barber("Kevin")
	go barber("Allen")

	go customers()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs
}

func randomPause(max int) {
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(max)))
}

// 理发师
func barber(name string) {
	for {
		// 等待一个用户
		log.Println(name + "老师尝试请求一个顾客")
		seats.Release()
		log.Println(name + "老师找到一位顾客，开始理发")

		randomPause(2000)
	}
}

// 模拟顾客陆陆续续的过来
func customers() {
	for {
		randomPause(1000)

		go customer()
	}
}

// 顾客
func customer() {
	if ok := seats.TryAcquire(); ok {
		log.Println("一位顾客开始坐下排队理发")
	} else {
		log.Println("没有空闲座位了，一位顾客离开了")
	}
}
