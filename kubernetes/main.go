package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"time"

	v1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
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
		fmt.Printf("Error client%s\n", err)
		os.Exit(1)
	}
	ctx := context.Background()
	if deployFile, err = deploy(ctx, client); err != nil {
		fmt.Printf("Error deploy%s\n", err)
		os.Exit(1)
	}
	if err = waitForPods(ctx, client, deployFile); err != nil {
		fmt.Printf("Error deploy%s\n", err)
		os.Exit(1)
	}
	fmt.Printf("Deploy %v\n", deployFile["app"])
}
func getClient() (*kubernetes.Clientset, error) {

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", filepath.Join(homedir.HomeDir(), ".kube", "config")) // carrega a config do cluster
	if err != nil {
		panic(err.Error())
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config) // inicializa as config do clust
	if err != nil {
		panic(err.Error())
	}
	return clientset, nil
}

func deploy(ctx context.Context, client *kubernetes.Clientset) (map[string]string, error) {
	var deployment *v1.Deployment
	appFile, err := os.ReadFile("app.yml")
	if err != nil {
		return nil, fmt.Errorf("erro file %s\n", err)
	}

	obj, group, err := scheme.Codecs.UniversalDeserializer().Decode(appFile, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("erro deployment %s\n", err)
	}
	fmt.Printf("OBJECT %#v\n", obj.(*v1.Deployment).ObjectMeta)
	switch obj.(type) {
	case *v1.Deployment:
		deployment = obj.(*v1.Deployment)
	default:
		return nil, fmt.Errorf("err type %s", group)
	}
	_, err = client.AppsV1().Deployments("default").Get(ctx, deployment.Name, metav1.GetOptions{})
	fmt.Printf("error %v\n", errors.IsNotFound(err))
	if err != nil && errors.IsNotFound(err) { // se o erro for = NotFound
		deploymentResponse, err := client.AppsV1().Deployments("default").Create(ctx, deployment, metav1.CreateOptions{})
		if err != nil {
			return nil, fmt.Errorf("erro deployment %s\n", err)
		}
		return deploymentResponse.Spec.Template.Labels, nil
	} else if err != nil && !errors.IsNotFound(err) { // se existir um deployment
		return nil, fmt.Errorf("get error %s\n", err)
	}
	deploymentResponse, err := client.AppsV1().Deployments("default").Update(ctx, deployment, metav1.UpdateOptions{})
	if err != nil {
		return nil, fmt.Errorf("erro deployment %s\n", err)
	}
	return deploymentResponse.Spec.Template.Labels, nil
}

func waitForPods(ctx context.Context, client *kubernetes.Clientset, deploymentLabels map[string]string) error {
	for {
		parshedLabels, err := labels.ValidatedSelectorFromSet(deploymentLabels)
		if err != nil {
			return fmt.Errorf("pods list error %s\n", err)
		}
		podList, err := client.CoreV1().Pods("default").List(ctx, metav1.ListOptions{
			LabelSelector: parshedLabels.String(),
		})
		if err != nil {
			return fmt.Errorf("pods list error %s\n", err)
		}
		podsNumberRunnig := 0
		for _, pod := range podList.Items {
			if pod.Status.Phase != "running" {
				podsNumberRunnig++
			}
		}
		fmt.Printf("wait for pods, ready (running %d | %d)\n", podsNumberRunnig, len(podList.Items))
		if podsNumberRunnig > 0 && podsNumberRunnig == len(podList.Items) {
			break
		}

		time.Sleep(1 * time.Second)
	}
	return nil
}
