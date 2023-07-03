package task

import (
	"context"
	"go-cronjobs/pkg/cronjobs"
	"log"
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type TaskParams struct {
	clientset *kubernetes.Clientset
	namespace string
	maxtime   int
}

func NewTask(c *kubernetes.Clientset, n string, mt int) *TaskParams {
	return &TaskParams{
		clientset: c,
		namespace: n,
		maxtime:   mt,
	}
}

var t *TaskParams

func NewParams(tp *TaskParams) {
	t = tp
}

func Task() {
	var jbsStringItems []string

	jobs, err := t.clientset.BatchV1().Jobs(t.namespace).List(context.Background(), v1.ListOptions{})
	if err != nil {
		log.Fatalln("Error creating the slientset", err)
	}

	for _, j := range jobs.Items {
		jbsStringItems = append(jbsStringItems, j.Name)
	}

	conf := cronjobs.Config{
		Namespace:   t.namespace,
		Clientset:   t.clientset,
		JobName:     jbsStringItems,
		MaxDuration: time.Duration(t.maxtime) * time.Minute,
	}

	if err := conf.DeleteJobTimstamp(); err != nil {
		log.Fatalln("Error deliting cronjob")

	}
}
