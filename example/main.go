package main

import (
	"context"
	"github.com/hamedcodelab/schedula"
	"log"
	"time"
)

type w1 struct {
}

func (w *w1) Run(ctx context.Context) {
	log.Println(" Run w1")
}

func (w *w1) Stop() {
	log.Println("Stop w1")
}

type w2 struct {
}

func (w *w2) Run(ctx context.Context) {
	log.Println(" Run w2")
}

func (w *w2) Stop() {
	log.Println("Stop w2")
}

func main() {
	sch := schedula.NewScheduler()
	sch.AddWorker("w1", &w1{}, schedula.SetWorkerType(schedula.ScheduledWorker), schedula.SetTimeTicker(time.Second))
	sch.RunWorker(context.Background(), "w1")

	sch.AddWorker("w2", &w2{}, schedula.SetWorkerType(schedula.EventDrivenWorker))
	sch.RunWorker(context.Background(), "w2")

	time.Sleep(5 * time.Second)
	sch.Stop()
	log.Println("Stop Scheduler")
	//

	time.Sleep(20 * time.Minute)
}
