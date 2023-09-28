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

type ActionThresholdInitParameters struct {

	// The type of threshold for a notification. Valid values are PERCENTAGE or ABSOLUTE_VALUE.
	ActionThresholdType *string `json:"actionThresholdType,omitempty" tf:"action_threshold_type,omitempty"`

	// The threshold of a notification.
	ActionThresholdValue *float64 `json:"actionThresholdValue,omitempty" tf:"action_threshold_value,omitempty"`
}

type ActionThresholdObservation struct {

	// The type of threshold for a notification. Valid values are PERCENTAGE or ABSOLUTE_VALUE.
	ActionThresholdType *string `json:"actionThresholdType,omitempty" tf:"action_threshold_type,omitempty"`

	// The threshold of a notification.
	ActionThresholdValue *float64 `json:"actionThresholdValue,omitempty" tf:"action_threshold_value,omitempty"`
}

type ActionThresholdParameters struct {

	// The type of threshold for a notification. Valid values are PERCENTAGE or ABSOLUTE_VALUE.
	// +kubebuilder:validation:Optional
	ActionThresholdType *string `json:"actionThresholdType" tf:"action_threshold_type,omitempty"`

	// The threshold of a notification.
	// +kubebuilder:validation:Optional
	ActionThresholdValue *float64 `json:"actionThresholdValue" tf:"action_threshold_value,omitempty"`
}

type BudgetActionInitParameters struct {

	// The ID of the target account for budget. Will use current user's account_id by default if omitted.
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// The trigger threshold of the action. See Action Threshold.
	ActionThreshold []ActionThresholdInitParameters `json:"actionThreshold,omitempty" tf:"action_threshold,omitempty"`

	// The type of action. This defines the type of tasks that can be carried out by this action. This field also determines the format for definition. Valid values are APPLY_IAM_POLICY, APPLY_SCP_POLICY, and RUN_SSM_DOCUMENTS.
	ActionType *string `json:"actionType,omitempty" tf:"action_type,omitempty"`

	// This specifies if the action needs manual or automatic approval. Valid values are AUTOMATIC and MANUAL.
	ApprovalModel *string `json:"approvalModel,omitempty" tf:"approval_model,omitempty"`

	// Specifies all of the type-specific parameters. See Definition.
	Definition []DefinitionInitParameters `json:"definition,omitempty" tf:"definition,omitempty"`

	// The type of a notification. Valid values are ACTUAL or FORECASTED.
	NotificationType *string `json:"notificationType,omitempty" tf:"notification_type,omitempty"`

	// A list of subscribers. See Subscriber.
	Subscriber []SubscriberInitParameters `json:"subscriber,omitempty" tf:"subscriber,omitempty"`
}

type BudgetActionObservation struct {

	// The ID of the target account for budget. Will use current user's account_id by default if omitted.
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// The id of the budget action.
	ActionID *string `json:"actionId,omitempty" tf:"action_id,omitempty"`

	// The trigger threshold of the action. See Action Threshold.
	ActionThreshold []ActionThresholdObservation `json:"actionThreshold,omitempty" tf:"action_threshold,omitempty"`

	// The type of action. This defines the type of tasks that can be carried out by this action. This field also determines the format for definition. Valid values are APPLY_IAM_POLICY, APPLY_SCP_POLICY, and RUN_SSM_DOCUMENTS.
	ActionType *string `json:"actionType,omitempty" tf:"action_type,omitempty"`

	// This specifies if the action needs manual or automatic approval. Valid values are AUTOMATIC and MANUAL.
	ApprovalModel *string `json:"approvalModel,omitempty" tf:"approval_model,omitempty"`

	// The ARN of the budget action.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The name of a budget.
	BudgetName *string `json:"budgetName,omitempty" tf:"budget_name,omitempty"`

	// Specifies all of the type-specific parameters. See Definition.
	Definition []DefinitionObservation `json:"definition,omitempty" tf:"definition,omitempty"`

	// The role passed for action execution and reversion. Roles and actions must be in the same account.
	ExecutionRoleArn *string `json:"executionRoleArn,omitempty" tf:"execution_role_arn,omitempty"`

	// ID of resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The type of a notification. Valid values are ACTUAL or FORECASTED.
	NotificationType *string `json:"notificationType,omitempty" tf:"notification_type,omitempty"`

	// The status of the budget action.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// A list of subscribers. See Subscriber.
	Subscriber []SubscriberObservation `json:"subscriber,omitempty" tf:"subscriber,omitempty"`
}

