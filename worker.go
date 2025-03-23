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

func NewWorker(name string, ticker time.Duration) Worker {
	return &worker{
		name:   name,
		ticker: ticker,
		status: "New",
		stop:   make(chan bool),
	}
}

func (s *worker) StartWorker() {
	s.status = "Running"
	for {
		select {
		case <-s.stop:
			fmt.Println(fmt.Sprintf("Worker %s cancelled"), s.name)
			return
		default:
			fmt.Println(fmt.Sprintf("Worker %s Working"), s.name)
			time.Sleep(s.ticker)
		}
	}

}

func (s *worker) StopWorker() {
	s.stop <- true
	fmt.Println(fmt.Sprintf("Worker %s Stop"), s.name)
}

func (s *worker) ReportStatus(workerName string) string {
	return s.status
}
