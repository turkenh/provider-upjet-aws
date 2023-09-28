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

type DestinationConfigOnFailureInitParameters struct {
}

type DestinationConfigOnFailureObservation struct {

	// Amazon Resource Name (ARN) of the destination resource. See the Lambda Developer Guide for acceptable resource types and associated IAM permissions.
	Destination *string `json:"destination,omitempty" tf:"destination,omitempty"`
}

type DestinationConfigOnFailureParameters struct {

	// Amazon Resource Name (ARN) of the destination resource. See the Lambda Developer Guide for acceptable resource types and associated IAM permissions.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/sqs/v1beta1.Queue
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.ARNExtractor()
	// +kubebuilder:validation:Optional
	Destination *string `json:"destination,omitempty" tf:"destination,omitempty"`

	// Reference to a Queue in sqs to populate destination.
	// +kubebuilder:validation:Optional
	DestinationRef *v1.Reference `json:"destinationRef,omitempty" tf:"-"`

	// Selector for a Queue in sqs to populate destination.
	// +kubebuilder:validation:Optional
	DestinationSelector *v1.Selector `json:"destinationSelector,omitempty" tf:"-"`
}

type FunctionEventInvokeConfigDestinationConfigInitParameters struct {

	// Configuration block with destination configuration for failed asynchronous invocations. See below for details.
	OnFailure []DestinationConfigOnFailureInitParameters `json:"onFailure,omitempty" tf:"on_failure,omitempty"`

	// Configuration block with destination configuration for successful asynchronous invocations. See below for details.
	OnSuccess []OnSuccessInitParameters `json:"onSuccess,omitempty" tf:"on_success,omitempty"`
}

type FunctionEventInvokeConfigDestinationConfigObservation struct {

	// Configuration block with destination configuration for failed asynchronous invocations. See below for details.
	OnFailure []DestinationConfigOnFailureObservation `json:"onFailure,omitempty" tf:"on_failure,omitempty"`

	// Configuration block with destination configuration for successful asynchronous invocations. See below for details.
	OnSuccess []OnSuccessObservation `json:"onSuccess,omitempty" tf:"on_success,omitempty"`
}

type FunctionEventInvokeConfigDestinationConfigParameters struct {

	// Configuration block with destination configuration for failed asynchronous invocations. See below for details.
	// +kubebuilder:validation:Optional
	OnFailure []DestinationConfigOnFailureParameters `json:"onFailure,omitempty" tf:"on_failure,omitempty"`

	// Configuration block with destination configuration for successful asynchronous invocations. See below for details.
	// +kubebuilder:validation:Optional
	OnSuccess []OnSuccessParameters `json:"onSuccess,omitempty" tf:"on_success,omitempty"`
}

type FunctionEventInvokeConfigInitParameters struct {

	// Configuration block with destination configuration. See below for details.
	DestinationConfig []FunctionEventInvokeConfigDestinationConfigInitParameters `json:"destinationConfig,omitempty" tf:"destination_config,omitempty"`

	// Name or Amazon Resource Name (ARN) of the Lambda Function, omitting any version or alias qualifier.
	FunctionName *string `json:"functionName,omitempty" tf:"function_name,omitempty"`

	// Maximum age of a request that Lambda sends to a function for processing in seconds. Valid values between 60 and 21600.
	MaximumEventAgeInSeconds *float64 `json:"maximumEventAgeInSeconds,omitempty" tf:"maximum_event_age_in_seconds,omitempty"`

	// Maximum number of times to retry when the function returns an error. Valid values between 0 and 2. Defaults to 2.
	MaximumRetryAttempts *float64 `json:"maximumRetryAttempts,omitempty" tf:"maximum_retry_attempts,omitempty"`

	// Lambda Function published version, $LATEST, or Lambda Alias name.
	Qualifier *string `json:"qualifier,omitempty" tf:"qualifier,omitempty"`
}

type FunctionEventInvokeConfigObservation struct {

	// Configuration block with destination configuration. See below for details.
	DestinationConfig []FunctionEventInvokeConfigDestinationConfigObservation `json:"destinationConfig,omitempty" tf:"destination_config,omitempty"`

	// Name or Amazon Resource Name (ARN) of the Lambda Function, omitting any version or alias qualifier.
	FunctionName *string `json:"functionName,omitempty" tf:"function_name,omitempty"`

	// Fully qualified Lambda Function name or Amazon Resource Name (ARN)
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Maximum age of a request that Lambda sends to a function for processing in seconds. Valid values between 60 and 21600.
	MaximumEventAgeInSeconds *float64 `json:"maximumEventAgeInSeconds,omitempty" tf:"maximum_event_age_in_seconds,omitempty"`

	// Maximum number of times to retry when the function returns an error. Valid values between 0 and 2. Defaults to 2.
	MaximumRetryAttempts *float64 `json:"maximumRetryAttempts,omitempty" tf:"maximum_retry_attempts,omitempty"`

	// Lambda Function published version, $LATEST, or Lambda Alias name.
	Qualifier *string `json:"qualifier,omitempty" tf:"qualifier,omitempty"`
}

