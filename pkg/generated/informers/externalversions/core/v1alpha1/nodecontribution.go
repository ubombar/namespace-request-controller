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

package v1alpha1

import (
	"context"
	time "time"

	corev1alpha1 "github.com/ubombar/namespace-request-controller/pkg/apis/core/v1alpha1"
	versioned "github.com/ubombar/namespace-request-controller/pkg/generated/clientset/versioned"
	internalinterfaces "github.com/ubombar/namespace-request-controller/pkg/generated/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/ubombar/namespace-request-controller/pkg/generated/listers/core/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// NodeContributionInformer provides access to a shared informer and lister for
// NodeContributions.
type NodeContributionInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.NodeContributionLister
}

type nodeContributionInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewNodeContributionInformer constructs a new informer for NodeContribution type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewNodeContributionInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredNodeContributionInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredNodeContributionInformer constructs a new informer for NodeContribution type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredNodeContributionInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CoreV1alpha1().NodeContributions().List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CoreV1alpha1().NodeContributions().Watch(context.TODO(), options)
			},
		},
		&corev1alpha1.NodeContribution{},
		resyncPeriod,
		indexers,
	)
}

func (f *nodeContributionInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredNodeContributionInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *nodeContributionInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&corev1alpha1.NodeContribution{}, f.defaultInformer)
}

func (f *nodeContributionInformer) Lister() v1alpha1.NodeContributionLister {
	return v1alpha1.NewNodeContributionLister(f.Informer().GetIndexer())
}
