package main

import (
	"context"
	"github.com/hamedcodelab/schedula"
	"log"
	"time"
)

type w1 struct {
}

func (w *w1) Run(ctx context.Context) error {
	log.Println(" Run w1")
	return nil
}

func (w *w1) Stop() error {
	log.Println("Stop w1")
	return nil
}

func main() {
	sch := schedula.NewScheduler()
	sch.AddWorker("w1", time.Second, &w1{})
	sch.RunWorker(context.Background(), "w1")
	time.Sleep(20 * time.Minute)
}
