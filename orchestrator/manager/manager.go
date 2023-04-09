package manager

import (
	"github.com/dkr290/go-projects/orchestrator/task"
	"github.com/golang-collections/collections/queue"
)

type Manager struct {
	Pending queue.Queue
	TaskDb  map[string][]task.Task
	EventDb map[string][]task.TaskEvent
}
