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

type ResourceInitParameters struct {

	// JSON string matching the CloudFormation resource type schema with desired configuration.
	DesiredState *string `json:"desiredState,omitempty" tf:"desired_state,omitempty"`

	// CloudFormation resource type name. For example, AWS::EC2::VPC.
	TypeName *string `json:"typeName,omitempty" tf:"type_name,omitempty"`

	// Identifier of the CloudFormation resource type version.
	TypeVersionID *string `json:"typeVersionId,omitempty" tf:"type_version_id,omitempty"`
}

type ResourceObservation struct {

	// JSON string matching the CloudFormation resource type schema with desired configuration.
	DesiredState *string `json:"desiredState,omitempty" tf:"desired_state,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// JSON string matching the CloudFormation resource type schema with current configuration. Underlying attributes can be referenced via the jsondecode() function, for example, jsondecode(data.aws_cloudcontrolapi_resource.example.properties)["example"].
	Properties *string `json:"properties,omitempty" tf:"properties,omitempty"`

	// Amazon Resource Name (ARN) of the IAM Role to assume for operations.
	RoleArn *string `json:"roleArn,omitempty" tf:"role_arn,omitempty"`

	// CloudFormation resource type name. For example, AWS::EC2::VPC.
	TypeName *string `json:"typeName,omitempty" tf:"type_name,omitempty"`

	// Identifier of the CloudFormation resource type version.
	TypeVersionID *string `json:"typeVersionId,omitempty" tf:"type_version_id,omitempty"`
}

type ResourceParameters struct {

	// JSON string matching the CloudFormation resource type schema with desired configuration.
	// +kubebuilder:validation:Optional
	DesiredState *string `json:"desiredState,omitempty" tf:"desired_state,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Amazon Resource Name (ARN) of the IAM Role to assume for operations.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.ARNExtractor()
	// +kubebuilder:validation:Optional
	RoleArn *string `json:"roleArn,omitempty" tf:"role_arn,omitempty"`

	// Reference to a Role in iam to populate roleArn.
	// +kubebuilder:validation:Optional
	RoleArnRef *v1.Reference `json:"roleArnRef,omitempty" tf:"-"`

	// Selector for a Role in iam to populate roleArn.
	// +kubebuilder:validation:Optional
	RoleArnSelector *v1.Selector `json:"roleArnSelector,omitempty" tf:"-"`

	// JSON string of the CloudFormation resource type schema which is used for plan time validation where possible. Automatically fetched if not provided. In large scale environments with multiple resources using the same type_name, it is recommended to fetch the schema once via the aws_cloudformation_type data source and use this argument to reduce DescribeType API operation throttling. This value is marked sensitive only to prevent large plan differences from showing.
	// +kubebuilder:validation:Optional
	SchemaSecretRef *v1.SecretKeySelector `json:"schemaSecretRef,omitempty" tf:"-"`

	// CloudFormation resource type name. For example, AWS::EC2::VPC.
	// +kubebuilder:validation:Optional
	TypeName *string `json:"typeName,omitempty" tf:"type_name,omitempty"`

	// Identifier of the CloudFormation resource type version.
	// +kubebuilder:validation:Optional
	TypeVersionID *string `json:"typeVersionId,omitempty" tf:"type_version_id,omitempty"`
}

// ResourceSpec defines the desired state of Resource
type ResourceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ResourceParameters `json:"forProvider"`
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
	InitProvider ResourceInitParameters `json:"initProvider,omitempty"`
}

// ResourceStatus defines the observed state of Resource.
type ResourceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ResourceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Resource is the Schema for the Resources API. Manages a Cloud Control API Resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Resource struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.desiredState) || (has(self.initProvider) && has(self.initProvider.desiredState))",message="spec.forProvider.desiredState is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.typeName) || (has(self.initProvider) && has(self.initProvider.typeName))",message="spec.forProvider.typeName is a required parameter"
	Spec   ResourceSpec   `json:"spec"`
	Status ResourceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ResourceList contains a list of Resources
type ResourceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Resource `json:"items"`
}

// Repository type metadata.
var (
	Resource_Kind             = "Resource"
	Resource_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Resource_Kind}.String()
	Resource_KindAPIVersion   = Resource_Kind + "." + CRDGroupVersion.String()
	Resource_GroupVersionKind = CRDGroupVersion.WithKind(Resource_Kind)
)

func init() {
	SchemeBuilder.Register(&Resource{}, &ResourceList{})
}
