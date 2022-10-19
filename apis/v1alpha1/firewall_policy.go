// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package v1alpha1

import (
	ackv1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// FirewallPolicySpec defines the desired state of FirewallPolicy.
//
// The firewall policy defines the behavior of a firewall using a collection
// of stateless and stateful rule groups and other settings. You can use one
// firewall policy for multiple firewalls.
//
// This, along with FirewallPolicyResponse, define the policy. You can retrieve
// all objects for a firewall policy by calling DescribeFirewallPolicy.
type FirewallPolicySpec struct {
	// A description of the firewall policy.
	Description *string `json:"description,omitempty"`
	// A complex type that contains settings for encryption of your firewall policy
	// resources.
	EncryptionConfiguration *EncryptionConfiguration `json:"encryptionConfiguration,omitempty"`
	// The rule groups and policy actions to use in the firewall policy.
	// +kubebuilder:validation:Required
	FirewallPolicy *FirewallPolicy_SDK `json:"firewallPolicy"`
	// The descriptive name of the firewall policy. You can't change the name of
	// a firewall policy after you create it.
	// +kubebuilder:validation:Required
	FirewallPolicyName *string `json:"firewallPolicyName"`
	// The key:value pairs to associate with the resource.
	Tags []*Tag `json:"tags,omitempty"`
}

// FirewallPolicyStatus defines the observed state of FirewallPolicy
type FirewallPolicyStatus struct {
	// All CRs managed by ACK have a common `Status.ACKResourceMetadata` member
	// that is used to contain resource sync state, account ownership,
	// constructed ARN for the resource
	// +kubebuilder:validation:Optional
	ACKResourceMetadata *ackv1alpha1.ResourceMetadata `json:"ackResourceMetadata"`
	// All CRS managed by ACK have a common `Status.Conditions` member that
	// contains a collection of `ackv1alpha1.Condition` objects that describe
	// the various terminal states of the CR and its backend AWS service API
	// resource
	// +kubebuilder:validation:Optional
	Conditions []*ackv1alpha1.Condition `json:"conditions"`
	// The high-level properties of a firewall policy. This, along with the FirewallPolicy,
	// define the policy. You can retrieve all objects for a firewall policy by
	// calling DescribeFirewallPolicy.
	// +kubebuilder:validation:Optional
	FirewallPolicyResponse *FirewallPolicyResponse `json:"firewallPolicyResponse,omitempty"`
	// A token used for optimistic locking. Network Firewall returns a token to
	// your requests that access the firewall policy. The token marks the state
	// of the policy resource at the time of the request.
	//
	// To make changes to the policy, you provide the token in your request. Network
	// Firewall uses the token to ensure that the policy hasn't changed since you
	// last retrieved it. If it has changed, the operation fails with an InvalidTokenException.
	// If this happens, retrieve the firewall policy again to get a current copy
	// of it with current token. Reapply your changes as needed, then try the operation
	// again using the new token.
	// +kubebuilder:validation:Optional
	UpdateToken *string `json:"updateToken,omitempty"`
}

// FirewallPolicy is the Schema for the FirewallPolicies API
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
type FirewallPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FirewallPolicySpec   `json:"spec,omitempty"`
	Status            FirewallPolicyStatus `json:"status,omitempty"`
}

// FirewallPolicyList contains a list of FirewallPolicy
// +kubebuilder:object:root=true
type FirewallPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FirewallPolicy `json:"items"`
}

func init() {
	SchemeBuilder.Register(&FirewallPolicy{}, &FirewallPolicyList{})
}
