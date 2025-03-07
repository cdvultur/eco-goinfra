package oran

import (
	"fmt"
	"testing"

	"github.com/openshift-kni/eco-goinfra/pkg/clients"
	"github.com/stretchr/testify/assert"
	"k8s.io/apimachinery/pkg/labels"
	runtimeclient "sigs.k8s.io/controller-runtime/pkg/client"
)

func TestListNodePools(t *testing.T) {
	testCases := []struct {
		nodePools     []*NodePoolBuilder
		listOptions   []runtimeclient.ListOptions
		client        bool
		expectedError error
	}{
		{
			nodePools:     []*NodePoolBuilder{buildValidNodePoolTestBuilder(buildTestClientWithDummyNodePool())},
			listOptions:   nil,
			client:        true,
			expectedError: nil,
		},
		{
			nodePools:     []*NodePoolBuilder{buildValidNodePoolTestBuilder(buildTestClientWithDummyNodePool())},
			listOptions:   []runtimeclient.ListOptions{{LabelSelector: labels.NewSelector()}},
			client:        true,
			expectedError: nil,
		},
		{
			nodePools: []*NodePoolBuilder{buildValidNodePoolTestBuilder(buildTestClientWithDummyNodePool())},
			listOptions: []runtimeclient.ListOptions{
				{LabelSelector: labels.NewSelector()},
				{LabelSelector: labels.NewSelector()},
			},
			client:        true,
			expectedError: fmt.Errorf("error: more than one ListOptions was passed"),
		},
		{
			nodePools:     []*NodePoolBuilder{buildValidNodePoolTestBuilder(buildTestClientWithDummyNodePool())},
			listOptions:   nil,
			client:        false,
			expectedError: fmt.Errorf("failed to list nodePools, 'apiClient' parameter is nil"),
		},
	}

	for _, testCase := range testCases {
		var testSettings *clients.Settings

		if testCase.client {
			testSettings = buildTestClientWithDummyNodePool()
		}

		builders, err := ListNodePools(testSettings, testCase.listOptions...)
		assert.Equal(t, testCase.expectedError, err)

		if testCase.expectedError == nil && len(testCase.listOptions) == 0 {
			assert.Equal(t, len(testCase.nodePools), len(builders))
		}
	}
}
