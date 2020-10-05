package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/uw-labs/sync/rungroup"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	g, ctx := rungroup.New(ctx)

	g.Go(func() error {
		return run(ctx, time.Second)
	})
	g.Go(func() error {
		time.Sleep(time.Millisecond * 50)
		return errors.New("component stopped")
	})

	err := g.Wait()
	fmt.Println(err)
}

func run(ctx context.Context, d time.Duration) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-time.After(d):
		return nil
	}
}
