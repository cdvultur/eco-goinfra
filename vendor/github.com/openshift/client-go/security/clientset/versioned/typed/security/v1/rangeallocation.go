// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"

	v1 "github.com/openshift/api/security/v1"
	securityv1 "github.com/openshift/client-go/security/applyconfigurations/security/v1"
	scheme "github.com/openshift/client-go/security/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// RangeAllocationsGetter has a method to return a RangeAllocationInterface.
// A group's client should implement this interface.
type RangeAllocationsGetter interface {
	RangeAllocations() RangeAllocationInterface
}

// RangeAllocationInterface has methods to work with RangeAllocation resources.
type RangeAllocationInterface interface {
	Create(ctx context.Context, rangeAllocation *v1.RangeAllocation, opts metav1.CreateOptions) (*v1.RangeAllocation, error)
	Update(ctx context.Context, rangeAllocation *v1.RangeAllocation, opts metav1.UpdateOptions) (*v1.RangeAllocation, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.RangeAllocation, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.RangeAllocationList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.RangeAllocation, err error)
	Apply(ctx context.Context, rangeAllocation *securityv1.RangeAllocationApplyConfiguration, opts metav1.ApplyOptions) (result *v1.RangeAllocation, err error)
	RangeAllocationExpansion
}

// rangeAllocations implements RangeAllocationInterface
type rangeAllocations struct {
	*gentype.ClientWithListAndApply[*v1.RangeAllocation, *v1.RangeAllocationList, *securityv1.RangeAllocationApplyConfiguration]
}

// newRangeAllocations returns a RangeAllocations
func newRangeAllocations(c *SecurityV1Client) *rangeAllocations {
	return &rangeAllocations{
		gentype.NewClientWithListAndApply[*v1.RangeAllocation, *v1.RangeAllocationList, *securityv1.RangeAllocationApplyConfiguration](
			"rangeallocations",
			c.RESTClient(),
			scheme.ParameterCodec,
			"",
			func() *v1.RangeAllocation { return &v1.RangeAllocation{} },
			func() *v1.RangeAllocationList { return &v1.RangeAllocationList{} }),
	}
}
