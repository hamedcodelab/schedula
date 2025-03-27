package schedula

import (
	"context"
	"time"
)

type Scheduler interface {
	AddWorker(workerName string, w Worker, opt ...Option)
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
