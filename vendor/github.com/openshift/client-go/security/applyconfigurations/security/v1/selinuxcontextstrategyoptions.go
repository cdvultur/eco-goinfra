// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/openshift/api/security/v1"
	corev1 "k8s.io/api/core/v1"
)

// SELinuxContextStrategyOptionsApplyConfiguration represents a declarative configuration of the SELinuxContextStrategyOptions type for use
// with apply.
type SELinuxContextStrategyOptionsApplyConfiguration struct {
	Type           *v1.SELinuxContextStrategyType `json:"type,omitempty"`
	SELinuxOptions *corev1.SELinuxOptions         `json:"seLinuxOptions,omitempty"`
}

// SELinuxContextStrategyOptionsApplyConfiguration constructs a declarative configuration of the SELinuxContextStrategyOptions type for use with
// apply.
func SELinuxContextStrategyOptions() *SELinuxContextStrategyOptionsApplyConfiguration {
	return &SELinuxContextStrategyOptionsApplyConfiguration{}
}

// WithType sets the Type field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Type field is set to the value of the last call.
func (b *SELinuxContextStrategyOptionsApplyConfiguration) WithType(value v1.SELinuxContextStrategyType) *SELinuxContextStrategyOptionsApplyConfiguration {
	b.Type = &value
	return b
}

// WithSELinuxOptions sets the SELinuxOptions field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SELinuxOptions field is set to the value of the last call.
func (b *SELinuxContextStrategyOptionsApplyConfiguration) WithSELinuxOptions(value corev1.SELinuxOptions) *SELinuxContextStrategyOptionsApplyConfiguration {
	b.SELinuxOptions = &value
	return b
}
