package schedula

import "time"

type scheduler struct {
	workers map[string]Worker
}

func NewScheduler() Scheduler {
	return &scheduler{}
}

func (s *scheduler) AddWorker(name string, ticker time.Duration) {
	newW := NewWorker(name, ticker)
	s.workers[name] = newW
}

func (s *scheduler) RemoveWorker(name string) {
	if s.workers == nil {
		return
	}

	if s.workers[name] == nil {
		return
	}
	if s.workers[name].ReportStatus(name) == "Running" {
		return
	}
	delete(s.workers, name)
}

func (s *scheduler) RunWorker(name string) error {
	go s.workers[name].StartWorker()
	return nil
}

func (s *scheduler) StopWorker(name string) error {
	s.workers[name].StopWorker()
	return nil
}
