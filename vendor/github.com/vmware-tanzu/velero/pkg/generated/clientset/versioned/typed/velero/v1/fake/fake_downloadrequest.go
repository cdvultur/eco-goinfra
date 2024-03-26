/*
Copyright the Velero contributors.

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

	velerov1 "github.com/vmware-tanzu/velero/pkg/apis/velero/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeDownloadRequests implements DownloadRequestInterface
type FakeDownloadRequests struct {
	Fake *FakeVeleroV1
	ns   string
}

var downloadrequestsResource = schema.GroupVersionResource{Group: "velero.io", Version: "v1", Resource: "downloadrequests"}

var downloadrequestsKind = schema.GroupVersionKind{Group: "velero.io", Version: "v1", Kind: "DownloadRequest"}

// Get takes name of the downloadRequest, and returns the corresponding downloadRequest object, and an error if there is any.
func (c *FakeDownloadRequests) Get(ctx context.Context, name string, options v1.GetOptions) (result *velerov1.DownloadRequest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(downloadrequestsResource, c.ns, name), &velerov1.DownloadRequest{})

	if obj == nil {
		return nil, err
	}
	return obj.(*velerov1.DownloadRequest), err
}

// List takes label and field selectors, and returns the list of DownloadRequests that match those selectors.
func (c *FakeDownloadRequests) List(ctx context.Context, opts v1.ListOptions) (result *velerov1.DownloadRequestList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(downloadrequestsResource, downloadrequestsKind, c.ns, opts), &velerov1.DownloadRequestList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &velerov1.DownloadRequestList{ListMeta: obj.(*velerov1.DownloadRequestList).ListMeta}
	for _, item := range obj.(*velerov1.DownloadRequestList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested downloadRequests.
func (c *FakeDownloadRequests) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(downloadrequestsResource, c.ns, opts))

}

// Create takes the representation of a downloadRequest and creates it.  Returns the server's representation of the downloadRequest, and an error, if there is any.
func (c *FakeDownloadRequests) Create(ctx context.Context, downloadRequest *velerov1.DownloadRequest, opts v1.CreateOptions) (result *velerov1.DownloadRequest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(downloadrequestsResource, c.ns, downloadRequest), &velerov1.DownloadRequest{})

	if obj == nil {
		return nil, err
	}
	return obj.(*velerov1.DownloadRequest), err
}

// Update takes the representation of a downloadRequest and updates it. Returns the server's representation of the downloadRequest, and an error, if there is any.
func (c *FakeDownloadRequests) Update(ctx context.Context, downloadRequest *velerov1.DownloadRequest, opts v1.UpdateOptions) (result *velerov1.DownloadRequest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(downloadrequestsResource, c.ns, downloadRequest), &velerov1.DownloadRequest{})

	if obj == nil {
		return nil, err
	}
	return obj.(*velerov1.DownloadRequest), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDownloadRequests) UpdateStatus(ctx context.Context, downloadRequest *velerov1.DownloadRequest, opts v1.UpdateOptions) (*velerov1.DownloadRequest, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(downloadrequestsResource, "status", c.ns, downloadRequest), &velerov1.DownloadRequest{})

	if obj == nil {
		return nil, err
	}
	return obj.(*velerov1.DownloadRequest), err
}

// Delete takes name of the downloadRequest and deletes it. Returns an error if one occurs.
func (c *FakeDownloadRequests) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(downloadrequestsResource, c.ns, name), &velerov1.DownloadRequest{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDownloadRequests) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(downloadrequestsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &velerov1.DownloadRequestList{})
	return err
}

// Patch applies the patch and returns the patched downloadRequest.
func (c *FakeDownloadRequests) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *velerov1.DownloadRequest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(downloadrequestsResource, c.ns, name, pt, data, subresources...), &velerov1.DownloadRequest{})

	if obj == nil {
		return nil, err
	}
	return obj.(*velerov1.DownloadRequest), err
}
