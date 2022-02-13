package main

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/fatih/color"
	"golang.org/x/sync/semaphore"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

// Chopstick 代表筷子.
type Chopstick struct{ sync.Mutex }

// Philosopher 代表哲学家.
type Philosopher struct {
	// 哲学家的名字
	name string
	// 左手一只和右手一只筷子
	leftChopstick, rightChopstick *Chopstick
	status                        string

	sema *semaphore.Weighted
}

// 无休止的进餐和冥想.
// 吃完睡(冥想、打坐), 睡完吃.
// 可以调整吃睡的时间来增加或者减少抢夺叉子的机会.
func (p *Philosopher) dine() {
	for {
		mark(p, "冥想")
		randomPause(10)

		mark(p, "饿了")
		p.sema.Acquire(context.Background(), 1)
		p.leftChopstick.Lock() // 先尝试拿起左手边的筷子
		mark(p, "拿起左手筷子")
		p.rightChopstick.Lock() // 再尝试拿起右手边的筷子
		mark(p, "拿起右手筷子")

		mark(p, "用膳")
		randomPause(10)

		p.rightChopstick.Unlock() // 先尝试放下右手边的筷子
		p.leftChopstick.Unlock()  // 再尝试拿起左手边的筷子
		p.sema.Release(1)
	}
}

// 随机暂停一段时
func randomPause(max int) {
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(max)))
}

// 显示此哲学家的状态
func mark(p *Philosopher, action string) {
	fmt.Printf("%s开始%s\n", p.name, action)
	p.status = fmt.Sprintf("%s开始%s\n", p.name, action)
}

func main() {
	go http.ListenAndServe("localhost:8972", nil)

	// 哲学家数量
	count := 5

	// 创建5根筷子
	chopsticks := make([]*Chopstick, count)
	for i := 0; i < count; i++ {
		chopsticks[i] = new(Chopstick)
	}

	sema := semaphore.NewWeighted(4)
	//
	names := []string{color.RedString("孔子"), color.MagentaString("庄子"), color.CyanString("墨子"), color.GreenString("孙子"), color.WhiteString("老子")}

	// 创建哲学家, 分配给他们左右手边的叉子，领他们做到圆餐桌上.
	philosophers := make([]*Philosopher, count)
	for i := 0; i < count; i++ {
		philosophers[i] = &Philosopher{
			name: names[i], leftChopstick: chopsticks[i], rightChopstick: chopsticks[(i+1)%count],
			sema: sema,
		}
		go philosophers[i].dine()
	}

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs

	fmt.Println("退出中... 每个哲学家的状态:")
	for _, p := range philosophers {
		fmt.Print(p.status)
	}
}
