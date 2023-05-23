package worker

import (
	"fmt"

	"github.com/dkr290/go-projects/orchestrator/task"
	"github.com/golang-collections/collections/queue"
	"github.com/google/uuid"
)

type Worker struct {
	Name      string
	Queue     queue.Queue
	Db        map[uuid.UUID]task.Task
	TaskCount int
}

func (w *Worker) RunTask() {
	fmt.Println("Running the task")
}

func (w *Worker) CollectStats() {
	fmt.Println("Collection the stats")
}

func (w *Worker) StartTask() {
	fmt.Println("starting the task")
}

func (w *Worker) StopTask() {
	fmt.Println("I will stop the taks")
}
