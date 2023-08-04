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
	"context"
	"time"

	v1alpha1 "github.com/ubombar/namespace-request-controller/pkg/apis/federation/v1alpha1"
	scheme "github.com/ubombar/namespace-request-controller/pkg/generated/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// SelectiveDeploymentAnchorsGetter has a method to return a SelectiveDeploymentAnchorInterface.
// A group's client should implement this interface.
type SelectiveDeploymentAnchorsGetter interface {
	SelectiveDeploymentAnchors(namespace string) SelectiveDeploymentAnchorInterface
}

// SelectiveDeploymentAnchorInterface has methods to work with SelectiveDeploymentAnchor resources.
type SelectiveDeploymentAnchorInterface interface {
	Create(ctx context.Context, selectiveDeploymentAnchor *v1alpha1.SelectiveDeploymentAnchor, opts v1.CreateOptions) (*v1alpha1.SelectiveDeploymentAnchor, error)
	Update(ctx context.Context, selectiveDeploymentAnchor *v1alpha1.SelectiveDeploymentAnchor, opts v1.UpdateOptions) (*v1alpha1.SelectiveDeploymentAnchor, error)
	UpdateStatus(ctx context.Context, selectiveDeploymentAnchor *v1alpha1.SelectiveDeploymentAnchor, opts v1.UpdateOptions) (*v1alpha1.SelectiveDeploymentAnchor, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.SelectiveDeploymentAnchor, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.SelectiveDeploymentAnchorList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.SelectiveDeploymentAnchor, err error)
	SelectiveDeploymentAnchorExpansion
}

// selectiveDeploymentAnchors implements SelectiveDeploymentAnchorInterface
type selectiveDeploymentAnchors struct {
	client rest.Interface
	ns     string
}

// newSelectiveDeploymentAnchors returns a SelectiveDeploymentAnchors
func newSelectiveDeploymentAnchors(c *FederationV1alpha1Client, namespace string) *selectiveDeploymentAnchors {
	return &selectiveDeploymentAnchors{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the selectiveDeploymentAnchor, and returns the corresponding selectiveDeploymentAnchor object, and an error if there is any.
func (c *selectiveDeploymentAnchors) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.SelectiveDeploymentAnchor, err error) {
	result = &v1alpha1.SelectiveDeploymentAnchor{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("selectivedeploymentanchors").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of SelectiveDeploymentAnchors that match those selectors.
func (c *selectiveDeploymentAnchors) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.SelectiveDeploymentAnchorList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.SelectiveDeploymentAnchorList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("selectivedeploymentanchors").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested selectiveDeploymentAnchors.
func (c *selectiveDeploymentAnchors) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("selectivedeploymentanchors").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a selectiveDeploymentAnchor and creates it.  Returns the server's representation of the selectiveDeploymentAnchor, and an error, if there is any.
func (c *selectiveDeploymentAnchors) Create(ctx context.Context, selectiveDeploymentAnchor *v1alpha1.SelectiveDeploymentAnchor, opts v1.CreateOptions) (result *v1alpha1.SelectiveDeploymentAnchor, err error) {
	result = &v1alpha1.SelectiveDeploymentAnchor{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("selectivedeploymentanchors").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(selectiveDeploymentAnchor).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a selectiveDeploymentAnchor and updates it. Returns the server's representation of the selectiveDeploymentAnchor, and an error, if there is any.
func (c *selectiveDeploymentAnchors) Update(ctx context.Context, selectiveDeploymentAnchor *v1alpha1.SelectiveDeploymentAnchor, opts v1.UpdateOptions) (result *v1alpha1.SelectiveDeploymentAnchor, err error) {
	result = &v1alpha1.SelectiveDeploymentAnchor{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("selectivedeploymentanchors").
		Name(selectiveDeploymentAnchor.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(selectiveDeploymentAnchor).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *selectiveDeploymentAnchors) UpdateStatus(ctx context.Context, selectiveDeploymentAnchor *v1alpha1.SelectiveDeploymentAnchor, opts v1.UpdateOptions) (result *v1alpha1.SelectiveDeploymentAnchor, err error) {
	result = &v1alpha1.SelectiveDeploymentAnchor{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("selectivedeploymentanchors").
		Name(selectiveDeploymentAnchor.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(selectiveDeploymentAnchor).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the selectiveDeploymentAnchor and deletes it. Returns an error if one occurs.
func (c *selectiveDeploymentAnchors) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("selectivedeploymentanchors").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *selectiveDeploymentAnchors) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("selectivedeploymentanchors").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched selectiveDeploymentAnchor.
func (c *selectiveDeploymentAnchors) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.SelectiveDeploymentAnchor, err error) {
	result = &v1alpha1.SelectiveDeploymentAnchor{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("selectivedeploymentanchors").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
