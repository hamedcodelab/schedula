package schedula

import "time"

type Scheduler interface {
	AddWorker(workerName string, ticker time.Duration, fn func() error)
	RemoveWorker(workerName string)
	RunWorker(workerName string)
	StopWorker(workerName string)
	Working()
	//RetryTask(taskID string) error
	//GetPendingTasks() []Task
	//MonitorWorkers() []WorkerStatus
}

type Worker interface {
	StartWorker()
	Run() error
	SetTask(fn func() error) error
	StopWorker()
	ReportStatus() string
	SetStatus(status string)
	//HandleFailure(task Task, err error) error
}
