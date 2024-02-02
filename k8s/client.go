package k8s

import (
	"fmt"
	"path/filepath"
	"playground/system"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

// Get kubeconfig from home directory
var bondeeKubeconfig = filepath.Join(system.HomeDir(), ".kube", "bondee_config")
var defaultKubeconfig = filepath.Join(system.HomeDir(), ".kube", "config")
var k8sClient *kubernetes.Clientset

func Init() {
	config, err := clientcmd.BuildConfigFromFlags("", bondeeKubeconfig)
	if err != nil {
		fmt.Printf("Failed to load bondee config: %v, try default config \n", err)
		config, err = clientcmd.BuildConfigFromFlags("", bondeeKubeconfig)
		if err != nil {
			panic(err.Error())
		}
	}

	k8sClient, err = kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
}
