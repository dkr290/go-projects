package main

import (
	"flag"
	"go-cronjobs/pkg/task"
	"log"
	"os"
	"strconv"

	"github.com/jasonlvhit/gocron"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

var (
	clientset *kubernetes.Clientset
	namespace string
	maxtime   int
)

func init() {

	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}

	clientset, err = kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	namespace = os.Getenv("KUBERNETES_NAMEPSACE")
	maxtime, err = strconv.Atoi(os.Getenv("KUBERNETES_MAXTIME"))
	if err != nil {
		log.Println("Something get wrong by conversion to maxtime please ensure KUBERNETES_MAXTIME variable is number (minutes like for example 10 min is  10)")
		log.Fatalln(err)
	}

	if namespace == "" {
		log.Println("The namespace must be set")
		flag.Usage()
		os.Exit(0)
	}
	if maxtime == 0 {

		log.Println("The max duration needs to be set in minutes just as for example: -maxtime 5")
		flag.Usage()
		os.Exit(0)
	}

}

func main() {

	params := task.NewTask(clientset, namespace, maxtime)
	task.NewParams(params)

	s := gocron.NewScheduler()

	s.Every(20).Minutes().Do(task.Task)

	<-s.Start()
}
