package worker

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/dkr290/go-projects/orchestrator/task"
	"github.com/docker/docker/client"
	"github.com/golang-collections/collections/queue"
	"github.com/google/uuid"
)

type Worker struct {
	Name      string
	Queue     queue.Queue
	Db        map[uuid.UUID]*task.Task
	TaskCount int
}

func (w *Worker) RunTask() task.ContainerResult {
	t := w.Queue.Dequeue()
	if t == nil {
		log.Println("No tasks in the queue")
		return task.ContainerResult{Error: nil}
	}

	taskQueued := t.(task.Task)
	taskPersisted := w.Db[taskQueued.ID]
	if taskPersisted == nil {
		taskPersisted = &taskQueued
		w.Db[taskQueued.ID] = &taskQueued
	}

	var result task.ContainerResult
	if task.ValidStateTransition(
		taskPersisted.State, taskQueued.State) {
		switch taskQueued.State {
		case task.Scheduled:
			result = w.StartTask(taskQueued)
		case task.Completed:
			result = w.StopTask(taskQueued)
		default:
			result.Error = errors.New("we should not get here")
		}
	} else {
		err := fmt.Errorf("invalid transition from %v to %v", taskPersisted.State, taskQueued.State)
		result.Error = err
	}

	return result
}

func (w *Worker) CollectStats() {
	fmt.Println("Collection the stats")
}

func (w *Worker) StartTask(t task.Task) task.ContainerResult {
	fmt.Println("starting the task")
	t.StartTime = time.Now().UTC()
	config := task.NewConfig(&t)
	newContainer, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}

	d := task.New(newContainer, config)
	result := d.Run()
	if result.Error != nil {
		log.Printf("Error running the task %v:%v\n", t.ID, result.Error)
		t.State = task.Failed
		w.Db[t.ID] = &t
		return result

	}

	t.ContainerId = result.ContainerId
	t.State = task.Running
	w.Db[t.ID] = &t
	return result

}

func (w *Worker) StopTask(t task.Task) task.ContainerResult {
	fmt.Println("I will stop the taks")
	config := task.NewConfig(&t)

	newContainer, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}

	d := task.New(newContainer, config)
	result := d.Stop(t.ContainerId)
	if result.Error != nil {
		log.Printf("Error stopping container  %v:%v", t.ContainerId, result.Error)

	}
	t.FinishTime = time.Now().UTC()
	t.State = task.Completed
	w.Db[t.ID] = &t
	log.Printf("Stopped and removed container %v for task %v", t.ContainerId, t.ID)
	return result
}
