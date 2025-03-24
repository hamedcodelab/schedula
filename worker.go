package schedula

import (
	"fmt"
	"time"
)

type worker struct {
	name string
	//payload interface{}
	ticker time.Duration
	status string
	stop   chan bool
}

func newWorker(name string, status string, ticker time.Duration) Worker {
	return &worker{
		name:   name,
		ticker: ticker,
		status: status,
		stop:   make(chan bool),
	}
}

func (s *worker) StartWorker() {
	for {
		select {
		case <-s.stop:
			fmt.Println(fmt.Sprintf("Worker %s cancelled", s.name))
			return
		default:
			fmt.Println(fmt.Sprintf("Worker %s Working", s.name))
			time.Sleep(s.ticker)
			fmt.Println(fmt.Sprintf("Worker %s Waiting Finish", s.name))
		}
	}

}

func (s *worker) StopWorker() {
	s.stop <- true
	fmt.Println(fmt.Sprintf("Worker %s Stop"), s.name)
}

func (s *worker) ReportStatus() string {
	return s.status
}

func (s *worker) SetStatus(status string) {
	s.status = status
}
