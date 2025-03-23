package schedula

type Scheduler interface {
	AddWorker() error
	RemoveWorker() error
	Schedule() error
	//RetryTask(taskID string) error
	//CancelTask(taskID string) error
	//GetPendingTasks() []Task
	//MonitorWorkers() []WorkerStatus
}

type Worker interface {
	StartWorker(workerID string) error
	//ExecuteTask(task Task) error
	StopWorker(workerID string) error
	//ReportStatus(workerID string) string
	//HandleFailure(task Task, err error) error
}
