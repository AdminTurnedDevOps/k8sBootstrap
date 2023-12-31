package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"

	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/chart/loader"
	"helm.sh/helm/v3/pkg/cli"
)

func main() {
	connectContext()
	installArgo()
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

func installArgo() {

	helmChart := "argo/argo-cd"
	releaseName := "argocd"

	// Call upon the CLI package
	settings := cli.New()

	// Create a new instance of the configuration
	config := new(action.Configuration)

	// Collect local Helm information
	if err := config.Init(settings.RESTClientGetter(), settings.Namespace(), os.Getenv("HELM_DRIVER"), log.Printf); err != nil {
		log.Printf("%+v", err)
	}

	// Create a new instance of the `Install` action, which is similar to running `helm instll`
	client := action.NewInstall(config)

	// Specify Namespace for ArgoCD
	client.CreateNamespace = true
	client.Namespace = "argocd"

	// values := map[string]interface{}{
	//     "redis": map[string]interface{}{
	//         "sentinel": map[string]interface{}{
	//             "masterName": "BigMaster",
	//             "pass":       "random",
	//             "addr":       "localhost",
	//             "port":       "26379",
	//         },
	//     },
	// }

	// Find the Helm Chart. similiar to a `helm add`
	cp, err := client.ChartPathOptions.LocateChart(helmChart, settings)

	if err != nil {
		log.Println(err)
	}

	chart, err := loader.Load(cp)
	if err != nil {
		log.Println(err)
	}

	// Install a helm chart
	client.ReleaseName = releaseName
	results, err := client.Run(chart, nil)
	if err != nil {
		log.Printf("%+v", err)
	}

	fmt.Println(results)
}
