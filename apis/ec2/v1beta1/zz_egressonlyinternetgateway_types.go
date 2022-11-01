/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type EgressOnlyInternetGatewayObservation struct {

	// The ID of the egress-only Internet gateway.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type EgressOnlyInternetGatewayParameters struct {

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The VPC ID to create in.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.VPC
	// +kubebuilder:validation:Optional
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	// Reference to a VPC in ec2 to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDRef *v1.Reference `json:"vpcIdRef,omitempty" tf:"-"`

	// Selector for a VPC in ec2 to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDSelector *v1.Selector `json:"vpcIdSelector,omitempty" tf:"-"`
}

// EgressOnlyInternetGatewaySpec defines the desired state of EgressOnlyInternetGateway
type EgressOnlyInternetGatewaySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     EgressOnlyInternetGatewayParameters `json:"forProvider"`
}

// EgressOnlyInternetGatewayStatus defines the observed state of EgressOnlyInternetGateway.
type EgressOnlyInternetGatewayStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        EgressOnlyInternetGatewayObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// EgressOnlyInternetGateway is the Schema for the EgressOnlyInternetGateways API. Provides a resource to create an egress-only Internet gateway.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type EgressOnlyInternetGateway struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EgressOnlyInternetGatewaySpec   `json:"spec"`
	Status            EgressOnlyInternetGatewayStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EgressOnlyInternetGatewayList contains a list of EgressOnlyInternetGateways
type EgressOnlyInternetGatewayList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EgressOnlyInternetGateway `json:"items"`
}

// Repository type metadata.
var (
	EgressOnlyInternetGateway_Kind             = "EgressOnlyInternetGateway"
	EgressOnlyInternetGateway_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: EgressOnlyInternetGateway_Kind}.String()
	EgressOnlyInternetGateway_KindAPIVersion   = EgressOnlyInternetGateway_Kind + "." + CRDGroupVersion.String()
	EgressOnlyInternetGateway_GroupVersionKind = CRDGroupVersion.WithKind(EgressOnlyInternetGateway_Kind)
)

func init() {
	SchemeBuilder.Register(&EgressOnlyInternetGateway{}, &EgressOnlyInternetGatewayList{})
}
