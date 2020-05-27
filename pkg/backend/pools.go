// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package backend

import (
	"context"
	"sync"
	"time"

	"github.com/talos-systems/talos-controller-manager/api/v1alpha1"
	"github.com/wailsapp/wails"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/scale/scheme"
	toolscache "k8s.io/client-go/tools/cache"
	"sigs.k8s.io/controller-runtime/pkg/cache"
)

type Pools struct {
	config *rest.Config
	log    *wails.CustomLogger

	pools []*v1alpha1.Pool

	sync.Mutex
}

func (c *Pools) WailsInit(runtime *wails.Runtime) error {
	c.log = runtime.Log.New("Pools")

	ch := make(chan []*v1alpha1.Pool, 100)

	go func() {
		err := c.watch(ch)
		if err != nil {
			c.log.Errorf("Pool watch failed: %v", err)
		}
	}()

	go func() {
		// TODO(andrewrynhard): There seems to be a race condition between the
		// frontend and the backend that causes the first events to be dropped by
		// the frontend. Remove this sleep once we have a fix.
		time.Sleep(1 * time.Second)

		for pools := range ch {
			c.log.Debugf("%+v", pools)
			runtime.Events.Emit("pools", pools)
		}
	}()

	return nil
}

func (c *Pools) Pools() []*v1alpha1.Pool {
	c.Lock()
	defer c.Unlock()

	return c.pools
}

func (c *Pools) watch(ch chan []*v1alpha1.Pool) error {
	s := runtime.NewScheme()
	_ = scheme.AddToScheme(s)
	_ = v1alpha1.AddToScheme(s)

	cache, err := cache.New(c.config, cache.Options{Scheme: s})
	if err != nil {
		return err
	}

	informer, err := cache.GetInformer(context.TODO(), &v1alpha1.Pool{})
	if err != nil {
		return err
	}

	informer.AddEventHandler(toolscache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			c.Lock()
			defer c.Unlock()

			pool := obj.(*v1alpha1.Pool)

			c.pools = append(c.pools, pool)

			ch <- c.pools
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			c.Lock()
			defer c.Unlock()

			pool := newObj.(*v1alpha1.Pool)

			for i, old := range c.pools {
				if old.UID == pool.UID {
					c.pools[i] = pool

					break
				}
			}

			ch <- c.pools
		},
		DeleteFunc: func(obj interface{}) {
			c.Lock()
			defer c.Unlock()

			pool := obj.(*v1alpha1.Pool)

			for i, old := range c.pools {
				if old.UID == pool.UID {
					c.pools[i] = c.pools[len(c.pools)-1]
					c.pools[len(c.pools)-1] = nil
					c.pools = c.pools[:len(c.pools)-1]

					break
				}
			}

			ch <- c.pools
		},
	})

	stopCh := make(chan struct{})
	defer close(stopCh)

	go cache.Start(stopCh)

	if ok := cache.WaitForCacheSync(stopCh); ok {
		c.log.Debug("Pool cache synced.")
	}

	<-stopCh

	return nil
}