type BudgetActionParameters struct {

	// The ID of the target account for budget. Will use current user's account_id by default if omitted.
	// +kubebuilder:validation:Optional
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// The trigger threshold of the action. See Action Threshold.
	// +kubebuilder:validation:Optional
	ActionThreshold []ActionThresholdParameters `json:"actionThreshold,omitempty" tf:"action_threshold,omitempty"`

	// The type of action. This defines the type of tasks that can be carried out by this action. This field also determines the format for definition. Valid values are APPLY_IAM_POLICY, APPLY_SCP_POLICY, and RUN_SSM_DOCUMENTS.
	// +kubebuilder:validation:Optional
	ActionType *string `json:"actionType,omitempty" tf:"action_type,omitempty"`

	// This specifies if the action needs manual or automatic approval. Valid values are AUTOMATIC and MANUAL.
	// +kubebuilder:validation:Optional
	ApprovalModel *string `json:"approvalModel,omitempty" tf:"approval_model,omitempty"`

	// The name of a budget.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/budgets/v1beta1.Budget
	// +kubebuilder:validation:Optional
	BudgetName *string `json:"budgetName,omitempty" tf:"budget_name,omitempty"`

	// Reference to a Budget in budgets to populate budgetName.
	// +kubebuilder:validation:Optional
	BudgetNameRef *v1.Reference `json:"budgetNameRef,omitempty" tf:"-"`

	// Selector for a Budget in budgets to populate budgetName.
	// +kubebuilder:validation:Optional
	BudgetNameSelector *v1.Selector `json:"budgetNameSelector,omitempty" tf:"-"`

	// Specifies all of the type-specific parameters. See Definition.
	// +kubebuilder:validation:Optional
	Definition []DefinitionParameters `json:"definition,omitempty" tf:"definition,omitempty"`

	// The role passed for action execution and reversion. Roles and actions must be in the same account.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.ARNExtractor()
	// +kubebuilder:validation:Optional
	ExecutionRoleArn *string `json:"executionRoleArn,omitempty" tf:"execution_role_arn,omitempty"`

	// Reference to a Role in iam to populate executionRoleArn.
	// +kubebuilder:validation:Optional
	ExecutionRoleArnRef *v1.Reference `json:"executionRoleArnRef,omitempty" tf:"-"`

	// Selector for a Role in iam to populate executionRoleArn.
	// +kubebuilder:validation:Optional
	ExecutionRoleArnSelector *v1.Selector `json:"executionRoleArnSelector,omitempty" tf:"-"`

	// The type of a notification. Valid values are ACTUAL or FORECASTED.
	// +kubebuilder:validation:Optional
	NotificationType *string `json:"notificationType,omitempty" tf:"notification_type,omitempty"`

	// The Region to run the SSM document.
	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// A list of subscribers. See Subscriber.
	// +kubebuilder:validation:Optional
	Subscriber []SubscriberParameters `json:"subscriber,omitempty" tf:"subscriber,omitempty"`
}

type DefinitionInitParameters struct {

	// The AWS Identity and Access Management (IAM) action definition details. See IAM Action Definition.
	IAMActionDefinition []IAMActionDefinitionInitParameters `json:"iamActionDefinition,omitempty" tf:"iam_action_definition,omitempty"`

	// The service control policies (SCPs) action definition details. See SCP Action Definition.
	ScpActionDefinition []ScpActionDefinitionInitParameters `json:"scpActionDefinition,omitempty" tf:"scp_action_definition,omitempty"`

	// The AWS Systems Manager (SSM) action definition details. See SSM Action Definition.
	SsmActionDefinition []SsmActionDefinitionInitParameters `json:"ssmActionDefinition,omitempty" tf:"ssm_action_definition,omitempty"`
}

type DefinitionObservation struct {

	// The AWS Identity and Access Management (IAM) action definition details. See IAM Action Definition.
	IAMActionDefinition []IAMActionDefinitionObservation `json:"iamActionDefinition,omitempty" tf:"iam_action_definition,omitempty"`

	// The service control policies (SCPs) action definition details. See SCP Action Definition.
	ScpActionDefinition []ScpActionDefinitionObservation `json:"scpActionDefinition,omitempty" tf:"scp_action_definition,omitempty"`

	// The AWS Systems Manager (SSM) action definition details. See SSM Action Definition.
	SsmActionDefinition []SsmActionDefinitionObservation `json:"ssmActionDefinition,omitempty" tf:"ssm_action_definition,omitempty"`
}

