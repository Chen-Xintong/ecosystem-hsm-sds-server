/*
Copyright The Istio Authors

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

	tcsv1alpha1 "github.com/intel-innersource/applications.services.cloud.hsm-sds-server/pkg/apis/tcs/v1alpha1"
	versioned "github.com/intel-innersource/applications.services.cloud.hsm-sds-server/pkg/client/clientset/versioned"
	internalinterfaces "github.com/intel-innersource/applications.services.cloud.hsm-sds-server/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/intel-innersource/applications.services.cloud.hsm-sds-server/pkg/client/listers/tcs/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// QuoteAttestationInformer provides access to a shared informer and lister for
// QuoteAttestations.
type QuoteAttestationInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.QuoteAttestationLister
}

type quoteAttestationInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewQuoteAttestationInformer constructs a new informer for QuoteAttestation type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewQuoteAttestationInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredQuoteAttestationInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredQuoteAttestationInformer constructs a new informer for QuoteAttestation type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredQuoteAttestationInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.TcsV1alpha1().QuoteAttestations(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.TcsV1alpha1().QuoteAttestations(namespace).Watch(context.TODO(), options)
			},
		},
		&tcsv1alpha1.QuoteAttestation{},
		resyncPeriod,
		indexers,
	)
}

func (f *quoteAttestationInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredQuoteAttestationInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *quoteAttestationInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&tcsv1alpha1.QuoteAttestation{}, f.defaultInformer)
}

func (f *quoteAttestationInformer) Lister() v1alpha1.QuoteAttestationLister {
	return v1alpha1.NewQuoteAttestationLister(f.Informer().GetIndexer())
}
