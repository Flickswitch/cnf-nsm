/*
 * Copyright (c) 2019 PANTHEON.tech s.r.o. All rights reserved.
 *
 * CNF-NSM License. For licensing terms please contact sales@pantheon.tech.
 */

// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "go.cdnf.io/cnf-nsm/plugins/crdplugin/pkg/apis/pantheontech/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// CNFConfigurationLister helps list CNFConfigurations.
type CNFConfigurationLister interface {
	// List lists all CNFConfigurations in the indexer.
	List(selector labels.Selector) (ret []*v1.CNFConfiguration, err error)
	// CNFConfigurations returns an object that can list and get CNFConfigurations.
	CNFConfigurations(namespace string) CNFConfigurationNamespaceLister
	CNFConfigurationListerExpansion
}

// cNFConfigurationLister implements the CNFConfigurationLister interface.
type cNFConfigurationLister struct {
	indexer cache.Indexer
}

// NewCNFConfigurationLister returns a new CNFConfigurationLister.
func NewCNFConfigurationLister(indexer cache.Indexer) CNFConfigurationLister {
	return &cNFConfigurationLister{indexer: indexer}
}

// List lists all CNFConfigurations in the indexer.
func (s *cNFConfigurationLister) List(selector labels.Selector) (ret []*v1.CNFConfiguration, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.CNFConfiguration))
	})
	return ret, err
}

// CNFConfigurations returns an object that can list and get CNFConfigurations.
func (s *cNFConfigurationLister) CNFConfigurations(namespace string) CNFConfigurationNamespaceLister {
	return cNFConfigurationNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// CNFConfigurationNamespaceLister helps list and get CNFConfigurations.
type CNFConfigurationNamespaceLister interface {
	// List lists all CNFConfigurations in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.CNFConfiguration, err error)
	// Get retrieves the CNFConfiguration from the indexer for a given namespace and name.
	Get(name string) (*v1.CNFConfiguration, error)
	CNFConfigurationNamespaceListerExpansion
}

// cNFConfigurationNamespaceLister implements the CNFConfigurationNamespaceLister
// interface.
type cNFConfigurationNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all CNFConfigurations in the indexer for a given namespace.
func (s cNFConfigurationNamespaceLister) List(selector labels.Selector) (ret []*v1.CNFConfiguration, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.CNFConfiguration))
	})
	return ret, err
}

// Get retrieves the CNFConfiguration from the indexer for a given namespace and name.
func (s cNFConfigurationNamespaceLister) Get(name string) (*v1.CNFConfiguration, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("cnfconfiguration"), name)
	}
	return obj.(*v1.CNFConfiguration), nil
}
