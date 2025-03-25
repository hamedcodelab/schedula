package schedula

import (
	"context"
	"sync"
	"time"
)

type scheduler struct {
	workers map[string]worker
	stop    chan struct{}
	wg      sync.WaitGroup
}

func NewScheduler() Scheduler {
	return &scheduler{
		workers: make(map[string]worker),
	}
}

func (s *scheduler) AddWorker(name string, ticker time.Duration, w Worker) {
	s.workers[name] = worker{
		Worker:  w,
		TimeRun: ticker,
	}
}

func (s *scheduler) RemoveWorker(name string) {
	delete(s.workers, name)
}

func (s *scheduler) RunWorker(ctx context.Context, name string) {
	s.wg.Add(1)
	go func(w worker) {
		ticker := time.NewTicker(w.TimeRun)
		defer s.wg.Done()
		for {
			select {
			case <-ctx.Done():
				s.StopScheduler()
				return
			case <-s.stop:
				s.StopScheduler()
				return
			case <-ticker.C:
				s.workers[name].Run(ctx)
			}
		}
	}(s.workers[name])
}

func (s *scheduler) StopScheduler() {
	s.stop <- struct{}{}
	s.wg.Wait()
}
