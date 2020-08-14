package main

import (
	"context"
	"fmt"
	"time"

	"github.com/airbloc/logger"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

var log = logger.New("rollerderby")

func main() {
	// creates the in-cluster config
	config, err := rest.InClusterConfig()
	if err != nil {
		log.Fatal("{} (Is it running on inside kubernetes?)", err.Error())
	}
	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal("Failed to initiate kubernetes client", err)
	}

	for {
		// get pods in all the namespaces by omitting namespace
		// Or specify namespace to get pods in particular namespace
		pods, err := clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			log.Fatal("Failed to list pods in cluster", err)
		}
		fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))

		// Examples for error handling:
		// - Use helper functions e.g. errors.IsNotFound()
		// - And/or cast to StatusError and use its properties like e.g. ErrStatus.Message
		_, err = clientset.CoreV1().Pods("default").
			Get(context.TODO(), "example-xxxxx", metav1.GetOptions{})

		if errors.IsNotFound(err) {
			log.Info("Pod example-xxxxx not found in default namespace")
		} else if statusError, isStatus := err.(*errors.StatusError); isStatus {
			log.Error("Error getting pod: {}", statusError.ErrStatus.Message)
		} else if err != nil {
			log.Error("Failed to resolve pod status", err)
		} else {
			log.Info("Found example-xxxxx pod in default namespace")
		}
		time.Sleep(10 * time.Second)
	}
}
