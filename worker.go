package schedula

type worker struct {
	ID         string
	Payload    interface{}
	timpestamp int64
}

func NewWorker() Worker {
	return &worker{}
}

func (s *worker) StartWorker(workerID string) error {
	return nil
}

func (s *worker) StopWorker(workerID string) error {
	return nil
}
