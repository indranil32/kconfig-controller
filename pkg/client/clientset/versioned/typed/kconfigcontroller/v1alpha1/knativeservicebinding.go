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

// Code generated by main. DO NOT EDIT.

package v1alpha1

import (
	"time"

	v1alpha1 "github.com/att-cloudnative-labs/kconfig-controller/pkg/apis/kconfigcontroller/v1alpha1"
	scheme "github.com/att-cloudnative-labs/kconfig-controller/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// KnativeServiceBindingsGetter has a method to return a KnativeServiceBindingInterface.
// A group's client should implement this interface.
type KnativeServiceBindingsGetter interface {
	KnativeServiceBindings(namespace string) KnativeServiceBindingInterface
}

// KnativeServiceBindingInterface has methods to work with KnativeServiceBinding resources.
type KnativeServiceBindingInterface interface {
	Create(*v1alpha1.KnativeServiceBinding) (*v1alpha1.KnativeServiceBinding, error)
	Update(*v1alpha1.KnativeServiceBinding) (*v1alpha1.KnativeServiceBinding, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.KnativeServiceBinding, error)
	List(opts v1.ListOptions) (*v1alpha1.KnativeServiceBindingList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.KnativeServiceBinding, err error)
	KnativeServiceBindingExpansion
}

// knativeServiceBindings implements KnativeServiceBindingInterface
type knativeServiceBindings struct {
	client rest.Interface
	ns     string
}

// newKnativeServiceBindings returns a KnativeServiceBindings
func newKnativeServiceBindings(c *KconfigcontrollerV1alpha1Client, namespace string) *knativeServiceBindings {
	return &knativeServiceBindings{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the knativeServiceBinding, and returns the corresponding knativeServiceBinding object, and an error if there is any.
func (c *knativeServiceBindings) Get(name string, options v1.GetOptions) (result *v1alpha1.KnativeServiceBinding, err error) {
	result = &v1alpha1.KnativeServiceBinding{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("knativeservicebindings").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of KnativeServiceBindings that match those selectors.
func (c *knativeServiceBindings) List(opts v1.ListOptions) (result *v1alpha1.KnativeServiceBindingList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.KnativeServiceBindingList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("knativeservicebindings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested knativeServiceBindings.
func (c *knativeServiceBindings) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("knativeservicebindings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a knativeServiceBinding and creates it.  Returns the server's representation of the knativeServiceBinding, and an error, if there is any.
func (c *knativeServiceBindings) Create(knativeServiceBinding *v1alpha1.KnativeServiceBinding) (result *v1alpha1.KnativeServiceBinding, err error) {
	result = &v1alpha1.KnativeServiceBinding{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("knativeservicebindings").
		Body(knativeServiceBinding).
		Do().
		Into(result)
	return
}

// Update takes the representation of a knativeServiceBinding and updates it. Returns the server's representation of the knativeServiceBinding, and an error, if there is any.
func (c *knativeServiceBindings) Update(knativeServiceBinding *v1alpha1.KnativeServiceBinding) (result *v1alpha1.KnativeServiceBinding, err error) {
	result = &v1alpha1.KnativeServiceBinding{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("knativeservicebindings").
		Name(knativeServiceBinding.Name).
		Body(knativeServiceBinding).
		Do().
		Into(result)
	return
}

// Delete takes name of the knativeServiceBinding and deletes it. Returns an error if one occurs.
func (c *knativeServiceBindings) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("knativeservicebindings").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *knativeServiceBindings) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("knativeservicebindings").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched knativeServiceBinding.
func (c *knativeServiceBindings) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.KnativeServiceBinding, err error) {
	result = &v1alpha1.KnativeServiceBinding{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("knativeservicebindings").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
