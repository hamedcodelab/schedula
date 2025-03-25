package schedula

import (
	"context"
	"time"
)

type Scheduler interface {
	AddWorker(workerName string, typ WorkerType, ticker time.Duration, w Worker)
	RunWorker(ctx context.Context, name string)
	Stop()
}

type Worker interface {
	Run(ctx context.Context)
	Stop()
}

type worker struct {
	Worker
	typeW   WorkerType
	timeRun time.Duration
	run     chan struct{}
}

type WorkerType string

const (
	ScheduledWorker   WorkerType = "TimeTicker"
	EventDrivenWorker WorkerType = "TimeLess"
)
