package schedula

import (
	"context"
	"time"
)

type Scheduler interface {
	// worker handler
	AddWorker(workerName string, ticker time.Duration, w Worker)
	RemoveWorker(workerName string)
	RunWorker(ctx context.Context, name string)
	//StopWorker(workerName string)
	StopScheduler()
}

type Worker interface {
	Run(ctx context.Context) error
	Stop() error
}
