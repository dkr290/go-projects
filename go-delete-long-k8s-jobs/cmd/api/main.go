package main

import (
	"context"
	"flag"
	"go-cronjobs/pkg/cronjobs"
	"os"
	"path/filepath"
	"time"

	"log"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

var (
	clientset *kubernetes.Clientset
	namespace string
	maxtime   int
	help      = flag.Bool("help", false, "Show help")
)

func init() {

	var kubeconfig *string

	if home := homedir.HomeDir(); home != "" {

		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")

	} else {

		kubeconfig = flag.String("kubeconfig", "  ", "absolute path to the kubeconfig file")

	}
	flag.StringVar(&namespace, "namespace", "", "The namespace to check")
	flag.IntVar(&maxtime, "maxtime", 0, "The max period of time the job runs and after that will be deleted in minutes ")

	flag.Parse()

	if *help {
		flag.Usage()
		os.Exit(0)
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

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)

	if err != nil {

		log.Fatal(err)

	}

	clientset, err = kubernetes.NewForConfig(config)

	if err != nil {

		log.Fatal(err)

	}

}
func main() {

	var jbsStringItems []string

	jobs, err := clientset.BatchV1().Jobs(namespace).List(context.Background(), v1.ListOptions{})
	if err != nil {
		log.Fatalln("Error creating the slientset", err)
	}

	for _, j := range jobs.Items {
		jbsStringItems = append(jbsStringItems, j.Name)
	}

	conf := cronjobs.Config{
		Namespace:   namespace,
		Clientset:   clientset,
		JobName:     jbsStringItems,
		MaxDuration: time.Duration(maxtime) * time.Minute,
	}

	if err := conf.DeleteJobTimstamp(); err != nil {
		log.Fatalln("Error deliting cronjob")

	}

}
