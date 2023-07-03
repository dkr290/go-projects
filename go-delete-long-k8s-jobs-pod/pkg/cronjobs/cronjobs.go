package cronjobs

import (
	"context"
	"log"
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type Config struct {
	JobName     []string
	Clientset   *kubernetes.Clientset
	Namespace   string
	MaxDuration time.Duration
}

func (c *Config) DeleteJobTimstamp() error {

	jobClient := c.Clientset.BatchV1().Jobs(c.Namespace)

	for _, jobStr := range c.JobName {

		job, err := jobClient.Get(context.Background(), jobStr, v1.GetOptions{})
		if err != nil {
			return err
		}

		jobCreation := job.CreationTimestamp
		timeNow := time.Now()

		lognTimeForJob := jobCreation.Add(c.MaxDuration)
		pp := v1.DeletePropagationBackground

		if timeNow.After(lognTimeForJob) && job.Status.Active == 1 {

			err := jobClient.Delete(context.Background(), jobStr, v1.DeleteOptions{PropagationPolicy: &pp})
			if err != nil {
				return err
			}

			log.Printf("The cronjob %s started at %v was deleted", jobStr, job.Status.StartTime)
		}

	}

	return nil

}
