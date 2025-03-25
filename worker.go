package schedula

import (
	"time"
)

// context
// ticker
// ticker := time.NewTicker(w.d)
// run in worker

type worker struct {
	Worker
	TimeRun time.Duration
}
