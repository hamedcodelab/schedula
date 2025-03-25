package main

import (
	"fmt"
	"github.com/hamedcodelab/schedula"
	"time"
)

func main() {
	sch := schedula.NewScheduler()
	sch.AddWorker("w1", time.Second, func() error {
		fmt.Println(fmt.Sprintf("Worker %s Running", "w1"))
		return nil
	})
	sch.AddWorker("w2", time.Second*30, func() error {
		fmt.Println(fmt.Sprintf("Worker %s Running", "w2"))
		return nil
	})
	sch.RunWorker("w1")
	sch.RunWorker("w2")
	go sch.Working()

	time.Sleep(time.Minute * 30)
}
