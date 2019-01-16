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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/gbraxton/kconfig/pkg/apis/kconfigcontroller/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeKconfigs implements KconfigInterface
type FakeKconfigs struct {
	Fake *FakeKconfigcontrollerV1alpha1
	ns   string
}

var kconfigsResource = schema.GroupVersionResource{Group: "kconfigcontroller.atteg.com", Version: "v1alpha1", Resource: "kconfigs"}

var kconfigsKind = schema.GroupVersionKind{Group: "kconfigcontroller.atteg.com", Version: "v1alpha1", Kind: "Kconfig"}

// Get takes name of the kconfig, and returns the corresponding kconfig object, and an error if there is any.
func (c *FakeKconfigs) Get(name string, options v1.GetOptions) (result *v1alpha1.Kconfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(kconfigsResource, c.ns, name), &v1alpha1.Kconfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Kconfig), err
}

// List takes label and field selectors, and returns the list of Kconfigs that match those selectors.
func (c *FakeKconfigs) List(opts v1.ListOptions) (result *v1alpha1.KconfigList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(kconfigsResource, kconfigsKind, c.ns, opts), &v1alpha1.KconfigList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.KconfigList{ListMeta: obj.(*v1alpha1.KconfigList).ListMeta}
	for _, item := range obj.(*v1alpha1.KconfigList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested kconfigs.
func (c *FakeKconfigs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(kconfigsResource, c.ns, opts))

}

// Create takes the representation of a kconfig and creates it.  Returns the server's representation of the kconfig, and an error, if there is any.
func (c *FakeKconfigs) Create(kconfig *v1alpha1.Kconfig) (result *v1alpha1.Kconfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(kconfigsResource, c.ns, kconfig), &v1alpha1.Kconfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Kconfig), err
}

// Update takes the representation of a kconfig and updates it. Returns the server's representation of the kconfig, and an error, if there is any.
func (c *FakeKconfigs) Update(kconfig *v1alpha1.Kconfig) (result *v1alpha1.Kconfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(kconfigsResource, c.ns, kconfig), &v1alpha1.Kconfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Kconfig), err
}

// Delete takes name of the kconfig and deletes it. Returns an error if one occurs.
func (c *FakeKconfigs) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(kconfigsResource, c.ns, name), &v1alpha1.Kconfig{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeKconfigs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(kconfigsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.KconfigList{})
	return err
}

// Patch applies the patch and returns the patched kconfig.
func (c *FakeKconfigs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Kconfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(kconfigsResource, c.ns, name, pt, data, subresources...), &v1alpha1.Kconfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Kconfig), err
}
