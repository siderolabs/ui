// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package backend

import (
	"context"
	"sync"
	"time"

	"github.com/wailsapp/wails"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/scale/scheme"
	toolscache "k8s.io/client-go/tools/cache"
	"sigs.k8s.io/cluster-api/api/v1alpha3"
	"sigs.k8s.io/controller-runtime/pkg/cache"
)

type Clusters struct {
	config *rest.Config
	log    *wails.CustomLogger

	clusters []*v1alpha3.Cluster

	sync.Mutex
}

func (c *Clusters) WailsInit(runtime *wails.Runtime) error {
	c.log = runtime.Log.New("Clusters")

	ch := make(chan []*v1alpha3.Cluster, 100)

	go func() {
		err := c.watch(ch)
		if err != nil {
			c.log.Errorf("Cluster watch failed: %v", err)
		}
	}()

	go func() {
		// TODO(andrewrynhard): There seems to be a race condition between the
		// frontend and the backend that causes the first events to be dropped by
		// the frontend. Remove this sleep once we have a fix.
		time.Sleep(1 * time.Second)

		for clusters := range ch {
			c.log.Debugf("%+v", clusters)
			runtime.Events.Emit("clusters", clusters)
		}
	}()

	return nil
}

func (c *Clusters) Clusters() []*v1alpha3.Cluster {
	c.Lock()
	defer c.Unlock()

	return c.clusters
}

func (c *Clusters) watch(ch chan []*v1alpha3.Cluster) error {
	s := runtime.NewScheme()
	_ = scheme.AddToScheme(s)
	_ = v1alpha3.AddToScheme(s)

	cache, err := cache.New(c.config, cache.Options{Scheme: s})
	if err != nil {
		return err
	}

	informer, err := cache.GetInformer(context.TODO(), &v1alpha3.Cluster{})
	if err != nil {
		return err
	}

	informer.AddEventHandler(toolscache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			c.Lock()
			defer c.Unlock()

			cluster := obj.(*v1alpha3.Cluster)

			// TODO(andrewrynhard): Remove this once we figure out why these
			// fields are causing the JSON decoder on the frontend to fail.
			cluster.ManagedFields = nil
			cluster.Annotations = nil

			c.clusters = append(c.clusters, cluster)

			ch <- c.clusters
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			c.Lock()
			defer c.Unlock()

			cluster := newObj.(*v1alpha3.Cluster)

			for i, old := range c.clusters {
				if old.UID == cluster.UID {
					// TODO(andrewrynhard): Remove this once we figure out why these
					// fields are causing the JSON decoder on the frontend to fail.
					cluster.ManagedFields = nil
					cluster.Annotations = nil

					c.clusters[i] = cluster

					break
				}
			}

			ch <- c.clusters
		},
		DeleteFunc: func(obj interface{}) {
			c.Lock()
			defer c.Unlock()

			cluster := obj.(*v1alpha3.Cluster)

			for i, old := range c.clusters {
				if old.UID == cluster.UID {
					c.clusters[i] = c.clusters[len(c.clusters)-1]
					c.clusters[len(c.clusters)-1] = nil
					c.clusters = c.clusters[:len(c.clusters)-1]

					break
				}
			}

			ch <- c.clusters
		},
	})

	stopCh := make(chan struct{})
	defer close(stopCh)

	go cache.Start(stopCh)

	if ok := cache.WaitForCacheSync(stopCh); ok {
		c.log.Debug("Cluster cache synced.")
	}

	<-stopCh

	return nil
}
