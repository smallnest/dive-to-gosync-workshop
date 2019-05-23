package main

import (
	"log"
	"math/rand"
	"sync"
	"time"

	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/clientv3/concurrency"
	recipe "github.com/coreos/etcd/contrib/recipes"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	endpoints := []string{"http://127.0.0.1:2379"}
	cli, err := clientv3.New(clientv3.Config{Endpoints: endpoints})
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()

	var lockName = "my-lock"

	var wg sync.WaitGroup
	wg.Add(20)

	for i := 0; i < 10; i++ {
		go startLockSession(i, cli, lockName, &wg)
	}

	for i := 0; i < 10; i++ {
		go startRLockSession(10+i, cli, lockName, &wg)
	}

	wg.Wait()
}

func startLockSession(id int, cli *clientv3.Client, lockName string, wg *sync.WaitGroup) {
	defer wg.Done()

	// 为锁生成session
	s1, err := concurrency.NewSession(cli)
	if err != nil {
		log.Fatal(err)
	}
	defer s1.Close()
	m1 := recipe.NewRWMutex(s1, lockName)

	// 请求锁
	log.Println("acquiring lock for ID:", id)
	if err := m1.Lock(); err != nil {
		log.Fatal(err)
	}
	log.Println("acquired lock for ID:", id)

	time.Sleep(time.Duration(rand.Intn(10)) * time.Second)

	if err := m1.Unlock(); err != nil {
		log.Fatal(err)
	}
	log.Println("released lock for ID:", id)
}

func startRLockSession(id int, cli *clientv3.Client, lockName string, wg *sync.WaitGroup) {
	defer wg.Done()

	// 为锁生成session
	s1, err := concurrency.NewSession(cli)
	if err != nil {
		log.Fatal(err)
	}
	defer s1.Close()
	m1 := recipe.NewRWMutex(s1, lockName)

	// 请求锁
	log.Println("acquiring rlock for ID:", id)
	if err := m1.RLock(); err != nil {
		log.Fatal(err)
	}
	log.Println("acquired lock for ID:", id)

	time.Sleep(time.Duration(rand.Intn(10)) * time.Second)

	if err := m1.RUnlock(); err != nil {
		log.Fatal(err)
	}
	log.Println("released rlock for ID:", id)
}
