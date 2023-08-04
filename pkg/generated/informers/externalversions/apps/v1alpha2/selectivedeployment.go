/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha2

import (
	"context"
	time "time"

	appsv1alpha2 "github.com/ubombar/namespace-request-controller/pkg/apis/apps/v1alpha2"
	versioned "github.com/ubombar/namespace-request-controller/pkg/generated/clientset/versioned"
	internalinterfaces "github.com/ubombar/namespace-request-controller/pkg/generated/informers/externalversions/internalinterfaces"
	v1alpha2 "github.com/ubombar/namespace-request-controller/pkg/generated/listers/apps/v1alpha2"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// SelectiveDeploymentInformer provides access to a shared informer and lister for
// SelectiveDeployments.
type SelectiveDeploymentInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha2.SelectiveDeploymentLister
}

type selectiveDeploymentInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewSelectiveDeploymentInformer constructs a new informer for SelectiveDeployment type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewSelectiveDeploymentInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredSelectiveDeploymentInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredSelectiveDeploymentInformer constructs a new informer for SelectiveDeployment type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredSelectiveDeploymentInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AppsV1alpha2().SelectiveDeployments(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AppsV1alpha2().SelectiveDeployments(namespace).Watch(context.TODO(), options)
			},
		},
		&appsv1alpha2.SelectiveDeployment{},
		resyncPeriod,
		indexers,
	)
}

func (f *selectiveDeploymentInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredSelectiveDeploymentInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *selectiveDeploymentInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&appsv1alpha2.SelectiveDeployment{}, f.defaultInformer)
}

func (f *selectiveDeploymentInformer) Lister() v1alpha2.SelectiveDeploymentLister {
	return v1alpha2.NewSelectiveDeploymentLister(f.Informer().GetIndexer())
}
