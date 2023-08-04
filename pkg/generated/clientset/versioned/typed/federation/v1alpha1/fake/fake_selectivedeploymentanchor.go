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
	"context"

	v1alpha1 "github.com/ubombar/namespace-request-controller/pkg/apis/federation/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeSelectiveDeploymentAnchors implements SelectiveDeploymentAnchorInterface
type FakeSelectiveDeploymentAnchors struct {
	Fake *FakeFederationV1alpha1
	ns   string
}

var selectivedeploymentanchorsResource = v1alpha1.SchemeGroupVersion.WithResource("selectivedeploymentanchors")

var selectivedeploymentanchorsKind = v1alpha1.SchemeGroupVersion.WithKind("SelectiveDeploymentAnchor")

// Get takes name of the selectiveDeploymentAnchor, and returns the corresponding selectiveDeploymentAnchor object, and an error if there is any.
func (c *FakeSelectiveDeploymentAnchors) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.SelectiveDeploymentAnchor, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(selectivedeploymentanchorsResource, c.ns, name), &v1alpha1.SelectiveDeploymentAnchor{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SelectiveDeploymentAnchor), err
}

// List takes label and field selectors, and returns the list of SelectiveDeploymentAnchors that match those selectors.
func (c *FakeSelectiveDeploymentAnchors) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.SelectiveDeploymentAnchorList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(selectivedeploymentanchorsResource, selectivedeploymentanchorsKind, c.ns, opts), &v1alpha1.SelectiveDeploymentAnchorList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.SelectiveDeploymentAnchorList{ListMeta: obj.(*v1alpha1.SelectiveDeploymentAnchorList).ListMeta}
	for _, item := range obj.(*v1alpha1.SelectiveDeploymentAnchorList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested selectiveDeploymentAnchors.
func (c *FakeSelectiveDeploymentAnchors) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(selectivedeploymentanchorsResource, c.ns, opts))

}

// Create takes the representation of a selectiveDeploymentAnchor and creates it.  Returns the server's representation of the selectiveDeploymentAnchor, and an error, if there is any.
func (c *FakeSelectiveDeploymentAnchors) Create(ctx context.Context, selectiveDeploymentAnchor *v1alpha1.SelectiveDeploymentAnchor, opts v1.CreateOptions) (result *v1alpha1.SelectiveDeploymentAnchor, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(selectivedeploymentanchorsResource, c.ns, selectiveDeploymentAnchor), &v1alpha1.SelectiveDeploymentAnchor{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SelectiveDeploymentAnchor), err
}

// Update takes the representation of a selectiveDeploymentAnchor and updates it. Returns the server's representation of the selectiveDeploymentAnchor, and an error, if there is any.
func (c *FakeSelectiveDeploymentAnchors) Update(ctx context.Context, selectiveDeploymentAnchor *v1alpha1.SelectiveDeploymentAnchor, opts v1.UpdateOptions) (result *v1alpha1.SelectiveDeploymentAnchor, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(selectivedeploymentanchorsResource, c.ns, selectiveDeploymentAnchor), &v1alpha1.SelectiveDeploymentAnchor{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SelectiveDeploymentAnchor), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeSelectiveDeploymentAnchors) UpdateStatus(ctx context.Context, selectiveDeploymentAnchor *v1alpha1.SelectiveDeploymentAnchor, opts v1.UpdateOptions) (*v1alpha1.SelectiveDeploymentAnchor, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(selectivedeploymentanchorsResource, "status", c.ns, selectiveDeploymentAnchor), &v1alpha1.SelectiveDeploymentAnchor{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SelectiveDeploymentAnchor), err
}

// Delete takes name of the selectiveDeploymentAnchor and deletes it. Returns an error if one occurs.
func (c *FakeSelectiveDeploymentAnchors) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(selectivedeploymentanchorsResource, c.ns, name, opts), &v1alpha1.SelectiveDeploymentAnchor{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSelectiveDeploymentAnchors) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(selectivedeploymentanchorsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.SelectiveDeploymentAnchorList{})
	return err
}

// Patch applies the patch and returns the patched selectiveDeploymentAnchor.
func (c *FakeSelectiveDeploymentAnchors) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.SelectiveDeploymentAnchor, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(selectivedeploymentanchorsResource, c.ns, name, pt, data, subresources...), &v1alpha1.SelectiveDeploymentAnchor{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SelectiveDeploymentAnchor), err
}
