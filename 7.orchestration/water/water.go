package water

import (
	"context"

	"github.com/marusama/cyclicbarrier"
	"golang.org/x/sync/semaphore"
)

type H2O struct {
	semaH *semaphore.Weighted
	semaO *semaphore.Weighted
	b1    cyclicbarrier.CyclicBarrier
	b2    cyclicbarrier.CyclicBarrier
}

func New() *H2O {
	return &H2O{
		semaH: semaphore.NewWeighted(2),
		semaO: semaphore.NewWeighted(1),
		b1:    cyclicbarrier.New(3),
		b2:    cyclicbarrier.New(3),
	}
}

func (h2o *H2O) hydrogen(releaseHydrogen func()) {
	h2o.semaH.Acquire(context.Background(), 1)

	// releaseHydrogen() outputs "H". Do not change or remove this line.
	releaseHydrogen()

	h2o.b1.Await(context.Background())
	h2o.b2.Await(context.Background())

	h2o.semaH.Release(1)
}

func (h2o *H2O) oxygen(releaseOxygen func()) {
	h2o.semaO.Acquire(context.Background(), 1)

	h2o.b1.Await(context.Background())
	// releaseOxygen() outputs "O". Do not change or remove this line.
	releaseOxygen()

	h2o.b2.Await(context.Background())

	h2o.semaO.Release(1)
}