type FunctionEventInvokeConfigParameters struct {

	// Configuration block with destination configuration. See below for details.
	// +kubebuilder:validation:Optional
	DestinationConfig []FunctionEventInvokeConfigDestinationConfigParameters `json:"destinationConfig,omitempty" tf:"destination_config,omitempty"`

	// Name or Amazon Resource Name (ARN) of the Lambda Function, omitting any version or alias qualifier.
	// +kubebuilder:validation:Optional
	FunctionName *string `json:"functionName,omitempty" tf:"function_name,omitempty"`

	// Maximum age of a request that Lambda sends to a function for processing in seconds. Valid values between 60 and 21600.
	// +kubebuilder:validation:Optional
	MaximumEventAgeInSeconds *float64 `json:"maximumEventAgeInSeconds,omitempty" tf:"maximum_event_age_in_seconds,omitempty"`

	// Maximum number of times to retry when the function returns an error. Valid values between 0 and 2. Defaults to 2.
	// +kubebuilder:validation:Optional
	MaximumRetryAttempts *float64 `json:"maximumRetryAttempts,omitempty" tf:"maximum_retry_attempts,omitempty"`

	// Lambda Function published version, $LATEST, or Lambda Alias name.
	// +kubebuilder:validation:Optional
	Qualifier *string `json:"qualifier,omitempty" tf:"qualifier,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

type OnSuccessInitParameters struct {
}

type OnSuccessObservation struct {

	// Amazon Resource Name (ARN) of the destination resource. See the Lambda Developer Guide for acceptable resource types and associated IAM permissions.
	Destination *string `json:"destination,omitempty" tf:"destination,omitempty"`
}

type OnSuccessParameters struct {

	// Amazon Resource Name (ARN) of the destination resource. See the Lambda Developer Guide for acceptable resource types and associated IAM permissions.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/sns/v1beta1.Topic
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.ARNExtractor()
	// +kubebuilder:validation:Optional
	Destination *string `json:"destination,omitempty" tf:"destination,omitempty"`

	// Reference to a Topic in sns to populate destination.
	// +kubebuilder:validation:Optional
	DestinationRef *v1.Reference `json:"destinationRef,omitempty" tf:"-"`

	// Selector for a Topic in sns to populate destination.
	// +kubebuilder:validation:Optional
	DestinationSelector *v1.Selector `json:"destinationSelector,omitempty" tf:"-"`
}

// FunctionEventInvokeConfigSpec defines the desired state of FunctionEventInvokeConfig
type FunctionEventInvokeConfigSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FunctionEventInvokeConfigParameters `json:"forProvider"`
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
	InitProvider FunctionEventInvokeConfigInitParameters `json:"initProvider,omitempty"`
}

// FunctionEventInvokeConfigStatus defines the observed state of FunctionEventInvokeConfig.
type FunctionEventInvokeConfigStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FunctionEventInvokeConfigObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// FunctionEventInvokeConfig is the Schema for the FunctionEventInvokeConfigs API. Manages an asynchronous invocation configuration for a Lambda Function or Alias.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type FunctionEventInvokeConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.functionName) || (has(self.initProvider) && has(self.initProvider.functionName))",message="spec.forProvider.functionName is a required parameter"
	Spec   FunctionEventInvokeConfigSpec   `json:"spec"`
	Status FunctionEventInvokeConfigStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FunctionEventInvokeConfigList contains a list of FunctionEventInvokeConfigs
type FunctionEventInvokeConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FunctionEventInvokeConfig `json:"items"`
}

// Repository type metadata.
var (
	FunctionEventInvokeConfig_Kind             = "FunctionEventInvokeConfig"
	FunctionEventInvokeConfig_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: FunctionEventInvokeConfig_Kind}.String()
	FunctionEventInvokeConfig_KindAPIVersion   = FunctionEventInvokeConfig_Kind + "." + CRDGroupVersion.String()
	FunctionEventInvokeConfig_GroupVersionKind = CRDGroupVersion.WithKind(FunctionEventInvokeConfig_Kind)
)

func init() {
	SchemeBuilder.Register(&FunctionEventInvokeConfig{}, &FunctionEventInvokeConfigList{})
}
