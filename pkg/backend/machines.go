// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package backend

import (
	"context"
	"sync"

	"github.com/wailsapp/wails"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/scale/scheme"
	toolscache "k8s.io/client-go/tools/cache"
	"sigs.k8s.io/cluster-api/api/v1alpha2"
	"sigs.k8s.io/controller-runtime/pkg/cache"
)

type Machines struct {
	config *rest.Config
	log    *wails.CustomLogger

	machines []*v1alpha2.Machine

	sync.Mutex
}

func (c *Machines) WailsInit(runtime *wails.Runtime) error {
	c.log = runtime.Log.New("Machines")

	ch, err := c.watch()
	if err != nil {
		return err
	}

	go func() {
		for machines := range ch {
			c.log.Debugf("%+v", machines)
			runtime.Events.Emit("machines", machines)
		}
	}()

	return nil
}

func (c *Machines) Machines() []*v1alpha2.Machine {
	c.Lock()
	defer c.Unlock()

	return c.machines
}

func (c *Machines) watch() (chan []*v1alpha2.Machine, error) {
	machineCh := make(chan []*v1alpha2.Machine)

	s := runtime.NewScheme()
	_ = scheme.AddToScheme(s)
	_ = v1alpha2.AddToScheme(s)

	cache, err := cache.New(c.config, cache.Options{Scheme: s})
	if err != nil {
		return nil, err
	}

	informer, err := cache.GetInformer(context.TODO(), &v1alpha2.Machine{})
	if err != nil {
		return nil, err
	}

	informer.AddEventHandler(toolscache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			c.Lock()
			defer c.Unlock()

			machine := obj.(*v1alpha2.Machine)

			c.machines = append(c.machines, machine)

			machineCh <- c.machines
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			c.Lock()
			defer c.Unlock()

			machine := newObj.(*v1alpha2.Machine)

			for i, old := range c.machines {
				if old.UID == machine.UID {
					c.machines[i] = machine

					break
				}
			}

			machineCh <- c.machines
		},
		DeleteFunc: func(obj interface{}) {
			c.Lock()
			defer c.Unlock()

			machine := obj.(*v1alpha2.Machine)

			for i, old := range c.machines {
				if old.UID == machine.UID {
					c.machines[i] = c.machines[len(c.machines)-1]
					c.machines[len(c.machines)-1] = nil
					c.machines = c.machines[:len(c.machines)-1]

					break
				}
			}

			machineCh <- c.machines
		},
	})

	stopCh := make(chan struct{})
	defer close(stopCh)

	go cache.Start(stopCh)

	if ok := cache.WaitForCacheSync(stopCh); ok {
		c.log.Debug("Machine cache synced.")
	}

	return machineCh, nil
}
