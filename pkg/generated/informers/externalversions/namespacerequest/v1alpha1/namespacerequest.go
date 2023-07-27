// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	namespacerequestv1alpha1 "github.com/ubombar/namespace-request-controller/pkg/apis/namespacerequest/v1alpha1"
	versioned "github.com/ubombar/namespace-request-controller/pkg/generated/clientset/versioned"
	internalinterfaces "github.com/ubombar/namespace-request-controller/pkg/generated/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/ubombar/namespace-request-controller/pkg/generated/listers/namespacerequest/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// NamespaceRequestInformer provides access to a shared informer and lister for
// NamespaceRequests.
type NamespaceRequestInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.NamespaceRequestLister
}

type namespaceRequestInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewNamespaceRequestInformer constructs a new informer for NamespaceRequest type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewNamespaceRequestInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredNamespaceRequestInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredNamespaceRequestInformer constructs a new informer for NamespaceRequest type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredNamespaceRequestInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.NamespacerequestV1alpha1().NamespaceRequests(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.NamespacerequestV1alpha1().NamespaceRequests(namespace).Watch(context.TODO(), options)
			},
		},
		&namespacerequestv1alpha1.NamespaceRequest{},
		resyncPeriod,
		indexers,
	)
}

func (f *namespaceRequestInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredNamespaceRequestInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *namespaceRequestInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&namespacerequestv1alpha1.NamespaceRequest{}, f.defaultInformer)
}

func (f *namespaceRequestInformer) Lister() v1alpha1.NamespaceRequestLister {
	return v1alpha1.NewNamespaceRequestLister(f.Informer().GetIndexer())
}