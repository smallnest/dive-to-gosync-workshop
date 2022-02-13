package water

import (
	"context"
	"sync"

	"golang.org/x/sync/semaphore"
)

type H2O struct {
	semaH *semaphore.Weighted
	semaO *semaphore.Weighted
	wg    sync.WaitGroup
}

func New() *H2O {
	var wg sync.WaitGroup
	wg.Add(3)

	return &H2O{
		semaH: semaphore.NewWeighted(2),
		semaO: semaphore.NewWeighted(1),
		wg:    wg,
	}
}

func (h2o *H2O) hydrogen(releaseHydrogen func()) {
	h2o.semaH.Acquire(context.Background(), 1)

	// releaseHydrogen() outputs "H". Do not change or remove this line.
	releaseHydrogen()

	// wait
	h2o.wg.Done()
	h2o.wg.Wait()

	h2o.semaH.Release(1)
}

func (h2o *H2O) oxygen(releaseOxygen func()) {
	h2o.semaO.Acquire(context.Background(), 1)

	// releaseOxygen() outputs "O". Do not change or remove this line.
	releaseOxygen()

	// wait and reset
	h2o.wg.Done()
	h2o.wg.Wait()
	h2o.wg.Add(3)

	h2o.semaO.Release(1)
}
