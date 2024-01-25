// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

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

type IdentityInfoInitParameters struct {

	// The email address. If you are using SAML for identity management and include this parameter, an error is returned. Note that updates to the email is supported. From the UpdateUserIdentityInfo API documentation it is strongly recommended to limit who has the ability to invoke UpdateUserIdentityInfo. Someone with that ability can change the login credentials of other users by changing their email address. This poses a security risk to your organization. They can change the email address of a user to the attacker's email address, and then reset the password through email. For more information, see Best Practices for Security Profiles in the Amazon Connect Administrator Guide.
	Email *string `json:"email,omitempty" tf:"email,omitempty"`

	// The first name. This is required if you are using Amazon Connect or SAML for identity management. Minimum length of 1. Maximum length of 100.
	FirstName *string `json:"firstName,omitempty" tf:"first_name,omitempty"`

	// The last name. This is required if you are using Amazon Connect or SAML for identity management. Minimum length of 1. Maximum length of 100.
	LastName *string `json:"lastName,omitempty" tf:"last_name,omitempty"`
}

type IdentityInfoObservation struct {

	// The email address. If you are using SAML for identity management and include this parameter, an error is returned. Note that updates to the email is supported. From the UpdateUserIdentityInfo API documentation it is strongly recommended to limit who has the ability to invoke UpdateUserIdentityInfo. Someone with that ability can change the login credentials of other users by changing their email address. This poses a security risk to your organization. They can change the email address of a user to the attacker's email address, and then reset the password through email. For more information, see Best Practices for Security Profiles in the Amazon Connect Administrator Guide.
	Email *string `json:"email,omitempty" tf:"email,omitempty"`

	// The first name. This is required if you are using Amazon Connect or SAML for identity management. Minimum length of 1. Maximum length of 100.
	FirstName *string `json:"firstName,omitempty" tf:"first_name,omitempty"`

	// The last name. This is required if you are using Amazon Connect or SAML for identity management. Minimum length of 1. Maximum length of 100.
	LastName *string `json:"lastName,omitempty" tf:"last_name,omitempty"`
}

type IdentityInfoParameters struct {

	// The email address. If you are using SAML for identity management and include this parameter, an error is returned. Note that updates to the email is supported. From the UpdateUserIdentityInfo API documentation it is strongly recommended to limit who has the ability to invoke UpdateUserIdentityInfo. Someone with that ability can change the login credentials of other users by changing their email address. This poses a security risk to your organization. They can change the email address of a user to the attacker's email address, and then reset the password through email. For more information, see Best Practices for Security Profiles in the Amazon Connect Administrator Guide.
	// +kubebuilder:validation:Optional
	Email *string `json:"email,omitempty" tf:"email,omitempty"`

	// The first name. This is required if you are using Amazon Connect or SAML for identity management. Minimum length of 1. Maximum length of 100.
	// +kubebuilder:validation:Optional
	FirstName *string `json:"firstName,omitempty" tf:"first_name,omitempty"`

	// The last name. This is required if you are using Amazon Connect or SAML for identity management. Minimum length of 1. Maximum length of 100.
	// +kubebuilder:validation:Optional
	LastName *string `json:"lastName,omitempty" tf:"last_name,omitempty"`
}

