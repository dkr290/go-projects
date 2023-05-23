package main

import (
	"fmt"
	"log"
	"time"

	"github.com/dkr290/go-projects/orchestrator/manager"
	"github.com/dkr290/go-projects/orchestrator/node"
	"github.com/dkr290/go-projects/orchestrator/task"
	"github.com/dkr290/go-projects/orchestrator/worker"
	"github.com/docker/docker/client"
	"github.com/golang-collections/collections/queue"
	"github.com/google/uuid"
)

func main() {

	t := task.Task{
		ID:     uuid.New(),
		Name:   "Task-1",
		State:  task.Pending,
		Image:  "busybox:latest",
		Memory: 1024,
		Disk:   1,
	}

	te := task.TaskEvent{
		ID:        uuid.New(),
		State:     task.Pending,
		Timestamp: time.Now(),
		Task:      t,
	}
	fmt.Printf("task: %v\n", t)

	fmt.Printf("task event: %v\n", te)

	w := worker.Worker{
		Name:  "Worker1",
		Queue: *queue.New(),
		Db:    make(map[uuid.UUID]task.Task),
	}

	fmt.Printf("worker: %v\n", w)

	w.CollectStats()
	w.RunTask()
	w.StartTask()
	w.StopTask()

	m := manager.Manager{
		Pending: *queue.New(),
		TaskDb:  make(map[string][]task.Task),
		EventDb: make(map[string][]task.TaskEvent),
		Workers: []string{w.Name},
	}
	fmt.Printf("manager: %v\n", m)
	m.SelectWorker()
	m.UpdateTasks()
	m.SendWork()

	n := node.Node{
		Name:   "Node-1",
		Ip:     "192.168.122.186",
		Cores:  2,
		Memory: 1024,
		Disk:   25,
		Role:   "worker",
	}

	fmt.Printf("node: %v\n", n)

}

func createContainer() (*task.Container, *task.ContainerResult) {
	c := task.Config{
		Name:  "test-container-1",
		Image: "postgres:14",
		Env:   []string{"POSTGRES_USER=user", "POSTGRES_PASSWORD=password1"},
	}

	newContainer, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}

	d := task.New(newContainer, c)
	result := d.Run()
	if result.Error != nil {
		log.Printf("%v\n", result.Error)
		return nil, nil

	}

	fmt.Printf("Container %s is running with config %v\n", result.ContainerId, c)

	return &d, &result

}
