package main

import (
	"github.com/hamedcodelab/schedula"
	"time"
)

func main() {
	sch := schedula.NewScheduler()
	sch.AddWorker("w1", time.Second)
	sch.AddWorker("w2", time.Minute)
	sch.RunWorker("w1")
	sch.RunWorker("w2")
	go sch.Working()

	time.Sleep(time.Minute * 30)
}
