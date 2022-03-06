package main

import (
	"log"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// Semaphore定义
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

// 顾客
var (
	// 控制顾客的总数
	customerMutex    sync.Mutex
	customerMaxCount = 20
	customerCount    = 0

	// 沙发的容量
	sofaSema Semaphore = make(chan struct{}, 4)
)

// 收银台
var (
	// 同时只有一对理发师和顾客结账
	paySema Semaphore = make(chan struct{}, 1)
	// 顾客拿到发票才会离开，控制开票
	receiptSema Semaphore = make(chan struct{}, 1)
)

// 理发师工作
func barber(name string) {
	for {
		// 等待一个用户
		log.Println(name + "老师尝试请求一个顾客")
		sofaSema.Release()
		log.Println(name + "老师找到一位顾客，开始理发")

		randomPause(2000)

		log.Println(name + "老师理完发，等待顾客付款")
		paySema.Acquire()
		log.Println(name + "老师给付完款的顾客发票")
		receiptSema.Release()
		log.Println(name + "老师服务完一位顾客")

	}
}

// 模拟顾客陆陆续续的过来
func customers() {
	for {
		randomPause(500)

		go customer()
	}
}

// 顾客
func customer() {
	customerMutex.Lock()
	if customerCount == customerMaxCount {
		log.Println("没有空闲座位了，一位顾客离开了")
		customerMutex.Unlock()
		return
	}
	customerCount++
	customerMutex.Unlock()

	log.Println("一位顾客开始等沙发坐下")
	sofaSema.Acquire()
	log.Println("一位顾客找到空闲沙发坐下,直到被理发师叫起理发")

	paySema.Release()
	log.Println("一位顾客已付完钱")
	receiptSema.Acquire()
	log.Println("一位顾客拿到发票，离开")

	customerMutex.Lock()
	customerCount--
	customerMutex.Unlock()
}

func randomPause(max int) {
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(max)))
}

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