type UserInitParameters struct {

	// The identifier of the user account in the directory used for identity management. If Amazon Connect cannot access the directory, you can specify this identifier to authenticate users. If you include the identifier, we assume that Amazon Connect cannot access the directory. Otherwise, the identity information is used to authenticate users from your directory. This parameter is required if you are using an existing directory for identity management in Amazon Connect when Amazon Connect cannot access your directory to authenticate users. If you are using SAML for identity management and include this parameter, an error is returned.
	DirectoryUserID *string `json:"directoryUserId,omitempty" tf:"directory_user_id,omitempty"`

	// The identifier of the hierarchy group for the user.
	HierarchyGroupID *string `json:"hierarchyGroupId,omitempty" tf:"hierarchy_group_id,omitempty"`

	// A block that contains information about the identity of the user. Documented below.
	IdentityInfo []IdentityInfoInitParameters `json:"identityInfo,omitempty" tf:"identity_info,omitempty"`

	// Specifies the identifier of the hosting Amazon Connect Instance.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/connect/v1beta1.Instance
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// Reference to a Instance in connect to populate instanceId.
	// +kubebuilder:validation:Optional
	InstanceIDRef *v1.Reference `json:"instanceIdRef,omitempty" tf:"-"`

	// Selector for a Instance in connect to populate instanceId.
	// +kubebuilder:validation:Optional
	InstanceIDSelector *v1.Selector `json:"instanceIdSelector,omitempty" tf:"-"`

	// The user name for the account. For instances not using SAML for identity management, the user name can include up to 20 characters. If you are using SAML for identity management, the user name can include up to 64 characters from [a-zA-Z0-9_-.\@]+.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// A block that contains information about the phone settings for the user. Documented below.
	PhoneConfig []UserPhoneConfigInitParameters `json:"phoneConfig,omitempty" tf:"phone_config,omitempty"`

	// The identifier of the routing profile for the user.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/connect/v1beta2.RoutingProfile
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("routing_profile_id",true)
	RoutingProfileID *string `json:"routingProfileId,omitempty" tf:"routing_profile_id,omitempty"`

	// Reference to a RoutingProfile in connect to populate routingProfileId.
	// +kubebuilder:validation:Optional
	RoutingProfileIDRef *v1.Reference `json:"routingProfileIdRef,omitempty" tf:"-"`

	// Selector for a RoutingProfile in connect to populate routingProfileId.
	// +kubebuilder:validation:Optional
	RoutingProfileIDSelector *v1.Selector `json:"routingProfileIdSelector,omitempty" tf:"-"`

	// A list of identifiers for the security profiles for the user. Specify a minimum of 1 and maximum of 10 security profile ids. For more information, see Best Practices for Security Profiles in the Amazon Connect Administrator Guide.
	// +listType=set
	SecurityProfileIds []*string `json:"securityProfileIds,omitempty" tf:"security_profile_ids,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type UserObservation struct {

	// The Amazon Resource Name (ARN) of the user.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The identifier of the user account in the directory used for identity management. If Amazon Connect cannot access the directory, you can specify this identifier to authenticate users. If you include the identifier, we assume that Amazon Connect cannot access the directory. Otherwise, the identity information is used to authenticate users from your directory. This parameter is required if you are using an existing directory for identity management in Amazon Connect when Amazon Connect cannot access your directory to authenticate users. If you are using SAML for identity management and include this parameter, an error is returned.
	DirectoryUserID *string `json:"directoryUserId,omitempty" tf:"directory_user_id,omitempty"`

	// The identifier of the hierarchy group for the user.
	HierarchyGroupID *string `json:"hierarchyGroupId,omitempty" tf:"hierarchy_group_id,omitempty"`

	// The identifier of the hosting Amazon Connect Instance and identifier of the user
	// separated by a colon (:).
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A block that contains information about the identity of the user. Documented below.
	IdentityInfo []IdentityInfoObservation `json:"identityInfo,omitempty" tf:"identity_info,omitempty"`

	// Specifies the identifier of the hosting Amazon Connect Instance.
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// The user name for the account. For instances not using SAML for identity management, the user name can include up to 20 characters. If you are using SAML for identity management, the user name can include up to 64 characters from [a-zA-Z0-9_-.\@]+.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// A block that contains information about the phone settings for the user. Documented below.
	PhoneConfig []UserPhoneConfigObservation `json:"phoneConfig,omitempty" tf:"phone_config,omitempty"`

	// The identifier of the routing profile for the user.
	RoutingProfileID *string `json:"routingProfileId,omitempty" tf:"routing_profile_id,omitempty"`

	// A list of identifiers for the security profiles for the user. Specify a minimum of 1 and maximum of 10 security profile ids. For more information, see Best Practices for Security Profiles in the Amazon Connect Administrator Guide.
	// +listType=set
	SecurityProfileIds []*string `json:"securityProfileIds,omitempty" tf:"security_profile_ids,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	// +mapType=granular
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// The identifier for the user.
	UserID *string `json:"userId,omitempty" tf:"user_id,omitempty"`
}

