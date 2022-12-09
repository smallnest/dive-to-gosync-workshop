package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/go-pkgz/syncs"
)

func main() {
	swg := syncs.NewSizedGroup(5)
	for i := 0; i < 10; i++ {
		swg.Go(func(ctx context.Context) {
			// doThings(ctx) // only 5 of these will run in parallel
		})
	}
	swg.Wait()

	ewg := syncs.NewErrSizedGroup(5, syncs.Preemptive) // error wait group with max size=5, don't try to start more if any error happened
	for i := 0; i < 10; i++ {
		ewg.Go(func() error { // Go here could be blocked if trying to run >5 at the same time
			// err := doThings(ctx) // only 5 of these will run in parallel
			return errors.New("some error")
		})
	}
	err := ewg.Wait()
	fmt.Println(err)
}
