// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package backend

import (
	"context"
	"sync"
	"time"

	"github.com/talos-systems/metal-controller-manager/api/v1alpha1"
	"github.com/wailsapp/wails"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/scale/scheme"
	toolscache "k8s.io/client-go/tools/cache"
	"sigs.k8s.io/controller-runtime/pkg/cache"
)

type Environments struct {
	config *rest.Config
	log    *wails.CustomLogger

	environments []*v1alpha1.Environment

	sync.Mutex
}

func (c *Environments) WailsInit(runtime *wails.Runtime) error {
	c.log = runtime.Log.New("Environments")

	ch := make(chan []*v1alpha1.Environment, 100)

	go func() {
		err := c.watch(ch)
		if err != nil {
			c.log.Errorf("Environment watch failed: %v", err)
		}
	}()

	go func() {
		// TODO(andrewrynhard): There seems to be a race condition between the
		// frontend and the backend that causes the first events to be dropped by
		// the frontend. Remove this sleep once we have a fix.
		time.Sleep(1 * time.Second)

		for environments := range ch {
			c.log.Debugf("%+v", environments)
			runtime.Events.Emit("environments", environments)
		}
	}()

	return nil
}

func (c *Environments) Environments() []*v1alpha1.Environment {
	c.Lock()
	defer c.Unlock()

	return c.environments
}

func (c *Environments) watch(ch chan []*v1alpha1.Environment) error {
	s := runtime.NewScheme()
	_ = scheme.AddToScheme(s)
	_ = v1alpha1.AddToScheme(s)

	cache, err := cache.New(c.config, cache.Options{Scheme: s})
	if err != nil {
		return err
	}

	informer, err := cache.GetInformer(context.TODO(), &v1alpha1.Environment{})
	if err != nil {
		return err
	}

	informer.AddEventHandler(toolscache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			c.Lock()
			defer c.Unlock()

			environment := obj.(*v1alpha1.Environment)

			// TODO(andrewrynhard): Remove this once we figure out why these
			// fields are causing the JSON decoder on the frontend to fail.
			environment.ManagedFields = nil
			environment.Annotations = nil

			c.environments = append(c.environments, environment)

			ch <- c.environments
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			c.Lock()
			defer c.Unlock()

			environment := newObj.(*v1alpha1.Environment)

			for i, old := range c.environments {
				if old.UID == environment.UID {
					// TODO(andrewrynhard): Remove this once we figure out why these
					// fields are causing the JSON decoder on the frontend to fail.
					environment.ManagedFields = nil
					environment.Annotations = nil

					c.environments[i] = environment

					break
				}
			}

			ch <- c.environments
		},
		DeleteFunc: func(obj interface{}) {
			c.Lock()
			defer c.Unlock()

			environment := obj.(*v1alpha1.Environment)

			for i, old := range c.environments {
				if old.UID == environment.UID {
					c.environments[i] = c.environments[len(c.environments)-1]
					c.environments[len(c.environments)-1] = nil
					c.environments = c.environments[:len(c.environments)-1]

					break
				}
			}

			ch <- c.environments
		},
	})

	stopCh := make(chan struct{})
	defer close(stopCh)

	go cache.Start(stopCh)

	if ok := cache.WaitForCacheSync(stopCh); ok {
		c.log.Debug("Environment cache synced.")
	}

	<-stopCh

	return nil
}
