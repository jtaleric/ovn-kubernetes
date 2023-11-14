/*


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

	egressservicev1 "github.com/ovn-org/ovn-kubernetes/go-controller/pkg/crd/egressservice/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeEgressServices implements EgressServiceInterface
type FakeEgressServices struct {
	Fake *FakeK8sV1
	ns   string
}

var egressservicesResource = schema.GroupVersionResource{Group: "k8s.ovn.org", Version: "v1", Resource: "egressservices"}

var egressservicesKind = schema.GroupVersionKind{Group: "k8s.ovn.org", Version: "v1", Kind: "EgressService"}

// Get takes name of the egressService, and returns the corresponding egressService object, and an error if there is any.
func (c *FakeEgressServices) Get(ctx context.Context, name string, options v1.GetOptions) (result *egressservicev1.EgressService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(egressservicesResource, c.ns, name), &egressservicev1.EgressService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*egressservicev1.EgressService), err
}

// List takes label and field selectors, and returns the list of EgressServices that match those selectors.
func (c *FakeEgressServices) List(ctx context.Context, opts v1.ListOptions) (result *egressservicev1.EgressServiceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(egressservicesResource, egressservicesKind, c.ns, opts), &egressservicev1.EgressServiceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &egressservicev1.EgressServiceList{ListMeta: obj.(*egressservicev1.EgressServiceList).ListMeta}
	for _, item := range obj.(*egressservicev1.EgressServiceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested egressServices.
func (c *FakeEgressServices) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(egressservicesResource, c.ns, opts))

}

// Create takes the representation of a egressService and creates it.  Returns the server's representation of the egressService, and an error, if there is any.
func (c *FakeEgressServices) Create(ctx context.Context, egressService *egressservicev1.EgressService, opts v1.CreateOptions) (result *egressservicev1.EgressService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(egressservicesResource, c.ns, egressService), &egressservicev1.EgressService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*egressservicev1.EgressService), err
}

// Update takes the representation of a egressService and updates it. Returns the server's representation of the egressService, and an error, if there is any.
func (c *FakeEgressServices) Update(ctx context.Context, egressService *egressservicev1.EgressService, opts v1.UpdateOptions) (result *egressservicev1.EgressService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(egressservicesResource, c.ns, egressService), &egressservicev1.EgressService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*egressservicev1.EgressService), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeEgressServices) UpdateStatus(ctx context.Context, egressService *egressservicev1.EgressService, opts v1.UpdateOptions) (*egressservicev1.EgressService, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(egressservicesResource, "status", c.ns, egressService), &egressservicev1.EgressService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*egressservicev1.EgressService), err
}

// Delete takes name of the egressService and deletes it. Returns an error if one occurs.
func (c *FakeEgressServices) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(egressservicesResource, c.ns, name), &egressservicev1.EgressService{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeEgressServices) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(egressservicesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &egressservicev1.EgressServiceList{})
	return err
}

// Patch applies the patch and returns the patched egressService.
func (c *FakeEgressServices) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *egressservicev1.EgressService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(egressservicesResource, c.ns, name, pt, data, subresources...), &egressservicev1.EgressService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*egressservicev1.EgressService), err
}
