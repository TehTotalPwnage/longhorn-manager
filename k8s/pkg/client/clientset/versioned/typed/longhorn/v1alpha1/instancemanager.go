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

package v1alpha1

import (
	"time"

	v1alpha1 "github.com/longhorn/longhorn-manager/k8s/pkg/apis/longhorn/v1alpha1"
	scheme "github.com/longhorn/longhorn-manager/k8s/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// InstanceManagersGetter has a method to return a InstanceManagerInterface.
// A group's client should implement this interface.
type InstanceManagersGetter interface {
	InstanceManagers(namespace string) InstanceManagerInterface
}

// InstanceManagerInterface has methods to work with InstanceManager resources.
type InstanceManagerInterface interface {
	Create(*v1alpha1.InstanceManager) (*v1alpha1.InstanceManager, error)
	Update(*v1alpha1.InstanceManager) (*v1alpha1.InstanceManager, error)
	UpdateStatus(*v1alpha1.InstanceManager) (*v1alpha1.InstanceManager, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.InstanceManager, error)
	List(opts v1.ListOptions) (*v1alpha1.InstanceManagerList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.InstanceManager, err error)
	InstanceManagerExpansion
}

// instanceManagers implements InstanceManagerInterface
type instanceManagers struct {
	client rest.Interface
	ns     string
}

// newInstanceManagers returns a InstanceManagers
func newInstanceManagers(c *LonghornV1alpha1Client, namespace string) *instanceManagers {
	return &instanceManagers{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the instanceManager, and returns the corresponding instanceManager object, and an error if there is any.
func (c *instanceManagers) Get(name string, options v1.GetOptions) (result *v1alpha1.InstanceManager, err error) {
	result = &v1alpha1.InstanceManager{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("instancemanagers").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of InstanceManagers that match those selectors.
func (c *instanceManagers) List(opts v1.ListOptions) (result *v1alpha1.InstanceManagerList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.InstanceManagerList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("instancemanagers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested instanceManagers.
func (c *instanceManagers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("instancemanagers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a instanceManager and creates it.  Returns the server's representation of the instanceManager, and an error, if there is any.
func (c *instanceManagers) Create(instanceManager *v1alpha1.InstanceManager) (result *v1alpha1.InstanceManager, err error) {
	result = &v1alpha1.InstanceManager{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("instancemanagers").
		Body(instanceManager).
		Do().
		Into(result)
	return
}

// Update takes the representation of a instanceManager and updates it. Returns the server's representation of the instanceManager, and an error, if there is any.
func (c *instanceManagers) Update(instanceManager *v1alpha1.InstanceManager) (result *v1alpha1.InstanceManager, err error) {
	result = &v1alpha1.InstanceManager{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("instancemanagers").
		Name(instanceManager.Name).
		Body(instanceManager).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *instanceManagers) UpdateStatus(instanceManager *v1alpha1.InstanceManager) (result *v1alpha1.InstanceManager, err error) {
	result = &v1alpha1.InstanceManager{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("instancemanagers").
		Name(instanceManager.Name).
		SubResource("status").
		Body(instanceManager).
		Do().
		Into(result)
	return
}

// Delete takes name of the instanceManager and deletes it. Returns an error if one occurs.
func (c *instanceManagers) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("instancemanagers").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *instanceManagers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("instancemanagers").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched instanceManager.
func (c *instanceManagers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.InstanceManager, err error) {
	result = &v1alpha1.InstanceManager{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("instancemanagers").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
