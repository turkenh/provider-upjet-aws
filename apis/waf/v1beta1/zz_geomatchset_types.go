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

type GeoMatchConstraintInitParameters struct {

	// The type of geographical area you want AWS WAF to search for. Currently Country is the only valid value.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// The country that you want AWS WAF to search for.
	// This is the two-letter country code, e.g., US, CA, RU, CN, etc.
	// See docs for all supported values.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type GeoMatchConstraintObservation struct {

	// The type of geographical area you want AWS WAF to search for. Currently Country is the only valid value.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// The country that you want AWS WAF to search for.
	// This is the two-letter country code, e.g., US, CA, RU, CN, etc.
	// See docs for all supported values.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type GeoMatchConstraintParameters struct {

	// The type of geographical area you want AWS WAF to search for. Currently Country is the only valid value.
	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`

	// The country that you want AWS WAF to search for.
	// This is the two-letter country code, e.g., US, CA, RU, CN, etc.
	// See docs for all supported values.
	// +kubebuilder:validation:Optional
	Value *string `json:"value" tf:"value,omitempty"`
}

type GeoMatchSetInitParameters struct {

	// The GeoMatchConstraint objects which contain the country that you want AWS WAF to search for.
	GeoMatchConstraint []GeoMatchConstraintInitParameters `json:"geoMatchConstraint,omitempty" tf:"geo_match_constraint,omitempty"`

	// The name or description of the GeoMatchSet.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type GeoMatchSetObservation struct {

	// Amazon Resource Name (ARN)
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The GeoMatchConstraint objects which contain the country that you want AWS WAF to search for.
	GeoMatchConstraint []GeoMatchConstraintObservation `json:"geoMatchConstraint,omitempty" tf:"geo_match_constraint,omitempty"`

	// The ID of the WAF GeoMatchSet.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name or description of the GeoMatchSet.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type GeoMatchSetParameters struct {

	// The GeoMatchConstraint objects which contain the country that you want AWS WAF to search for.
	// +kubebuilder:validation:Optional
	GeoMatchConstraint []GeoMatchConstraintParameters `json:"geoMatchConstraint,omitempty" tf:"geo_match_constraint,omitempty"`

	// The name or description of the GeoMatchSet.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// GeoMatchSetSpec defines the desired state of GeoMatchSet
type GeoMatchSetSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     GeoMatchSetParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider GeoMatchSetInitParameters `json:"initProvider,omitempty"`
}

// GeoMatchSetStatus defines the observed state of GeoMatchSet.
type GeoMatchSetStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        GeoMatchSetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// GeoMatchSet is the Schema for the GeoMatchSets API. Provides a AWS WAF GeoMatchSet resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type GeoMatchSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   GeoMatchSetSpec   `json:"spec"`
	Status GeoMatchSetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GeoMatchSetList contains a list of GeoMatchSets
type GeoMatchSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GeoMatchSet `json:"items"`
}

// Repository type metadata.
var (
	GeoMatchSet_Kind             = "GeoMatchSet"
	GeoMatchSet_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: GeoMatchSet_Kind}.String()
	GeoMatchSet_KindAPIVersion   = GeoMatchSet_Kind + "." + CRDGroupVersion.String()
	GeoMatchSet_GroupVersionKind = CRDGroupVersion.WithKind(GeoMatchSet_Kind)
)

func init() {
	SchemeBuilder.Register(&GeoMatchSet{}, &GeoMatchSetList{})
}
