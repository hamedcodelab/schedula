package schedula

type scheduler struct {
	Worker
}

func NewScheduler() Scheduler {
	return &scheduler{}
}

func (s *scheduler) AddWorker() error {
	return nil
}

func (s *scheduler) RemoveWorker() error {
	return nil
}

func (s *scheduler) Schedule() error {
	return nil
}
