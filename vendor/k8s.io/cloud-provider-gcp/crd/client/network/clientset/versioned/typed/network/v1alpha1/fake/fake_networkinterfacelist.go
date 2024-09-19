/*
Copyright 2023 The Kubernetes Authors.

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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	testing "k8s.io/client-go/testing"
	v1alpha1 "k8s.io/cloud-provider-gcp/crd/apis/network/v1alpha1"
)

// FakeNetworkInterfaceLists implements NetworkInterfaceListInterface
type FakeNetworkInterfaceLists struct {
	Fake *FakeNetworkingV1alpha1
	ns   string
}

var networkinterfacelistsResource = v1alpha1.SchemeGroupVersion.WithResource("networkinterfacelists")

var networkinterfacelistsKind = v1alpha1.SchemeGroupVersion.WithKind("NetworkInterfaceList")

// Get takes name of the networkInterfaceList, and returns the corresponding networkInterfaceList object, and an error if there is any.
func (c *FakeNetworkInterfaceLists) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.NetworkInterfaceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(networkinterfacelistsResource, c.ns, name), &v1alpha1.NetworkInterfaceList{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NetworkInterfaceList), err
}