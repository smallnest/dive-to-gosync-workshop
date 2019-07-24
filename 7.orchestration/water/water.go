package water

import (
	"context"

	"golang.org/x/sync/semaphore"
)

type H2O struct {
	semaH   *semaphore.Weighted
	semaO   *semaphore.Weighted
	phaser1 chan struct{}
	phaser2 chan struct{}
}

func New() *H2O {
	return &H2O{
		semaH:   semaphore.NewWeighted(2),
		semaO:   semaphore.NewWeighted(1),
		phaser1: make(chan struct{}, 2),
		phaser2: make(chan struct{}, 2),
	}
}

func (h2o *H2O) hydrogen(releaseHydrogen func()) {
	h2o.semaH.Acquire(context.Background(), 1)

	// releaseHydrogen() outputs "H". Do not change or remove this line.
	releaseHydrogen()

	<-h2o.phaser1 //wait oxygen goroutine for preparing oxygen

	<-h2o.phaser2 // wait water molecule

	h2o.semaH.Release(1)
}

func (h2o *H2O) oxygen(releaseOxygen func()) {
	h2o.semaO.Acquire(context.Background(), 1)

	h2o.phaser1 <- struct{}{} // trigger hydrogen goroutine that oxygen is ready
	h2o.phaser1 <- struct{}{} // trigger another hydrogen goroutine that oxygen is ready

	// releaseOxygen() outputs "O". Do not change or remove this line.
	releaseOxygen()

	h2o.phaser2 <- struct{}{} // trigger hydrogen goroutine to prepare next water molecule
	h2o.phaser2 <- struct{}{} // trigger another hydrogen goroutine to prepare next water molecule

	h2o.semaO.Release(1)
}
