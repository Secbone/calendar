package main

import (
	"fmt"
	"time"
	"sync"
	"context"
)

type Job func(ctx context.Context)

type Scheduler struct {
	group				*sync.WaitGroup
	cancellations	[]context.CancelFunc
}

func NewScheduler() *Scheduler {
	return &Scheduler{
		group: new(sync.WaitGroup),
		cancellations: make([]context.CancelFunc, 0),
	}
}

func (s *Scheduler) Add(ctx context.Context, j Job, interval time.Duration) {
	ctx, cancel := context.WithCancel(ctx)
	s.cancellations = append(s.cancellations, cancel)

	s.group.Add(1)
	go s.process(ctx, j, interval)
}

func (s *Scheduler) Stop() {
	for _, cancel := range s.cancellations {
		cancel()
	}
	s.group.Wait()
	fmt.Println("stop called")
}

func (s *Scheduler) process(ctx context.Context, j Job, interval time.Duration) {
	ticker := time.NewTicker(interval)

	for {
		select {
		case <-ticker.C:
			j(ctx)
		case <-ctx.Done():
			s.group.Done()
			return
		}
	}
}