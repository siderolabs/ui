package backend

import (
	"os"
	"path/filepath"

	"k8s.io/client-go/tools/clientcmd"
)

type Backend struct {
	Kubernetes *Kubernetes
	Talos      *Talos
}

type Kubernetes struct {
	Clusters *Clusters
	Machines *Machines
}

type Talos struct{}

func NewBackend() (*Backend, error) {
	var kubeconfig string

	if env, ok := os.LookupEnv("KUBECONFIG"); ok {
		kubeconfig = env
	} else {
		kubeconfig = filepath.Join(os.Getenv("HOME"), ".kube", "config")
	}

	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		return nil, err
	}

	clusters := &Clusters{config: config}
	machines := &Machines{config: config}

	return &Backend{
		Kubernetes: &Kubernetes{
			Clusters: clusters,
			Machines: machines,
		},
	}, nil
}
