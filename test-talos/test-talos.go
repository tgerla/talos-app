package main

import (
	//	"bytes"
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	//	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	//	k8sjson "k8s.io/apimachinery/pkg/runtime/serializer/json"
	"k8s.io/client-go/kubernetes"
	//	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/tools/clientcmd"

	"github.com/pkg/errors"
	"github.com/talos-systems/talos/cmd/osctl/pkg/client"
	"github.com/talos-systems/talos/pkg/constants"
)

func newTalosClient(target string) (*client.Client, error) {
	home := homeDir()
	_, creds, err := client.NewClientTargetAndCredentialsFromConfig(filepath.Join(home, ".talos", "config"))
	if err != nil {
		return nil, errors.Wrap(err, "error getting client credentials")
	}
	log.Printf("Target: %s", target)
	c, err := client.NewClient(creds, target, constants.OsdPort)
	if err != nil {
		return nil, errors.Wrap(err, "error constructing client")
	}

	return c, nil
}

func newKubernetesClient() (*kubernetes.Clientset, error) {
	var kubeconfig *string
	if home := homeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	return clientset, err
}

// from https://github.com/kubernetes/client-go/blob/master/examples/out-of-cluster-client-configuration/main.go
func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}

func main() {
	var clientset *kubernetes.Clientset

	clientset, err := newKubernetesClient()
	if err != nil {
		panic(err.Error())
	}

	nodes, err := clientset.CoreV1().Nodes().List(metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	log.Printf("There are %d nodes in the cluster\n", len(nodes.Items))

	for _, n := range nodes.Items {
		log.Printf("Node: %s (%s)\n", n.Status.Addresses[0].Address, n.Status.Addresses[0].Type)

		c, err := newTalosClient(n.Status.Addresses[0].Address)
		if err != nil {
			log.Printf("%+v", err)
		}
		reply, err := c.ServiceList(context.Background())
		if err != nil {
			log.Printf("error getting services: %+v", err)
		}
		for _, n := range reply.Response {
			for _, s := range n.Services {
				fmt.Printf("Service name %s (status: %s)\n", s.GetId(), s.GetState())
			}
			//			fmt.Printf("Service: %s (%s)\n", s.GetId(), s.GetState())
		}

		fmt.Println("------------------------------------------------")
	}

}