type UserParameters struct {

	// The identifier of the user account in the directory used for identity management. If Amazon Connect cannot access the directory, you can specify this identifier to authenticate users. If you include the identifier, we assume that Amazon Connect cannot access the directory. Otherwise, the identity information is used to authenticate users from your directory. This parameter is required if you are using an existing directory for identity management in Amazon Connect when Amazon Connect cannot access your directory to authenticate users. If you are using SAML for identity management and include this parameter, an error is returned.
	// +kubebuilder:validation:Optional
	DirectoryUserID *string `json:"directoryUserId,omitempty" tf:"directory_user_id,omitempty"`

	// The identifier of the hierarchy group for the user.
	// +kubebuilder:validation:Optional
	HierarchyGroupID *string `json:"hierarchyGroupId,omitempty" tf:"hierarchy_group_id,omitempty"`

	// A block that contains information about the identity of the user. Documented below.
	// +kubebuilder:validation:Optional
	IdentityInfo []IdentityInfoParameters `json:"identityInfo,omitempty" tf:"identity_info,omitempty"`

	// Specifies the identifier of the hosting Amazon Connect Instance.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/connect/v1beta1.Instance
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// Reference to a Instance in connect to populate instanceId.
	// +kubebuilder:validation:Optional
	InstanceIDRef *v1.Reference `json:"instanceIdRef,omitempty" tf:"-"`

	// Selector for a Instance in connect to populate instanceId.
	// +kubebuilder:validation:Optional
	InstanceIDSelector *v1.Selector `json:"instanceIdSelector,omitempty" tf:"-"`

	// The user name for the account. For instances not using SAML for identity management, the user name can include up to 20 characters. If you are using SAML for identity management, the user name can include up to 64 characters from [a-zA-Z0-9_-.\@]+.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The password for the user account. A password is required if you are using Amazon Connect for identity management. Otherwise, it is an error to include a password.
	// +kubebuilder:validation:Optional
	PasswordSecretRef *v1.SecretKeySelector `json:"passwordSecretRef,omitempty" tf:"-"`

	// A block that contains information about the phone settings for the user. Documented below.
	// +kubebuilder:validation:Optional
	PhoneConfig []UserPhoneConfigParameters `json:"phoneConfig,omitempty" tf:"phone_config,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The identifier of the routing profile for the user.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/connect/v1beta2.RoutingProfile
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("routing_profile_id",true)
	// +kubebuilder:validation:Optional
	RoutingProfileID *string `json:"routingProfileId,omitempty" tf:"routing_profile_id,omitempty"`

	// Reference to a RoutingProfile in connect to populate routingProfileId.
	// +kubebuilder:validation:Optional
	RoutingProfileIDRef *v1.Reference `json:"routingProfileIdRef,omitempty" tf:"-"`

	// Selector for a RoutingProfile in connect to populate routingProfileId.
	// +kubebuilder:validation:Optional
	RoutingProfileIDSelector *v1.Selector `json:"routingProfileIdSelector,omitempty" tf:"-"`

	// A list of identifiers for the security profiles for the user. Specify a minimum of 1 and maximum of 10 security profile ids. For more information, see Best Practices for Security Profiles in the Amazon Connect Administrator Guide.
	// +kubebuilder:validation:Optional
	// +listType=set
	SecurityProfileIds []*string `json:"securityProfileIds,omitempty" tf:"security_profile_ids,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type UserPhoneConfigInitParameters struct {

	// The After Call Work (ACW) timeout setting, in seconds. Minimum value of 0.
	AfterContactWorkTimeLimit *float64 `json:"afterContactWorkTimeLimit,omitempty" tf:"after_contact_work_time_limit,omitempty"`

	// When Auto-Accept Call is enabled for an available agent, the agent connects to contacts automatically.
	AutoAccept *bool `json:"autoAccept,omitempty" tf:"auto_accept,omitempty"`

	// The phone number for the user's desk phone. Required if phone_type is set as DESK_PHONE.
	DeskPhoneNumber *string `json:"deskPhoneNumber,omitempty" tf:"desk_phone_number,omitempty"`

	// The phone type. Valid values are DESK_PHONE and SOFT_PHONE.
	PhoneType *string `json:"phoneType,omitempty" tf:"phone_type,omitempty"`
}

