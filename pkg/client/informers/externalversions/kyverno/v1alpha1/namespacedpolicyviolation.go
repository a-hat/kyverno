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
	time "time"

	kyvernov1alpha1 "github.com/nirmata/kyverno/pkg/api/kyverno/v1alpha1"
	versioned "github.com/nirmata/kyverno/pkg/client/clientset/versioned"
	internalinterfaces "github.com/nirmata/kyverno/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/nirmata/kyverno/pkg/client/listers/kyverno/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// NamespacedPolicyViolationInformer provides access to a shared informer and lister for
// NamespacedPolicyViolations.
type NamespacedPolicyViolationInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.NamespacedPolicyViolationLister
}

type namespacedPolicyViolationInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewNamespacedPolicyViolationInformer constructs a new informer for NamespacedPolicyViolation type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewNamespacedPolicyViolationInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredNamespacedPolicyViolationInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredNamespacedPolicyViolationInformer constructs a new informer for NamespacedPolicyViolation type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredNamespacedPolicyViolationInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.KyvernoV1alpha1().NamespacedPolicyViolations(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.KyvernoV1alpha1().NamespacedPolicyViolations(namespace).Watch(options)
			},
		},
		&kyvernov1alpha1.NamespacedPolicyViolation{},
		resyncPeriod,
		indexers,
	)
}

func (f *namespacedPolicyViolationInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredNamespacedPolicyViolationInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *namespacedPolicyViolationInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&kyvernov1alpha1.NamespacedPolicyViolation{}, f.defaultInformer)
}

func (f *namespacedPolicyViolationInformer) Lister() v1alpha1.NamespacedPolicyViolationLister {
	return v1alpha1.NewNamespacedPolicyViolationLister(f.Informer().GetIndexer())
}
