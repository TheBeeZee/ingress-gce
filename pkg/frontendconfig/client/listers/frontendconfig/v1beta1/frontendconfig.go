/*
Copyright 2024 The Kubernetes Authors.

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

// Code generated by lister-gen. DO NOT EDIT.

package v1beta1

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	v1beta1 "k8s.io/ingress-gce/pkg/apis/frontendconfig/v1beta1"
)

// FrontendConfigLister helps list FrontendConfigs.
// All objects returned here must be treated as read-only.
type FrontendConfigLister interface {
	// List lists all FrontendConfigs in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1beta1.FrontendConfig, err error)
	// FrontendConfigs returns an object that can list and get FrontendConfigs.
	FrontendConfigs(namespace string) FrontendConfigNamespaceLister
	FrontendConfigListerExpansion
}

// frontendConfigLister implements the FrontendConfigLister interface.
type frontendConfigLister struct {
	indexer cache.Indexer
}

// NewFrontendConfigLister returns a new FrontendConfigLister.
func NewFrontendConfigLister(indexer cache.Indexer) FrontendConfigLister {
	return &frontendConfigLister{indexer: indexer}
}

// List lists all FrontendConfigs in the indexer.
func (s *frontendConfigLister) List(selector labels.Selector) (ret []*v1beta1.FrontendConfig, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.FrontendConfig))
	})
	return ret, err
}

// FrontendConfigs returns an object that can list and get FrontendConfigs.
func (s *frontendConfigLister) FrontendConfigs(namespace string) FrontendConfigNamespaceLister {
	return frontendConfigNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// FrontendConfigNamespaceLister helps list and get FrontendConfigs.
// All objects returned here must be treated as read-only.
type FrontendConfigNamespaceLister interface {
	// List lists all FrontendConfigs in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1beta1.FrontendConfig, err error)
	// Get retrieves the FrontendConfig from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1beta1.FrontendConfig, error)
	FrontendConfigNamespaceListerExpansion
}

// frontendConfigNamespaceLister implements the FrontendConfigNamespaceLister
// interface.
type frontendConfigNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all FrontendConfigs in the indexer for a given namespace.
func (s frontendConfigNamespaceLister) List(selector labels.Selector) (ret []*v1beta1.FrontendConfig, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.FrontendConfig))
	})
	return ret, err
}

// Get retrieves the FrontendConfig from the indexer for a given namespace and name.
func (s frontendConfigNamespaceLister) Get(name string) (*v1beta1.FrontendConfig, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1beta1.Resource("frontendconfig"), name)
	}
	return obj.(*v1beta1.FrontendConfig), nil
}
