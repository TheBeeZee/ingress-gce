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

// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	v1 "k8s.io/ingress-gce/pkg/apis/serviceattachment/v1"
	scheme "k8s.io/ingress-gce/pkg/serviceattachment/client/clientset/versioned/scheme"
)

// ServiceAttachmentsGetter has a method to return a ServiceAttachmentInterface.
// A group's client should implement this interface.
type ServiceAttachmentsGetter interface {
	ServiceAttachments(namespace string) ServiceAttachmentInterface
}

// ServiceAttachmentInterface has methods to work with ServiceAttachment resources.
type ServiceAttachmentInterface interface {
	Create(ctx context.Context, serviceAttachment *v1.ServiceAttachment, opts metav1.CreateOptions) (*v1.ServiceAttachment, error)
	Update(ctx context.Context, serviceAttachment *v1.ServiceAttachment, opts metav1.UpdateOptions) (*v1.ServiceAttachment, error)
	UpdateStatus(ctx context.Context, serviceAttachment *v1.ServiceAttachment, opts metav1.UpdateOptions) (*v1.ServiceAttachment, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.ServiceAttachment, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.ServiceAttachmentList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.ServiceAttachment, err error)
	ServiceAttachmentExpansion
}

// serviceAttachments implements ServiceAttachmentInterface
type serviceAttachments struct {
	client rest.Interface
	ns     string
}

// newServiceAttachments returns a ServiceAttachments
func newServiceAttachments(c *NetworkingV1Client, namespace string) *serviceAttachments {
	return &serviceAttachments{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the serviceAttachment, and returns the corresponding serviceAttachment object, and an error if there is any.
func (c *serviceAttachments) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.ServiceAttachment, err error) {
	result = &v1.ServiceAttachment{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("serviceattachments").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ServiceAttachments that match those selectors.
func (c *serviceAttachments) List(ctx context.Context, opts metav1.ListOptions) (result *v1.ServiceAttachmentList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.ServiceAttachmentList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("serviceattachments").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested serviceAttachments.
func (c *serviceAttachments) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("serviceattachments").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a serviceAttachment and creates it.  Returns the server's representation of the serviceAttachment, and an error, if there is any.
func (c *serviceAttachments) Create(ctx context.Context, serviceAttachment *v1.ServiceAttachment, opts metav1.CreateOptions) (result *v1.ServiceAttachment, err error) {
	result = &v1.ServiceAttachment{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("serviceattachments").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(serviceAttachment).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a serviceAttachment and updates it. Returns the server's representation of the serviceAttachment, and an error, if there is any.
func (c *serviceAttachments) Update(ctx context.Context, serviceAttachment *v1.ServiceAttachment, opts metav1.UpdateOptions) (result *v1.ServiceAttachment, err error) {
	result = &v1.ServiceAttachment{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("serviceattachments").
		Name(serviceAttachment.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(serviceAttachment).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *serviceAttachments) UpdateStatus(ctx context.Context, serviceAttachment *v1.ServiceAttachment, opts metav1.UpdateOptions) (result *v1.ServiceAttachment, err error) {
	result = &v1.ServiceAttachment{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("serviceattachments").
		Name(serviceAttachment.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(serviceAttachment).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the serviceAttachment and deletes it. Returns an error if one occurs.
func (c *serviceAttachments) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("serviceattachments").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *serviceAttachments) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("serviceattachments").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched serviceAttachment.
func (c *serviceAttachments) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.ServiceAttachment, err error) {
	result = &v1.ServiceAttachment{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("serviceattachments").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
