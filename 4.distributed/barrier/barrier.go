package main

import (
	"log"
	"math/rand"
	"sync"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/concurrency"
	recipe "go.etcd.io/etcd/client/v3/experimental/recipes"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	barrier()
}

func barrier() {
	endpoints := []string{"http://127.0.0.1:2379"}
	cli, err := clientv3.New(clientv3.Config{Endpoints: endpoints})
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()

	barrierName := "my-test"

	b := recipe.NewBarrier(cli, barrierName)
	err = b.Hold()
	if err != nil {
		panic(err)
	}
	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		i := i
		go func() {
			b := recipe.NewBarrier(cli, barrierName)

			time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
			log.Println("enter for ID:", i)
			err := b.Wait()
			if err != nil {
				panic(err)
			}
			log.Println("entered for ID:", i)
			wg.Done()
		}()
	}

	time.Sleep(12 * time.Second)
	err = b.Release()
	if err != nil {
		panic(err)
	}

	wg.Wait()
}

func doubleBarrier() {
	endpoints := []string{"http://127.0.0.1:2379"}
	cli, err := clientv3.New(clientv3.Config{Endpoints: endpoints})
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()

	s, err := concurrency.NewSession(cli)
	if err != nil {
		log.Fatal(err)
	}
	defer s.Close()

	barrierName := "my-test"

	var wg sync.WaitGroup
	wg.Add(10)

	var leaveWG sync.WaitGroup
	leaveWG.Add(10)

	for i := 0; i < 10; i++ {
		i := i
		go func() {
			b := recipe.NewDoubleBarrier(s, barrierName, 10)

			time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
			log.Println("enter for ID:", i)
			b.Enter()
			log.Println("entered for ID:", i)
			wg.Done()

			time.Sleep(time.Duration(rand.Intn(20)) * time.Second)
			log.Println("leave for ID:", i)
			b.Leave()
			log.Println("left for ID:", i)
			leaveWG.Done()
		}()
	}

	wg.Wait()

	leaveWG.Wait()
}
