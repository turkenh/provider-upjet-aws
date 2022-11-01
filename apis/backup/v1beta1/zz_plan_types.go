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

type AdvancedBackupSettingObservation struct {
}

type AdvancedBackupSettingParameters struct {

	// Specifies the backup option for a selected resource. This option is only available for Windows VSS backup jobs. Set to { WindowsVSS = "enabled" } to enable Windows VSS backup option and create a VSS Windows backup.
	// +kubebuilder:validation:Required
	BackupOptions map[string]*string `json:"backupOptions" tf:"backup_options,omitempty"`

	// The type of AWS resource to be backed up. For VSS Windows backups, the only supported resource type is Amazon EC2. Valid values: EC2.
	// +kubebuilder:validation:Required
	ResourceType *string `json:"resourceType" tf:"resource_type,omitempty"`
}

type CopyActionObservation struct {
}

type CopyActionParameters struct {

	// An Amazon Resource Name (ARN) that uniquely identifies the destination backup vault for the copied backup.
	// +kubebuilder:validation:Required
	DestinationVaultArn *string `json:"destinationVaultArn" tf:"destination_vault_arn,omitempty"`

	// The lifecycle defines when a protected resource is transitioned to cold storage and when it expires.  Fields documented below.
	// +kubebuilder:validation:Optional
	Lifecycle []LifecycleParameters `json:"lifecycle,omitempty" tf:"lifecycle,omitempty"`
}

type LifecycleObservation struct {
}

type LifecycleParameters struct {

	// Specifies the number of days after creation that a recovery point is moved to cold storage.
	// +kubebuilder:validation:Optional
	ColdStorageAfter *float64 `json:"coldStorageAfter,omitempty" tf:"cold_storage_after,omitempty"`

	// Specifies the number of days after creation that a recovery point is deleted. Must be 90 days greater than cold_storage_after.
	// +kubebuilder:validation:Optional
	DeleteAfter *float64 `json:"deleteAfter,omitempty" tf:"delete_after,omitempty"`
}

type PlanObservation struct {

	// The ARN of the backup plan.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The id of the backup plan.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// Unique, randomly generated, Unicode, UTF-8 encoded string that serves as the version ID of the backup plan.
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type PlanParameters struct {

	// An object that specifies backup options for each resource type.
	// +kubebuilder:validation:Optional
	AdvancedBackupSetting []AdvancedBackupSettingParameters `json:"advancedBackupSetting,omitempty" tf:"advanced_backup_setting,omitempty"`

	// The display name of a backup plan.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// A rule object that specifies a scheduled task that is used to back up a selection of resources.
	// +kubebuilder:validation:Required
	Rule []RuleParameters `json:"rule" tf:"rule,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type RuleLifecycleObservation struct {
}

type RuleLifecycleParameters struct {

	// Specifies the number of days after creation that a recovery point is moved to cold storage.
	// +kubebuilder:validation:Optional
	ColdStorageAfter *float64 `json:"coldStorageAfter,omitempty" tf:"cold_storage_after,omitempty"`

	// Specifies the number of days after creation that a recovery point is deleted. Must be 90 days greater than cold_storage_after.
	// +kubebuilder:validation:Optional
	DeleteAfter *float64 `json:"deleteAfter,omitempty" tf:"delete_after,omitempty"`
}

type RuleObservation struct {
}

type RuleParameters struct {

	// The amount of time AWS Backup attempts a backup before canceling the job and returning an error.
	// +kubebuilder:validation:Optional
	CompletionWindow *float64 `json:"completionWindow,omitempty" tf:"completion_window,omitempty"`

	// Configuration block(s) with copy operation settings. Detailed below.
	// +kubebuilder:validation:Optional
	CopyAction []CopyActionParameters `json:"copyAction,omitempty" tf:"copy_action,omitempty"`

	// Enable continuous backups for supported resources.
	// +kubebuilder:validation:Optional
	EnableContinuousBackup *bool `json:"enableContinuousBackup,omitempty" tf:"enable_continuous_backup,omitempty"`

	// The lifecycle defines when a protected resource is transitioned to cold storage and when it expires.  Fields documented below.
	// +kubebuilder:validation:Optional
	Lifecycle []RuleLifecycleParameters `json:"lifecycle,omitempty" tf:"lifecycle,omitempty"`

	// Metadata that you can assign to help organize the resources that you create.
	// +kubebuilder:validation:Optional
	RecoveryPointTags map[string]*string `json:"recoveryPointTags,omitempty" tf:"recovery_point_tags,omitempty"`

	// An display name for a backup rule.
	// +kubebuilder:validation:Required
	RuleName *string `json:"ruleName" tf:"rule_name,omitempty"`

	// A CRON expression specifying when AWS Backup initiates a backup job.
	// +kubebuilder:validation:Optional
	Schedule *string `json:"schedule,omitempty" tf:"schedule,omitempty"`

	// The amount of time in minutes before beginning a backup.
	// +kubebuilder:validation:Optional
	StartWindow *float64 `json:"startWindow,omitempty" tf:"start_window,omitempty"`

	// The name of a logical container where backups are stored.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/backup/v1beta1.Vault
	// +kubebuilder:validation:Optional
	TargetVaultName *string `json:"targetVaultName,omitempty" tf:"target_vault_name,omitempty"`

	// Reference to a Vault in backup to populate targetVaultName.
	// +kubebuilder:validation:Optional
	TargetVaultNameRef *v1.Reference `json:"targetVaultNameRef,omitempty" tf:"-"`

	// Selector for a Vault in backup to populate targetVaultName.
	// +kubebuilder:validation:Optional
	TargetVaultNameSelector *v1.Selector `json:"targetVaultNameSelector,omitempty" tf:"-"`
}

// PlanSpec defines the desired state of Plan
type PlanSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PlanParameters `json:"forProvider"`
}

// PlanStatus defines the observed state of Plan.
type PlanStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PlanObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Plan is the Schema for the Plans API. Provides an AWS Backup plan resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Plan struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PlanSpec   `json:"spec"`
	Status            PlanStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PlanList contains a list of Plans
type PlanList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Plan `json:"items"`
}

// Repository type metadata.
var (
	Plan_Kind             = "Plan"
	Plan_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Plan_Kind}.String()
	Plan_KindAPIVersion   = Plan_Kind + "." + CRDGroupVersion.String()
	Plan_GroupVersionKind = CRDGroupVersion.WithKind(Plan_Kind)
)

func init() {
	SchemeBuilder.Register(&Plan{}, &PlanList{})
}