type UserPhoneConfigObservation struct {

	// The After Call Work (ACW) timeout setting, in seconds. Minimum value of 0.
	AfterContactWorkTimeLimit *float64 `json:"afterContactWorkTimeLimit,omitempty" tf:"after_contact_work_time_limit,omitempty"`

	// When Auto-Accept Call is enabled for an available agent, the agent connects to contacts automatically.
	AutoAccept *bool `json:"autoAccept,omitempty" tf:"auto_accept,omitempty"`

	// The phone number for the user's desk phone. Required if phone_type is set as DESK_PHONE.
	DeskPhoneNumber *string `json:"deskPhoneNumber,omitempty" tf:"desk_phone_number,omitempty"`

	// The phone type. Valid values are DESK_PHONE and SOFT_PHONE.
	PhoneType *string `json:"phoneType,omitempty" tf:"phone_type,omitempty"`
}

type UserPhoneConfigParameters struct {

	// The After Call Work (ACW) timeout setting, in seconds. Minimum value of 0.
	// +kubebuilder:validation:Optional
	AfterContactWorkTimeLimit *float64 `json:"afterContactWorkTimeLimit,omitempty" tf:"after_contact_work_time_limit,omitempty"`

	// When Auto-Accept Call is enabled for an available agent, the agent connects to contacts automatically.
	// +kubebuilder:validation:Optional
	AutoAccept *bool `json:"autoAccept,omitempty" tf:"auto_accept,omitempty"`

	// The phone number for the user's desk phone. Required if phone_type is set as DESK_PHONE.
	// +kubebuilder:validation:Optional
	DeskPhoneNumber *string `json:"deskPhoneNumber,omitempty" tf:"desk_phone_number,omitempty"`

	// The phone type. Valid values are DESK_PHONE and SOFT_PHONE.
	// +kubebuilder:validation:Optional
	PhoneType *string `json:"phoneType" tf:"phone_type,omitempty"`
}

// UserSpec defines the desired state of User
type UserSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     UserParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider UserInitParameters `json:"initProvider,omitempty"`
}

// UserStatus defines the observed state of User.
type UserStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        UserObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// User is the Schema for the Users API. Provides details about a specific Amazon Connect User
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type User struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.phoneConfig) || (has(self.initProvider) && has(self.initProvider.phoneConfig))",message="spec.forProvider.phoneConfig is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.securityProfileIds) || (has(self.initProvider) && has(self.initProvider.securityProfileIds))",message="spec.forProvider.securityProfileIds is a required parameter"
	Spec   UserSpec   `json:"spec"`
	Status UserStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// UserList contains a list of Users
type UserList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []User `json:"items"`
}

// Repository type metadata.
var (
	User_Kind             = "User"
	User_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: User_Kind}.String()
	User_KindAPIVersion   = User_Kind + "." + CRDGroupVersion.String()
	User_GroupVersionKind = CRDGroupVersion.WithKind(User_Kind)
)

func init() {
	SchemeBuilder.Register(&User{}, &UserList{})
}
