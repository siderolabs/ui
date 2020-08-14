// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package backend

import (
	"os"
	"path/filepath"

	"github.com/talos-systems/ui/pkg/backend/common"
	"github.com/talos-systems/ui/pkg/backend/talos"
	"k8s.io/client-go/tools/clientcmd"
)

type Backend struct {
	Kubernetes *Kubernetes
	Talos      *Talos
}

type Kubernetes struct {
	Clusters *talos.Clusters
	V1       *common.CAPI
}

type Talos struct{}

func NewBackend() (*Backend, error) {
	var kubeconfig string

	if env, ok := os.LookupEnv("KUBECONFIG"); ok {
		kubeconfig = env
	} else {
		kubeconfig = filepath.Join(os.Getenv("HOME"), ".kube", "config")
	}

	config, err := clientcmd.LoadFromFile(kubeconfig)
	if err != nil {
		return nil, err
	}

	clusters := talos.NewClusters(config)
	v1 := common.NewCAPI(common.GetScheme())

	return &Backend{
		Kubernetes: &Kubernetes{
			Clusters: clusters,
			V1:       v1,
		},
	}, nil
}
