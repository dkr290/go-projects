package main

import (
	"fmt"
	"time"

	"github.com/dkr290/go-projects/orchestrator/node"
	"github.com/dkr290/go-projects/orchestrator/task"
	"github.com/dkr290/go-projects/orchestrator/worker"
	"github.com/golang-collections/collections/queue"
	"github.com/google/uuid"
)

func main() {

	t := task.Task{

		ID:     uuid.New(),
		Name:   "Task-1",
		State:  task.Pending,
		Image:  "nginx",
		Memory: 1024,
		Disk:   1,
	}

	te := task.TaskEvent{
		ID:       uuid.New(),
		State:    task.Pending,
		TimStamp: time.Now(),
		Task:     t,
	}

	fmt.Printf("task: %v\n", t)
	fmt.Printf("task event: %v\n", te)

	w := worker.Worker{
		Queue: *queue.New(),
		Db:    make(map[uuid.UUID]task.Task),
	}

	fmt.Printf("worker: %v\n", w)
	w.CollectStats()
	w.RunTask()
	w.StartTask()
	w.StopTask()

	n := node.Node{
		Name:   "Node-1",
		Ip:     "192.168.1.1",
		Cores:  4,
		Memory: 1024,
		Disk:   25,
		Role:   "worker",
	}

	fmt.Printf("node: %v\n", n)
}
