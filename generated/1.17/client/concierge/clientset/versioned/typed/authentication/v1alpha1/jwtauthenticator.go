// Copyright 2020 the Pinniped contributors. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"time"

	v1alpha1 "go.pinniped.dev/generated/1.17/apis/concierge/authentication/v1alpha1"
	scheme "go.pinniped.dev/generated/1.17/client/concierge/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// JWTAuthenticatorsGetter has a method to return a JWTAuthenticatorInterface.
// A group's client should implement this interface.
type JWTAuthenticatorsGetter interface {
	JWTAuthenticators(namespace string) JWTAuthenticatorInterface
}

// JWTAuthenticatorInterface has methods to work with JWTAuthenticator resources.
type JWTAuthenticatorInterface interface {
	Create(*v1alpha1.JWTAuthenticator) (*v1alpha1.JWTAuthenticator, error)
	Update(*v1alpha1.JWTAuthenticator) (*v1alpha1.JWTAuthenticator, error)
	UpdateStatus(*v1alpha1.JWTAuthenticator) (*v1alpha1.JWTAuthenticator, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.JWTAuthenticator, error)
	List(opts v1.ListOptions) (*v1alpha1.JWTAuthenticatorList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.JWTAuthenticator, err error)
	JWTAuthenticatorExpansion
}

// jWTAuthenticators implements JWTAuthenticatorInterface
type jWTAuthenticators struct {
	client rest.Interface
	ns     string
}

// newJWTAuthenticators returns a JWTAuthenticators
func newJWTAuthenticators(c *AuthenticationV1alpha1Client, namespace string) *jWTAuthenticators {
	return &jWTAuthenticators{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the jWTAuthenticator, and returns the corresponding jWTAuthenticator object, and an error if there is any.
func (c *jWTAuthenticators) Get(name string, options v1.GetOptions) (result *v1alpha1.JWTAuthenticator, err error) {
	result = &v1alpha1.JWTAuthenticator{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("jwtauthenticators").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of JWTAuthenticators that match those selectors.
func (c *jWTAuthenticators) List(opts v1.ListOptions) (result *v1alpha1.JWTAuthenticatorList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.JWTAuthenticatorList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("jwtauthenticators").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested jWTAuthenticators.
func (c *jWTAuthenticators) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("jwtauthenticators").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a jWTAuthenticator and creates it.  Returns the server's representation of the jWTAuthenticator, and an error, if there is any.
func (c *jWTAuthenticators) Create(jWTAuthenticator *v1alpha1.JWTAuthenticator) (result *v1alpha1.JWTAuthenticator, err error) {
	result = &v1alpha1.JWTAuthenticator{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("jwtauthenticators").
		Body(jWTAuthenticator).
		Do().
		Into(result)
	return
}

// Update takes the representation of a jWTAuthenticator and updates it. Returns the server's representation of the jWTAuthenticator, and an error, if there is any.
func (c *jWTAuthenticators) Update(jWTAuthenticator *v1alpha1.JWTAuthenticator) (result *v1alpha1.JWTAuthenticator, err error) {
	result = &v1alpha1.JWTAuthenticator{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("jwtauthenticators").
		Name(jWTAuthenticator.Name).
		Body(jWTAuthenticator).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *jWTAuthenticators) UpdateStatus(jWTAuthenticator *v1alpha1.JWTAuthenticator) (result *v1alpha1.JWTAuthenticator, err error) {
	result = &v1alpha1.JWTAuthenticator{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("jwtauthenticators").
		Name(jWTAuthenticator.Name).
		SubResource("status").
		Body(jWTAuthenticator).
		Do().
		Into(result)
	return
}

// Delete takes name of the jWTAuthenticator and deletes it. Returns an error if one occurs.
func (c *jWTAuthenticators) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("jwtauthenticators").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *jWTAuthenticators) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("jwtauthenticators").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched jWTAuthenticator.
func (c *jWTAuthenticators) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.JWTAuthenticator, err error) {
	result = &v1alpha1.JWTAuthenticator{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("jwtauthenticators").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
