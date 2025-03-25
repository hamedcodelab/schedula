package schedula

import (
	"context"
	"time"
)

type Scheduler interface {
	AddWorker(workerName string, typ WorkerSchemaType, ticker time.Duration, w Worker)
	RemoveWorker(workerName string)
	RunWorker(ctx context.Context, name string)
	Stop()
}

type Worker interface {
	Run(ctx context.Context)
	Stop()
}

type worker struct {
	Worker
	typeW   WorkerSchemaType
	timeRun time.Duration
	run     chan struct{}
}

type WorkerSchemaType string

const (
	WorkerTimeTicker WorkerSchemaType = "TimeTicker"
	WorkerTimeLess   WorkerSchemaType = "TimeLess"
)
