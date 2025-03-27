package schedula

import (
	"context"
	"sync"
	"time"
)

type scheduler struct {
	workers map[string]*worker
	mu      sync.Mutex
	stop    chan struct{}
	wg      sync.WaitGroup
	once    sync.Once
}

func NewScheduler() Scheduler {
	return &scheduler{
		workers: make(map[string]*worker),
		stop:    make(chan struct{}),
		wg:      sync.WaitGroup{},
	}
}

type Option func(w *worker)

func SetTimeTicker(ticker time.Duration) Option {
	return func(w *worker) {
		w.timeRun = ticker
	}
}

func SetWorkerType(typ WorkerType) Option {
	return func(w *worker) {
		w.typeW = typ
	}
}

func (s *scheduler) AddWorker(name string, w Worker, opt ...Option) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.workers[name]; exists {
		return
	}

	s.workers[name] = &worker{
		Worker: w,
		run:    make(chan struct{}),
	}
	for _, fn := range opt {
		fn(s.workers[name])
	}

	if s.workers[name].typeW == "" {
		s.workers[name].typeW = ScheduledWorker
	}

	if s.workers[name].timeRun == 0 {
		s.workers[name].timeRun = time.Second
	}

	if s.workers[name].typeW == EventDrivenWorker {
		s.workers[name].run = make(chan struct{})
		go func() {
			s.workers[name].run <- struct{}{}
		}()
	}

	return
}

func (s *scheduler) RunWorker(ctx context.Context, name string) {
	s.mu.Lock()
	w, exists := s.workers[name]
	s.mu.Unlock()

	if !exists {
		return
	}

	s.wg.Add(1)
	go func(w *worker) {
		defer s.wg.Done()
		switch w.typeW {
		case ScheduledWorker:
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
		case EventDrivenWorker:
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
	}(w)
}

func (s *scheduler) Stop() {
	////s.stop <- struct{}{}
	//close(s.stop)
	//s.wg.Wait()
	s.once.Do(func() {
		close(s.stop)
		s.wg.Wait()
	})
}
