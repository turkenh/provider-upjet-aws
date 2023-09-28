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

type LifecycleHookInitParameters struct {

	// Defines the action the Auto Scaling group should take when the lifecycle hook timeout elapses or if an unexpected failure occurs. The value for this parameter can be either CONTINUE or ABANDON. The default value for this parameter is ABANDON.
	DefaultResult *string `json:"defaultResult,omitempty" tf:"default_result,omitempty"`

	// Defines the amount of time, in seconds, that can elapse before the lifecycle hook times out. When the lifecycle hook times out, Auto Scaling performs the action defined in the DefaultResult parameter
	HeartbeatTimeout *float64 `json:"heartbeatTimeout,omitempty" tf:"heartbeat_timeout,omitempty"`

	// Instance state to which you want to attach the lifecycle hook. For a list of lifecycle hook types, see describe-lifecycle-hook-types
	LifecycleTransition *string `json:"lifecycleTransition,omitempty" tf:"lifecycle_transition,omitempty"`

	// Contains additional information that you want to include any time Auto Scaling sends a message to the notification target.
	NotificationMetadata *string `json:"notificationMetadata,omitempty" tf:"notification_metadata,omitempty"`

	// ARN of the notification target that Auto Scaling will use to notify you when an instance is in the transition state for the lifecycle hook. This ARN target can be either an SQS queue or an SNS topic.
	NotificationTargetArn *string `json:"notificationTargetArn,omitempty" tf:"notification_target_arn,omitempty"`
}

type LifecycleHookObservation struct {

	// Name of the Auto Scaling group to which you want to assign the lifecycle hook
	AutoscalingGroupName *string `json:"autoscalingGroupName,omitempty" tf:"autoscaling_group_name,omitempty"`

	// Defines the action the Auto Scaling group should take when the lifecycle hook timeout elapses or if an unexpected failure occurs. The value for this parameter can be either CONTINUE or ABANDON. The default value for this parameter is ABANDON.
	DefaultResult *string `json:"defaultResult,omitempty" tf:"default_result,omitempty"`

	// Defines the amount of time, in seconds, that can elapse before the lifecycle hook times out. When the lifecycle hook times out, Auto Scaling performs the action defined in the DefaultResult parameter
	HeartbeatTimeout *float64 `json:"heartbeatTimeout,omitempty" tf:"heartbeat_timeout,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Instance state to which you want to attach the lifecycle hook. For a list of lifecycle hook types, see describe-lifecycle-hook-types
	LifecycleTransition *string `json:"lifecycleTransition,omitempty" tf:"lifecycle_transition,omitempty"`

	// Contains additional information that you want to include any time Auto Scaling sends a message to the notification target.
	NotificationMetadata *string `json:"notificationMetadata,omitempty" tf:"notification_metadata,omitempty"`

	// ARN of the notification target that Auto Scaling will use to notify you when an instance is in the transition state for the lifecycle hook. This ARN target can be either an SQS queue or an SNS topic.
	NotificationTargetArn *string `json:"notificationTargetArn,omitempty" tf:"notification_target_arn,omitempty"`

	// ARN of the IAM role that allows the Auto Scaling group to publish to the specified notification target.
	RoleArn *string `json:"roleArn,omitempty" tf:"role_arn,omitempty"`
}

type LifecycleHookParameters struct {

	// Name of the Auto Scaling group to which you want to assign the lifecycle hook
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/autoscaling/v1beta1.AutoscalingGroup
	// +kubebuilder:validation:Optional
	AutoscalingGroupName *string `json:"autoscalingGroupName,omitempty" tf:"autoscaling_group_name,omitempty"`

	// Reference to a AutoscalingGroup in autoscaling to populate autoscalingGroupName.
	// +kubebuilder:validation:Optional
	AutoscalingGroupNameRef *v1.Reference `json:"autoscalingGroupNameRef,omitempty" tf:"-"`

	// Selector for a AutoscalingGroup in autoscaling to populate autoscalingGroupName.
	// +kubebuilder:validation:Optional
	AutoscalingGroupNameSelector *v1.Selector `json:"autoscalingGroupNameSelector,omitempty" tf:"-"`

	// Defines the action the Auto Scaling group should take when the lifecycle hook timeout elapses or if an unexpected failure occurs. The value for this parameter can be either CONTINUE or ABANDON. The default value for this parameter is ABANDON.
	// +kubebuilder:validation:Optional
	DefaultResult *string `json:"defaultResult,omitempty" tf:"default_result,omitempty"`

	// Defines the amount of time, in seconds, that can elapse before the lifecycle hook times out. When the lifecycle hook times out, Auto Scaling performs the action defined in the DefaultResult parameter
	// +kubebuilder:validation:Optional
	HeartbeatTimeout *float64 `json:"heartbeatTimeout,omitempty" tf:"heartbeat_timeout,omitempty"`

	// Instance state to which you want to attach the lifecycle hook. For a list of lifecycle hook types, see describe-lifecycle-hook-types
	// +kubebuilder:validation:Optional
	LifecycleTransition *string `json:"lifecycleTransition,omitempty" tf:"lifecycle_transition,omitempty"`

	// Contains additional information that you want to include any time Auto Scaling sends a message to the notification target.
	// +kubebuilder:validation:Optional
	NotificationMetadata *string `json:"notificationMetadata,omitempty" tf:"notification_metadata,omitempty"`

	// ARN of the notification target that Auto Scaling will use to notify you when an instance is in the transition state for the lifecycle hook. This ARN target can be either an SQS queue or an SNS topic.
	// +kubebuilder:validation:Optional
	NotificationTargetArn *string `json:"notificationTargetArn,omitempty" tf:"notification_target_arn,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// ARN of the IAM role that allows the Auto Scaling group to publish to the specified notification target.
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
}

// LifecycleHookSpec defines the desired state of LifecycleHook
type LifecycleHookSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LifecycleHookParameters `json:"forProvider"`
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
	InitProvider LifecycleHookInitParameters `json:"initProvider,omitempty"`
}

// LifecycleHookStatus defines the observed state of LifecycleHook.
type LifecycleHookStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LifecycleHookObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// LifecycleHook is the Schema for the LifecycleHooks API. Provides an AutoScaling Lifecycle Hook resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type LifecycleHook struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.lifecycleTransition) || (has(self.initProvider) && has(self.initProvider.lifecycleTransition))",message="spec.forProvider.lifecycleTransition is a required parameter"
	Spec   LifecycleHookSpec   `json:"spec"`
	Status LifecycleHookStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LifecycleHookList contains a list of LifecycleHooks
type LifecycleHookList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LifecycleHook `json:"items"`
}

// Repository type metadata.
var (
	LifecycleHook_Kind             = "LifecycleHook"
	LifecycleHook_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: LifecycleHook_Kind}.String()
	LifecycleHook_KindAPIVersion   = LifecycleHook_Kind + "." + CRDGroupVersion.String()
	LifecycleHook_GroupVersionKind = CRDGroupVersion.WithKind(LifecycleHook_Kind)
)

func init() {
	SchemeBuilder.Register(&LifecycleHook{}, &LifecycleHookList{})
}
