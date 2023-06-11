package main

import (
	"flag"
	"fmt"
	"path/filepath"

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

	fmt.Println(*kubeConfig)

}
