// Copyright 2020-2021 the Pinniped contributors. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"time"

	v1alpha1 "go.pinniped.dev/generated/1.17/apis/supervisor/idp/v1alpha1"
	scheme "go.pinniped.dev/generated/1.17/client/supervisor/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// LDAPIdentityProvidersGetter has a method to return a LDAPIdentityProviderInterface.
// A group's client should implement this interface.
type LDAPIdentityProvidersGetter interface {
	LDAPIdentityProviders(namespace string) LDAPIdentityProviderInterface
}

// LDAPIdentityProviderInterface has methods to work with LDAPIdentityProvider resources.
type LDAPIdentityProviderInterface interface {
	Create(*v1alpha1.LDAPIdentityProvider) (*v1alpha1.LDAPIdentityProvider, error)
	Update(*v1alpha1.LDAPIdentityProvider) (*v1alpha1.LDAPIdentityProvider, error)
	UpdateStatus(*v1alpha1.LDAPIdentityProvider) (*v1alpha1.LDAPIdentityProvider, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.LDAPIdentityProvider, error)
	List(opts v1.ListOptions) (*v1alpha1.LDAPIdentityProviderList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.LDAPIdentityProvider, err error)
	LDAPIdentityProviderExpansion
}

// lDAPIdentityProviders implements LDAPIdentityProviderInterface
type lDAPIdentityProviders struct {
	client rest.Interface
	ns     string
}

// newLDAPIdentityProviders returns a LDAPIdentityProviders
func newLDAPIdentityProviders(c *IDPV1alpha1Client, namespace string) *lDAPIdentityProviders {
	return &lDAPIdentityProviders{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the lDAPIdentityProvider, and returns the corresponding lDAPIdentityProvider object, and an error if there is any.
func (c *lDAPIdentityProviders) Get(name string, options v1.GetOptions) (result *v1alpha1.LDAPIdentityProvider, err error) {
	result = &v1alpha1.LDAPIdentityProvider{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("ldapidentityproviders").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of LDAPIdentityProviders that match those selectors.
func (c *lDAPIdentityProviders) List(opts v1.ListOptions) (result *v1alpha1.LDAPIdentityProviderList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.LDAPIdentityProviderList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("ldapidentityproviders").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested lDAPIdentityProviders.
func (c *lDAPIdentityProviders) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("ldapidentityproviders").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a lDAPIdentityProvider and creates it.  Returns the server's representation of the lDAPIdentityProvider, and an error, if there is any.
func (c *lDAPIdentityProviders) Create(lDAPIdentityProvider *v1alpha1.LDAPIdentityProvider) (result *v1alpha1.LDAPIdentityProvider, err error) {
	result = &v1alpha1.LDAPIdentityProvider{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("ldapidentityproviders").
		Body(lDAPIdentityProvider).
		Do().
		Into(result)
	return
}

// Update takes the representation of a lDAPIdentityProvider and updates it. Returns the server's representation of the lDAPIdentityProvider, and an error, if there is any.
func (c *lDAPIdentityProviders) Update(lDAPIdentityProvider *v1alpha1.LDAPIdentityProvider) (result *v1alpha1.LDAPIdentityProvider, err error) {
	result = &v1alpha1.LDAPIdentityProvider{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("ldapidentityproviders").
		Name(lDAPIdentityProvider.Name).
		Body(lDAPIdentityProvider).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *lDAPIdentityProviders) UpdateStatus(lDAPIdentityProvider *v1alpha1.LDAPIdentityProvider) (result *v1alpha1.LDAPIdentityProvider, err error) {
	result = &v1alpha1.LDAPIdentityProvider{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("ldapidentityproviders").
		Name(lDAPIdentityProvider.Name).
		SubResource("status").
		Body(lDAPIdentityProvider).
		Do().
		Into(result)
	return
}

// Delete takes name of the lDAPIdentityProvider and deletes it. Returns an error if one occurs.
func (c *lDAPIdentityProviders) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("ldapidentityproviders").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *lDAPIdentityProviders) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("ldapidentityproviders").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched lDAPIdentityProvider.
func (c *lDAPIdentityProviders) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.LDAPIdentityProvider, err error) {
	result = &v1alpha1.LDAPIdentityProvider{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("ldapidentityproviders").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
