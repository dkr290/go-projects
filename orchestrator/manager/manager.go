package manager

import (
	"fmt"

	"github.com/dkr290/go-projects/orchestrator/task"
	"github.com/golang-collections/collections/queue"
	"github.com/google/uuid"
)

type Manager struct {
	Pending       queue.Queue
	TaskDb        map[string][]task.Task
	EventDb       map[string][]task.TaskEvent
	Workers       []string               //workers
	WorkerTaskMap map[string][]uuid.UUID //jobs that are assigned to each worker.
	TaskWorkerMap map[uuid.UUID]string
}

// SelectWorker responsible for looking at the requirements specified in
// a Task and evaluating the resources available in the pool of workers to see which worker is best suited to run the task.
// select the worker
func (m *Manager) SelectWorker() {

	fmt.Println("I will select appropriate worker")
}

// UpdateTasks keep list of tasks and their states and the machines they run
func (m *Manager) UpdateTasks() {
	fmt.Println("I will update tasks")

}

// SendWork() will send task to the workers
func (m *Manager) SendWork() {
	fmt.Println("I will send work to workers")

}
