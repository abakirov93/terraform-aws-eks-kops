/*
Copyright 2020 The Kubernetes Authors.

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

package internalversion

import (
	"context"
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	kops "k8s.io/kops/pkg/apis/kops"
	scheme "k8s.io/kops/pkg/client/clientset_generated/internalclientset/scheme"
)

// SSHCredentialsGetter has a method to return a SSHCredentialInterface.
// A group's client should implement this interface.
type SSHCredentialsGetter interface {
	SSHCredentials(namespace string) SSHCredentialInterface
}

// SSHCredentialInterface has methods to work with SSHCredential resources.
type SSHCredentialInterface interface {
	Create(ctx context.Context, sSHCredential *kops.SSHCredential, opts v1.CreateOptions) (*kops.SSHCredential, error)
	Update(ctx context.Context, sSHCredential *kops.SSHCredential, opts v1.UpdateOptions) (*kops.SSHCredential, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*kops.SSHCredential, error)
	List(ctx context.Context, opts v1.ListOptions) (*kops.SSHCredentialList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *kops.SSHCredential, err error)
	SSHCredentialExpansion
}

// sSHCredentials implements SSHCredentialInterface
type sSHCredentials struct {
	client rest.Interface
	ns     string
}

// newSSHCredentials returns a SSHCredentials
func newSSHCredentials(c *KopsClient, namespace string) *sSHCredentials {
	return &sSHCredentials{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the sSHCredential, and returns the corresponding sSHCredential object, and an error if there is any.
func (c *sSHCredentials) Get(ctx context.Context, name string, options v1.GetOptions) (result *kops.SSHCredential, err error) {
	result = &kops.SSHCredential{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("sshcredentials").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of SSHCredentials that match those selectors.
func (c *sSHCredentials) List(ctx context.Context, opts v1.ListOptions) (result *kops.SSHCredentialList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &kops.SSHCredentialList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("sshcredentials").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested sSHCredentials.
func (c *sSHCredentials) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("sshcredentials").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a sSHCredential and creates it.  Returns the server's representation of the sSHCredential, and an error, if there is any.
func (c *sSHCredentials) Create(ctx context.Context, sSHCredential *kops.SSHCredential, opts v1.CreateOptions) (result *kops.SSHCredential, err error) {
	result = &kops.SSHCredential{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("sshcredentials").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(sSHCredential).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a sSHCredential and updates it. Returns the server's representation of the sSHCredential, and an error, if there is any.
func (c *sSHCredentials) Update(ctx context.Context, sSHCredential *kops.SSHCredential, opts v1.UpdateOptions) (result *kops.SSHCredential, err error) {
	result = &kops.SSHCredential{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("sshcredentials").
		Name(sSHCredential.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(sSHCredential).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the sSHCredential and deletes it. Returns an error if one occurs.
func (c *sSHCredentials) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("sshcredentials").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *sSHCredentials) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("sshcredentials").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched sSHCredential.
func (c *sSHCredentials) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *kops.SSHCredential, err error) {
	result = &kops.SSHCredential{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("sshcredentials").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
