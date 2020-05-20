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
	"sigs.k8s.io/cluster-api/api/v1alpha2"
	"sigs.k8s.io/controller-runtime/pkg/cache"
)

type Clusters struct {
	config *rest.Config
	log    *wails.CustomLogger

	clusters []*v1alpha2.Cluster

	sync.Mutex
}

func (c *Clusters) WailsInit(runtime *wails.Runtime) error {
	c.log = runtime.Log.New("Clusters")

	ch, err := c.watch()
	if err != nil {
		return err
	}

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

func (c *Clusters) watch() (chan []*v1alpha2.Cluster, error) {
	clusterCh := make(chan []*v1alpha2.Cluster, 100)

	s := runtime.NewScheme()
	_ = scheme.AddToScheme(s)
	_ = v1alpha2.AddToScheme(s)

	cache, err := cache.New(c.config, cache.Options{Scheme: s})
	if err != nil {
		return nil, err
	}

	informer, err := cache.GetInformer(context.TODO(), &v1alpha2.Cluster{})
	if err != nil {
		return nil, err
	}

	informer.AddEventHandler(toolscache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			c.Lock()
			defer c.Unlock()

			cluster := obj.(*v1alpha2.Cluster)

			c.clusters = append(c.clusters, cluster)

			clusterCh <- c.clusters
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			c.Lock()
			defer c.Unlock()

			cluster := newObj.(*v1alpha2.Cluster)

			for i, old := range c.clusters {
				if old.UID == cluster.UID {
					c.clusters[i] = cluster

					break
				}
			}

			clusterCh <- c.clusters
		},
		DeleteFunc: func(obj interface{}) {
			c.Lock()
			defer c.Unlock()

			cluster := obj.(*v1alpha2.Cluster)

			for i, old := range c.clusters {
				if old.UID == cluster.UID {
					c.clusters[i] = c.clusters[len(c.clusters)-1]
					c.clusters[len(c.clusters)-1] = nil
					c.clusters = c.clusters[:len(c.clusters)-1]

					break
				}
			}

			clusterCh <- c.clusters
		},
	})

	stopCh := make(chan struct{})
	defer close(stopCh)

	go cache.Start(stopCh)

	if ok := cache.WaitForCacheSync(stopCh); ok {
		c.log.Debug("Cluster cache synced.")
	}

	return clusterCh, nil
}
