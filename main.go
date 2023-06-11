package main

import (
	"flag"
	"fmt"
	"log"
	"path/filepath"

	"github.com/opentracing/opentracing-go/log"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func main() {
	connectContext()
}

func connectContext() {
	home := homedir.HomeDir()

	var kubeConfig *string

	if home == "" {
		fmt.Print("Enter Your Kubeconfig Context: ")
		fmt.Scan(&kubeConfig)

	} else {
		kubeConfig = flag.String("kubeConfig", filepath.Join(home, ".kube", "config"), "")
	}

	kubeConfigPath, err := clientcmd.BuildConfigFromFlags("", *kubeConfig)
	if err != nil {
		log.Fatalln(err)
		fmt.Print("Enter Your Kubeconfig Context: ")
		fmt.Scan(&kubeConfig)
	}

	clientConfig, err := dynamic.NewForConfig(kubeConfigPath)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(clientConfig)

}
