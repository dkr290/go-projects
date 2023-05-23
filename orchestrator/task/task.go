package task

import (
<<<<<<< HEAD
	"context"
	"fmt"
	"io"
	"log"
	"os"
=======
>>>>>>> bad74e0 (ds)
	"time"

	"github.com/docker/go-connections/nat"
	"github.com/google/uuid"
)

type State int

var containerclient *client.Client

const (
	Pending State = iota
	Scheduled
	Running
	Completed
	Failed
)

type Task struct {
	ID            uuid.UUID
<<<<<<< HEAD
	ContainerId   string
=======
>>>>>>> bad74e0 (ds)
	Name          string
	State         State
	Image         string
	Memory        int
	Disk          int
	ExposedPorts  nat.PortSet
	PortBindings  map[string]string
	RestartPolicy string
	StartTime     time.Time
	FinishTime    time.Time
}
<<<<<<< HEAD

type TaskEvent struct {
	ID        uuid.UUID
	State     State
	Timestamp time.Time
	Task      Task
}

type Config struct {
	Name                  string
	AttachStdin           bool
	AttachStdout          bool
	AttachStderr          bool
	Image                 string
	Memory                int64
	Disk                  int64
	Env                   []string
	RestartPolicy         string
	RestartPolicyMaxRetry int
}

type Container struct {
	Config      Config
	ContainerId string
}

type ContainerResult struct {
	Error       error
	Action      string
	ContainerId string
	Result      string
}

func New(c *client.Client, conf Config) Container {
	containerclient = c
	return Container{
		Config: conf,
	}

}

func NewConfig(t *Task) Config {
	return Config{
		Name:          t.Name,
		Image:         t.Image,
		Memory:        int64(t.Memory),
		Disk:          int64(t.Disk),
		RestartPolicy: t.RestartPolicy,
	}
}

func (c *Container) Run() ContainerResult {

	ctx := context.Background()
	reader, err := containerclient.ImagePull(ctx, c.Config.Image, types.ImagePullOptions{})
	if err != nil {
		log.Printf("Error pulling image %s:%v\n", c.Config.Image, err)
		return ContainerResult{
			Error: err,
		}

	}
	io.Copy(os.Stdout, reader)

	restartPolicy := container.RestartPolicy{
		Name:              c.Config.RestartPolicy,
		MaximumRetryCount: c.Config.RestartPolicyMaxRetry,
	}

	resourceTypes := container.Resources{
		Memory: c.Config.Memory,
	}

	containerConfig := container.Config{
		Image: c.Config.Image,
		Env:   c.Config.Env,
	}

	containerHostConfig := container.HostConfig{
		RestartPolicy:   restartPolicy,
		Resources:       resourceTypes,
		PublishAllPorts: true,
	}

	resp, err := containerclient.ContainerCreate(ctx, &containerConfig, &containerHostConfig, nil, nil, c.Config.Name)
	if err != nil {
		log.Printf("Error creating container using image %s: %v\n", c.Config.Image, err)
		return ContainerResult{Error: err}
	}

	if err := containerclient.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		log.Printf("Error starting container %s:%v\n", resp.ID, err)
		return ContainerResult{
			Error: err,
		}
	}

	c.ContainerId = resp.ID

	out, err := containerclient.ContainerLogs(ctx, resp.ID, types.ContainerLogsOptions{
		ShowStdout: true,
		ShowStderr: true,
	})

	if err != nil {
		log.Printf("Error getting logs for container %s:%v\n", resp.ID, err)
		return ContainerResult{
			Error: err,
		}
	}

	stdcopy.StdCopy(os.Stdout, os.Stderr, out)

	return ContainerResult{
		ContainerId: resp.ID,
		Action:      "start",
		Result:      "sucess",
	}

}

func (c *Container) Stop(ContainerId string) ContainerResult {
	ctx := context.Background()
	log.Println("Attempting to stop the container with id", ContainerId)
	err := containerclient.ContainerStop(ctx, ContainerId, container.StopOptions{})
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	err = containerclient.ContainerRemove(ctx, ContainerId, types.ContainerRemoveOptions{
		RemoveVolumes: true,
		RemoveLinks:   false,
		Force:         false,
	})
	if err != nil {
		panic(err)
	}

	return ContainerResult{
		Action:      "stop",
		ContainerId: ContainerId,
		Error:       nil,
		Result:      "success",
	}

}
=======
>>>>>>> bad74e0 (ds)
