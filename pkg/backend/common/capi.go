// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package common

import (
	"context"
	"fmt"
	"sync"

	"k8s.io/apimachinery/pkg/runtime"

	"github.com/wailsapp/wails"
	toolscache "k8s.io/client-go/tools/cache"
	"sigs.k8s.io/controller-runtime/pkg/cache"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
)

const (
	EventUpdate = "update"
	EventAdd    = "added"
	EventDelete = "delete"
)

type event struct {
	EventType     string      `json:"type"`
	ContextName   string      `json:"contextName"`
	RuntimeObject string      `json:"runtimeObject"`
	Payload       interface{} `json:"payload"`
}

type updatePayload struct {
	Old interface{} `json:"oldObj"`
	New interface{} `json:"newObj"`
}

type subscription struct {
	cache  cache.Cache
	stopCh chan struct{}
}

// CAPI provides connection between k8s cache informer and JS view.
type CAPI struct {
	sync.RWMutex
	subscriptions map[string]*subscription
	wailsRuntime  *wails.Runtime
	scheme        *runtime.Scheme
}

// NewCAPI creates new capi instance.
func NewCAPI(scheme *runtime.Scheme) *CAPI {
	return &CAPI{
		subscriptions: map[string]*subscription{},
		scheme:        scheme,
	}
}

func (w *CAPI) WailsInit(wailsRuntime *wails.Runtime) error {
	w.wailsRuntime = wailsRuntime
	return nil
}

// GetSubscriptionID returns subscription id for context + type.
func (w *CAPI) GetSubscriptionID(contextName, runtimeObjectName string) string {
	return fmt.Sprintf("%s/%s", contextName, runtimeObjectName)
}

// Subscribe to an informer.
// Returns unique subscription id for cluster/informer type.
func (w *CAPI) Subscribe(contextName, runtimeObjectName string) error {
	if w.hasSubscription(contextName, runtimeObjectName) {
		err := w.Unsubscribe(contextName, runtimeObjectName)
		if err != nil {
			return err
		}
	}

	obj, ok := Types[runtimeObjectName]
	if !ok {
		return fmt.Errorf("No mapping defined for runtimeObjectName %s", runtimeObjectName)
	}

	restConfig, err := config.GetConfigWithContext(contextName)
	if err != nil {
		return err
	}

	ctx := context.Background()
	id := w.GetSubscriptionID(contextName, runtimeObjectName)

	cache, err := cache.New(restConfig, cache.Options{Scheme: w.scheme})
	if err != nil {
		return err
	}

	informer, err := cache.GetInformer(ctx, obj)
	if err != nil {
		return err
	}

	w.Lock()
	defer w.Unlock()

	informer.AddEventHandler(toolscache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			w.wailsRuntime.Events.Emit("capiEvent", &event{
				EventType:     EventAdd,
				ContextName:   contextName,
				RuntimeObject: runtimeObjectName,
				Payload:       obj,
			})
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			w.wailsRuntime.Events.Emit("capiEvent", &event{
				EventType:     EventUpdate,
				ContextName:   contextName,
				RuntimeObject: runtimeObjectName,
				Payload: &updatePayload{
					Old: oldObj,
					New: newObj,
				},
			})
		},
		DeleteFunc: func(obj interface{}) {
			w.wailsRuntime.Events.Emit("capiEvent", &event{
				EventType:     EventDelete,
				ContextName:   contextName,
				RuntimeObject: runtimeObjectName,
				Payload:       obj,
			})
		},
	})

	stopCh := make(chan struct{})
	w.subscriptions[id] = &subscription{
		stopCh: stopCh,
		cache:  cache,
	}
	go cache.Start(stopCh) // nolint:errcheck
	if ok := cache.WaitForCacheSync(stopCh); !ok {
		close(stopCh)
		return fmt.Errorf("Failed to sync cache")
	}

	return nil
}

// Unsubscribe from an informer.
func (w *CAPI) Unsubscribe(contextName, runtimeObjectName string) error {
	w.Lock()
	defer w.Unlock()

	id := w.GetSubscriptionID(contextName, runtimeObjectName)

	if s, ok := w.subscriptions[id]; ok {
		if s.stopCh != nil {
			close(s.stopCh)
		}
		delete(w.subscriptions, id)
	} else {
		return fmt.Errorf("No subscription with id %s", id)
	}
	return nil
}

func (w *CAPI) hasSubscription(contextName, runtimeObjectName string) bool {
	w.RLock()
	defer w.RUnlock()

	return w.subscriptions[w.GetSubscriptionID(contextName, runtimeObjectName)] != nil
}
