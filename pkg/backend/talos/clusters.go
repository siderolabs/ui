// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package talos

import (
	"sync"

	"github.com/wailsapp/wails"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
)

func NewClusters(config *clientcmdapi.Config) *Clusters {
	return &Clusters{
		config: config,
	}
}

type Cluster struct {
	*clientcmdapi.Cluster
	*clientcmdapi.Context
	ContextName string `json:"contextName"`
}

type Clusters struct {
	config *clientcmdapi.Config
	log    *wails.CustomLogger

	sync.Mutex
}

func (c *Clusters) WailsInit(runtime *wails.Runtime) error {
	c.log = runtime.Log.New("Clusters")
	return nil
}

func (c *Clusters) Clusters() map[string]*Cluster {
	c.Lock()
	defer c.Unlock()

	clusters := map[string]*Cluster{}

	for name, context := range c.config.Contexts {
		cluster, ok := c.config.Clusters[context.Cluster]
		if !ok {
			continue
		}

		clusters[context.Cluster] = &Cluster{
			Cluster:     cluster,
			Context:     context,
			ContextName: name,
		}
	}

	return clusters
}