type DefinitionParameters struct {

	// The AWS Identity and Access Management (IAM) action definition details. See IAM Action Definition.
	// +kubebuilder:validation:Optional
	IAMActionDefinition []IAMActionDefinitionParameters `json:"iamActionDefinition,omitempty" tf:"iam_action_definition,omitempty"`

	// The service control policies (SCPs) action definition details. See SCP Action Definition.
	// +kubebuilder:validation:Optional
	ScpActionDefinition []ScpActionDefinitionParameters `json:"scpActionDefinition,omitempty" tf:"scp_action_definition,omitempty"`

	// The AWS Systems Manager (SSM) action definition details. See SSM Action Definition.
	// +kubebuilder:validation:Optional
	SsmActionDefinition []SsmActionDefinitionParameters `json:"ssmActionDefinition,omitempty" tf:"ssm_action_definition,omitempty"`
}

type IAMActionDefinitionInitParameters struct {

	// A list of groups to be attached. There must be at least one group.
	Groups []*string `json:"groups,omitempty" tf:"groups,omitempty"`

	// A list of roles to be attached. There must be at least one role.
	Roles []*string `json:"roles,omitempty" tf:"roles,omitempty"`

	// A list of users to be attached. There must be at least one user.
	Users []*string `json:"users,omitempty" tf:"users,omitempty"`
}

type IAMActionDefinitionObservation struct {

	// A list of groups to be attached. There must be at least one group.
	Groups []*string `json:"groups,omitempty" tf:"groups,omitempty"`

	// The Amazon Resource Name (ARN) of the policy to be attached.
	PolicyArn *string `json:"policyArn,omitempty" tf:"policy_arn,omitempty"`

	// A list of roles to be attached. There must be at least one role.
	Roles []*string `json:"roles,omitempty" tf:"roles,omitempty"`

	// A list of users to be attached. There must be at least one user.
	Users []*string `json:"users,omitempty" tf:"users,omitempty"`
}

type IAMActionDefinitionParameters struct {

	// A list of groups to be attached. There must be at least one group.
	// +kubebuilder:validation:Optional
	Groups []*string `json:"groups,omitempty" tf:"groups,omitempty"`

	// The Amazon Resource Name (ARN) of the policy to be attached.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/iam/v1beta1.Policy
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("arn",true)
	// +kubebuilder:validation:Optional
	PolicyArn *string `json:"policyArn,omitempty" tf:"policy_arn,omitempty"`

	// Reference to a Policy in iam to populate policyArn.
	// +kubebuilder:validation:Optional
	PolicyArnRef *v1.Reference `json:"policyArnRef,omitempty" tf:"-"`

	// Selector for a Policy in iam to populate policyArn.
	// +kubebuilder:validation:Optional
	PolicyArnSelector *v1.Selector `json:"policyArnSelector,omitempty" tf:"-"`

	// A list of roles to be attached. There must be at least one role.
	// +kubebuilder:validation:Optional
	Roles []*string `json:"roles,omitempty" tf:"roles,omitempty"`

	// A list of users to be attached. There must be at least one user.
	// +kubebuilder:validation:Optional
	Users []*string `json:"users,omitempty" tf:"users,omitempty"`
}

type ScpActionDefinitionInitParameters struct {

	// The policy ID attached.
	PolicyID *string `json:"policyId,omitempty" tf:"policy_id,omitempty"`

	// A list of target IDs.
	TargetIds []*string `json:"targetIds,omitempty" tf:"target_ids,omitempty"`
}

type ScpActionDefinitionObservation struct {

	// The policy ID attached.
	PolicyID *string `json:"policyId,omitempty" tf:"policy_id,omitempty"`

	// A list of target IDs.
	TargetIds []*string `json:"targetIds,omitempty" tf:"target_ids,omitempty"`
}

type ScpActionDefinitionParameters struct {

	// The policy ID attached.
	// +kubebuilder:validation:Optional
	PolicyID *string `json:"policyId" tf:"policy_id,omitempty"`

	// A list of target IDs.
	// +kubebuilder:validation:Optional
	TargetIds []*string `json:"targetIds" tf:"target_ids,omitempty"`
}

