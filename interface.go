package schedula

import "time"

type Scheduler interface {
	AddWorker(workerName string, ticker time.Duration)
	RemoveWorker(workerName string)
	RunWorker(workerName string) error
	StopWorker(workerName string) error
	Working()
	//RetryTask(taskID string) error
	//GetPendingTasks() []Task
	//MonitorWorkers() []WorkerStatus
}

type Worker interface {
	StartWorker()
	//ExecuteTask(task Task) error
	StopWorker()
	ReportStatus() string
	SetStatus(status string)
	//HandleFailure(task Task, err error) error
}
