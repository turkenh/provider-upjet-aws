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

type MaintenanceWindowTargetInitParameters struct {

	// The description of the maintenance window target.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The name of the maintenance window target.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// User-provided value that will be included in any CloudWatch events raised while running tasks for these targets in this Maintenance Window.
	OwnerInformation *string `json:"ownerInformation,omitempty" tf:"owner_information,omitempty"`

	// The type of target being registered with the Maintenance Window. Possible values are INSTANCE and RESOURCE_GROUP.
	ResourceType *string `json:"resourceType,omitempty" tf:"resource_type,omitempty"`

	// The targets to register with the maintenance window. In other words, the instances to run commands on when the maintenance window runs. You can specify targets using instance IDs, resource group names, or tags that have been applied to instances. For more information about these examples formats see
	// (https://docs.aws.amazon.com/systems-manager/latest/userguide/mw-cli-tutorial-targets-examples.html)
	Targets []MaintenanceWindowTargetTargetsInitParameters `json:"targets,omitempty" tf:"targets,omitempty"`
}

type MaintenanceWindowTargetObservation struct {

	// The description of the maintenance window target.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The ID of the maintenance window target.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the maintenance window target.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// User-provided value that will be included in any CloudWatch events raised while running tasks for these targets in this Maintenance Window.
	OwnerInformation *string `json:"ownerInformation,omitempty" tf:"owner_information,omitempty"`

	// The type of target being registered with the Maintenance Window. Possible values are INSTANCE and RESOURCE_GROUP.
	ResourceType *string `json:"resourceType,omitempty" tf:"resource_type,omitempty"`

	// The targets to register with the maintenance window. In other words, the instances to run commands on when the maintenance window runs. You can specify targets using instance IDs, resource group names, or tags that have been applied to instances. For more information about these examples formats see
	// (https://docs.aws.amazon.com/systems-manager/latest/userguide/mw-cli-tutorial-targets-examples.html)
	Targets []MaintenanceWindowTargetTargetsObservation `json:"targets,omitempty" tf:"targets,omitempty"`

	// The Id of the maintenance window to register the target with.
	WindowID *string `json:"windowId,omitempty" tf:"window_id,omitempty"`
}

type MaintenanceWindowTargetParameters struct {

	// The description of the maintenance window target.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The name of the maintenance window target.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// User-provided value that will be included in any CloudWatch events raised while running tasks for these targets in this Maintenance Window.
	// +kubebuilder:validation:Optional
	OwnerInformation *string `json:"ownerInformation,omitempty" tf:"owner_information,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The type of target being registered with the Maintenance Window. Possible values are INSTANCE and RESOURCE_GROUP.
	// +kubebuilder:validation:Optional
	ResourceType *string `json:"resourceType,omitempty" tf:"resource_type,omitempty"`

	// The targets to register with the maintenance window. In other words, the instances to run commands on when the maintenance window runs. You can specify targets using instance IDs, resource group names, or tags that have been applied to instances. For more information about these examples formats see
	// (https://docs.aws.amazon.com/systems-manager/latest/userguide/mw-cli-tutorial-targets-examples.html)
	// +kubebuilder:validation:Optional
	Targets []MaintenanceWindowTargetTargetsParameters `json:"targets,omitempty" tf:"targets,omitempty"`

	// The Id of the maintenance window to register the target with.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ssm/v1beta1.MaintenanceWindow
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	WindowID *string `json:"windowId,omitempty" tf:"window_id,omitempty"`

	// Reference to a MaintenanceWindow in ssm to populate windowId.
	// +kubebuilder:validation:Optional
	WindowIDRef *v1.Reference `json:"windowIdRef,omitempty" tf:"-"`

	// Selector for a MaintenanceWindow in ssm to populate windowId.
	// +kubebuilder:validation:Optional
	WindowIDSelector *v1.Selector `json:"windowIdSelector,omitempty" tf:"-"`
}

type MaintenanceWindowTargetTargetsInitParameters struct {
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type MaintenanceWindowTargetTargetsObservation struct {
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type MaintenanceWindowTargetTargetsParameters struct {

	// +kubebuilder:validation:Optional
	Key *string `json:"key" tf:"key,omitempty"`

	// +kubebuilder:validation:Optional
	Values []*string `json:"values" tf:"values,omitempty"`
}

// MaintenanceWindowTargetSpec defines the desired state of MaintenanceWindowTarget
type MaintenanceWindowTargetSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     MaintenanceWindowTargetParameters `json:"forProvider"`
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
	InitProvider MaintenanceWindowTargetInitParameters `json:"initProvider,omitempty"`
}

// MaintenanceWindowTargetStatus defines the observed state of MaintenanceWindowTarget.
type MaintenanceWindowTargetStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        MaintenanceWindowTargetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// MaintenanceWindowTarget is the Schema for the MaintenanceWindowTargets API. Provides an SSM Maintenance Window Target resource
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type MaintenanceWindowTarget struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.resourceType) || (has(self.initProvider) && has(self.initProvider.resourceType))",message="spec.forProvider.resourceType is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.targets) || (has(self.initProvider) && has(self.initProvider.targets))",message="spec.forProvider.targets is a required parameter"
	Spec   MaintenanceWindowTargetSpec   `json:"spec"`
	Status MaintenanceWindowTargetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MaintenanceWindowTargetList contains a list of MaintenanceWindowTargets
type MaintenanceWindowTargetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MaintenanceWindowTarget `json:"items"`
}

// Repository type metadata.
var (
	MaintenanceWindowTarget_Kind             = "MaintenanceWindowTarget"
	MaintenanceWindowTarget_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: MaintenanceWindowTarget_Kind}.String()
	MaintenanceWindowTarget_KindAPIVersion   = MaintenanceWindowTarget_Kind + "." + CRDGroupVersion.String()
	MaintenanceWindowTarget_GroupVersionKind = CRDGroupVersion.WithKind(MaintenanceWindowTarget_Kind)
)

func init() {
	SchemeBuilder.Register(&MaintenanceWindowTarget{}, &MaintenanceWindowTargetList{})
}