type SsmActionDefinitionInitParameters struct {

	// The action subType. Valid values are STOP_EC2_INSTANCES or STOP_RDS_INSTANCES.
	ActionSubType *string `json:"actionSubType,omitempty" tf:"action_sub_type,omitempty"`

	// The EC2 and RDS instance IDs.
	InstanceIds []*string `json:"instanceIds,omitempty" tf:"instance_ids,omitempty"`
}

type SsmActionDefinitionObservation struct {

	// The action subType. Valid values are STOP_EC2_INSTANCES or STOP_RDS_INSTANCES.
	ActionSubType *string `json:"actionSubType,omitempty" tf:"action_sub_type,omitempty"`

	// The EC2 and RDS instance IDs.
	InstanceIds []*string `json:"instanceIds,omitempty" tf:"instance_ids,omitempty"`

	// The Region to run the SSM document.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

type SsmActionDefinitionParameters struct {

	// The action subType. Valid values are STOP_EC2_INSTANCES or STOP_RDS_INSTANCES.
	// +kubebuilder:validation:Optional
	ActionSubType *string `json:"actionSubType" tf:"action_sub_type,omitempty"`

	// The EC2 and RDS instance IDs.
	// +kubebuilder:validation:Optional
	InstanceIds []*string `json:"instanceIds" tf:"instance_ids,omitempty"`

	// The Region to run the SSM document.
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"region,omitempty"`
}

type SubscriberInitParameters struct {

	// The address that AWS sends budget notifications to, either an SNS topic or an email.
	Address *string `json:"address,omitempty" tf:"address,omitempty"`

	// The type of notification that AWS sends to a subscriber. Valid values are SNS or EMAIL.
	SubscriptionType *string `json:"subscriptionType,omitempty" tf:"subscription_type,omitempty"`
}

type SubscriberObservation struct {

	// The address that AWS sends budget notifications to, either an SNS topic or an email.
	Address *string `json:"address,omitempty" tf:"address,omitempty"`

	// The type of notification that AWS sends to a subscriber. Valid values are SNS or EMAIL.
	SubscriptionType *string `json:"subscriptionType,omitempty" tf:"subscription_type,omitempty"`
}

type SubscriberParameters struct {

	// The address that AWS sends budget notifications to, either an SNS topic or an email.
	// +kubebuilder:validation:Optional
	Address *string `json:"address" tf:"address,omitempty"`

	// The type of notification that AWS sends to a subscriber. Valid values are SNS or EMAIL.
	// +kubebuilder:validation:Optional
	SubscriptionType *string `json:"subscriptionType" tf:"subscription_type,omitempty"`
}

// BudgetActionSpec defines the desired state of BudgetAction
type BudgetActionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BudgetActionParameters `json:"forProvider"`
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
	InitProvider BudgetActionInitParameters `json:"initProvider,omitempty"`
}

// BudgetActionStatus defines the observed state of BudgetAction.
type BudgetActionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BudgetActionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// BudgetAction is the Schema for the BudgetActions API. Provides a budget action resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type BudgetAction struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.actionThreshold) || (has(self.initProvider) && has(self.initProvider.actionThreshold))",message="spec.forProvider.actionThreshold is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.actionType) || (has(self.initProvider) && has(self.initProvider.actionType))",message="spec.forProvider.actionType is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.approvalModel) || (has(self.initProvider) && has(self.initProvider.approvalModel))",message="spec.forProvider.approvalModel is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.definition) || (has(self.initProvider) && has(self.initProvider.definition))",message="spec.forProvider.definition is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.notificationType) || (has(self.initProvider) && has(self.initProvider.notificationType))",message="spec.forProvider.notificationType is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.subscriber) || (has(self.initProvider) && has(self.initProvider.subscriber))",message="spec.forProvider.subscriber is a required parameter"
	Spec   BudgetActionSpec   `json:"spec"`
	Status BudgetActionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BudgetActionList contains a list of BudgetActions
type BudgetActionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BudgetAction `json:"items"`
}

// Repository type metadata.
var (
	BudgetAction_Kind             = "BudgetAction"
	BudgetAction_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: BudgetAction_Kind}.String()
	BudgetAction_KindAPIVersion   = BudgetAction_Kind + "." + CRDGroupVersion.String()
	BudgetAction_GroupVersionKind = CRDGroupVersion.WithKind(BudgetAction_Kind)
)

func init() {
	SchemeBuilder.Register(&BudgetAction{}, &BudgetActionList{})
}
