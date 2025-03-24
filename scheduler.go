package schedula

import "time"

type scheduler struct {
	workers map[string]Worker
}

func NewScheduler() Scheduler {
	return &scheduler{
		workers: make(map[string]Worker),
	}
}

func (s *scheduler) AddWorker(name string, ticker time.Duration) {
	newW := newWorker(name, "New", ticker)
	s.workers[name] = newW
}

func (s *scheduler) RemoveWorker(name string) {
	if s.workers == nil {
		return
	}

	if s.workers[name] == nil {
		return
	}
	if s.workers[name].ReportStatus() == "Running" {
		return
	}
	delete(s.workers, name)
}

func (s *scheduler) RunWorker(name string) {
	s.workers[name].SetStatus("Running")
}

func (s *scheduler) StopWorker(name string) {
	s.workers[name].StopWorker()
	s.workers[name].SetStatus("Stop")
}

func (s *scheduler) Working() {
	for _, w := range s.workers {
		if w.ReportStatus() == "Running" {
			go w.StartWorker()
		}
	}
}
