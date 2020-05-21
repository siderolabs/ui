// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package backend

import (
	"context"
	"sync"

	"github.com/talos-systems/metal-controller-manager/api/v1alpha1"
	"github.com/wailsapp/wails"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/scale/scheme"
	toolscache "k8s.io/client-go/tools/cache"
	"sigs.k8s.io/controller-runtime/pkg/cache"
)

type ServerClasses struct {
	config *rest.Config
	log    *wails.CustomLogger

	serverClasses []*v1alpha1.ServerClass

	sync.Mutex
}

func (c *ServerClasses) WailsInit(runtime *wails.Runtime) error {
	c.log = runtime.Log.New("ServerClasses")

	ch, err := c.watch()
	if err != nil {
		return err
	}

	go func() {
		for serverClasses := range ch {
			c.log.Debugf("%+v", serverClasses)
			runtime.Events.Emit("serverClasses", serverClasses)
		}
	}()

	return nil
}

func (c *ServerClasses) ServerClasses() []*v1alpha1.ServerClass {
	c.Lock()
	defer c.Unlock()

	return c.serverClasses
}

func (c *ServerClasses) watch() (chan []*v1alpha1.ServerClass, error) {
	serverCh := make(chan []*v1alpha1.ServerClass, 100)

	s := runtime.NewScheme()
	_ = scheme.AddToScheme(s)
	_ = v1alpha1.AddToScheme(s)

	cache, err := cache.New(c.config, cache.Options{Scheme: s})
	if err != nil {
		return nil, err
	}

	informer, err := cache.GetInformer(context.TODO(), &v1alpha1.ServerClass{})
	if err != nil {
		return nil, err
	}

	informer.AddEventHandler(toolscache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			c.Lock()
			defer c.Unlock()

			server := obj.(*v1alpha1.ServerClass)

			// TODO(andrewrynhard): Remove this once we figure out why these
			// fields are causing the JSON decoder on the frontend to fail.
			server.ManagedFields = nil
			server.Annotations = nil

			c.serverClasses = append(c.serverClasses, server)

			serverCh <- c.serverClasses
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			c.Lock()
			defer c.Unlock()

			server := newObj.(*v1alpha1.ServerClass)

			for i, old := range c.serverClasses {
				if old.UID == server.UID {
					// TODO(andrewrynhard): Remove this once we figure out why these
					// fields are causing the JSON decoder on the frontend to fail.
					server.ManagedFields = nil
					server.Annotations = nil

					c.serverClasses[i] = server

					break
				}
			}

			serverCh <- c.serverClasses
		},
		DeleteFunc: func(obj interface{}) {
			c.Lock()
			defer c.Unlock()

			server := obj.(*v1alpha1.ServerClass)

			for i, old := range c.serverClasses {
				if old.UID == server.UID {
					c.serverClasses[i] = c.serverClasses[len(c.serverClasses)-1]
					c.serverClasses[len(c.serverClasses)-1] = nil
					c.serverClasses = c.serverClasses[:len(c.serverClasses)-1]

					break
				}
			}

			serverCh <- c.serverClasses
		},
	})

	stopCh := make(chan struct{})
	defer close(stopCh)

	go cache.Start(stopCh)

	if ok := cache.WaitForCacheSync(stopCh); ok {
		c.log.Debug("Server cache synced.")
	}

	return serverCh, nil
}
