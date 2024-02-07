package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	v1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func main() {
	var (
		client     *kubernetes.Clientset
		err        error
		deployFile map[string]string
	)
	if client, err = getClient(); err != nil {
		fmt.Printf("Error client%s", err)
		os.Exit(1)
	}
	ctx := context.Background()
	if deployFile, err = deploy(ctx, client); err != nil {
		fmt.Printf("Error deploy%s", err)
		os.Exit(1)
	}
	fmt.Printf("Deploy %v", deployFile)
}
func getClient() (*kubernetes.Clientset, error) {

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", filepath.Join(homedir.HomeDir(), ".kube", "config"))
	if err != nil {
		panic(err.Error())
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	return clientset, nil
}

func deploy(ctx context.Context, client *kubernetes.Clientset) (map[string]string, error) {
	var deployment *v1.Deployment
	appFile, err := os.ReadFile("app.yml")
	if err != nil {
		return nil, fmt.Errorf("erro file %s", err)
	}
	obj, group, err := scheme.Codecs.UniversalDeserializer().Decode(appFile, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("erro deployment %s", err)
	}
	switch obj.(type) {
	case *v1.Deployment:
		deployment = obj.(*v1.Deployment)
	default:
		return nil, fmt.Errorf("err type %s", group)
	}
	deploymentResponse, err := client.AppsV1().Deployments("default").Create(ctx, deployment, metav1.CreateOptions{})
	if err != nil {
		return nil, fmt.Errorf("erro deployment %s", err)
	}
	return deploymentResponse.Spec.Template.Labels, nil
}
