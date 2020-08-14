// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package arges

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

type Machines struct {
	config *rest.Config
	log    *wails.CustomLogger

	machines []*v1alpha3.Machine

	sync.Mutex
}

func (c *Machines) WailsInit(runtime *wails.Runtime) error {
	c.log = runtime.Log.New("Machines")

	ch := make(chan []*v1alpha3.Machine, 100)

	go func() {
		err := c.watch(ch)
		if err != nil {
			c.log.Errorf("Machine watch failed: %v", err)
		}
	}()

	go func() {
		// TODO(andrewrynhard): There seems to be a race condition between the
		// frontend and the backend that causes the first events to be dropped by
		// the frontend. Remove this sleep once we have a fix.
		time.Sleep(1 * time.Second)

		for machines := range ch {
			c.log.Debugf("%+v", machines)
			runtime.Events.Emit("machines", machines)
		}
	}()

	return nil
}

func (c *Machines) Machines(cluster string) []*v1alpha3.Machine {
	c.Lock()
	defer c.Unlock()

	machines := []*v1alpha3.Machine{}

	for _, machine := range c.machines {
		if machine.Spec.ClusterName == cluster {
			machines = append(machines, machine)
		}
	}

	return machines
}

func (c *Machines) watch(ch chan []*v1alpha3.Machine) error {
	s := runtime.NewScheme()
	_ = scheme.AddToScheme(s)
	_ = v1alpha3.AddToScheme(s)

	cache, err := cache.New(c.config, cache.Options{Scheme: s})
	if err != nil {
		return err
	}

	informer, err := cache.GetInformer(context.TODO(), &v1alpha3.Machine{})
	if err != nil {
		return err
	}

	informer.AddEventHandler(toolscache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			c.Lock()
			defer c.Unlock()

			machine := obj.(*v1alpha3.Machine)

			c.machines = append(c.machines, machine)

			ch <- c.machines
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			c.Lock()
			defer c.Unlock()

			machine := newObj.(*v1alpha3.Machine)

			for i, old := range c.machines {
				if old.UID == machine.UID {
					c.machines[i] = machine

					break
				}
			}

			ch <- c.machines
		},
		DeleteFunc: func(obj interface{}) {
			c.Lock()
			defer c.Unlock()

			machine := obj.(*v1alpha3.Machine)

			for i, old := range c.machines {
				if old.UID == machine.UID {
					c.machines[i] = c.machines[len(c.machines)-1]
					c.machines[len(c.machines)-1] = nil
					c.machines = c.machines[:len(c.machines)-1]

					break
				}
			}

			ch <- c.machines
		},
	})

	stopCh := make(chan struct{})
	defer close(stopCh)

	go cache.Start(stopCh)

	if ok := cache.WaitForCacheSync(stopCh); ok {
		c.log.Debug("Machine cache synced.")
	}

	<-stopCh

	return nil
}
