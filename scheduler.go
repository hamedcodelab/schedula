package schedula

import (
	"context"
	"log"
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
		stop:    make(chan struct{}),
		wg:      sync.WaitGroup{},
	}
}

func (s *scheduler) AddWorker(name string, typ WorkerSchemaType, ticker time.Duration, w Worker) {
	switch typ {
	case WorkerTimeTicker:
		s.workers[name] = worker{
			Worker:  w,
			typeW:   typ,
			timeRun: ticker,
			run:     make(chan struct{}),
		}
	case WorkerTimeLess:
		s.workers[name] = worker{
			Worker:  w,
			typeW:   typ,
			timeRun: 0,
			run:     make(chan struct{}),
		}
		go func() {
			s.workers[name].run <- struct{}{}
		}()
	default:
		log.Println("please choice worker type")
	}
}

func (s *scheduler) RemoveWorker(name string) {
	delete(s.workers, name)
}

func (s *scheduler) RunWorker(ctx context.Context, name string) {
	s.wg.Add(1)
	go func(w worker) {
		defer s.wg.Done()
		switch w.typeW {
		case WorkerTimeTicker:
			ticker := time.NewTicker(w.timeRun)
			for {
				select {
				case <-ctx.Done():
					w.Stop()
					return
				case <-s.stop:
					w.Stop()
					return
				case <-ticker.C:
					w.Run(ctx)
				}
			}
		case WorkerTimeLess:
			for {
				select {
				case <-ctx.Done():
					w.Stop()
					return
				case <-s.stop:
					w.Stop()
					return
				case <-w.run:
					w.Run(ctx)
				}
			}
		}
	}(s.workers[name])
}

func (s *scheduler) Stop() {
	s.stop <- struct{}{}
	s.wg.Wait()
}
