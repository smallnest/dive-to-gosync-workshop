package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {
	g, ctx := errgroup.WithContext(context.Background())

	g.Go(func() error {
		time.Sleep(5 * time.Second)
		fmt.Println("exec #1")
		return nil
	})

	g.Go(func() error {
		time.Sleep(10 * time.Second)
		fmt.Println("exec #2")
		return errors.New("failed to exec #2")
	})

	g.Go(func() error {
		time.Sleep(15 * time.Second)
		fmt.Println("exec #3")
		return nil
	})

	select {
	case <-ctx.Done():
		fmt.Println("done:", ctx.Err())
	}

	if err := g.Wait(); err == nil {
		fmt.Println("Successfully exec all")
	} else {
		fmt.Println("Successfully exec:", err)
	}

}
