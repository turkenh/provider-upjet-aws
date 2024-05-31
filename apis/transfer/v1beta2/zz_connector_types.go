// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1beta2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type As2ConfigInitParameters struct {

	// Specifies weather AS2 file is compressed. The valud values are ZLIB and  DISABLED.
	Compression *string `json:"compression,omitempty" tf:"compression,omitempty"`

	// The algorithm that is used to encrypt the file. The valid values are AES128_CBC | AES192_CBC | AES256_CBC | NONE.
	EncryptionAlgorithm *string `json:"encryptionAlgorithm,omitempty" tf:"encryption_algorithm,omitempty"`

	// The unique identifier for the AS2 local profile.
	LocalProfileID *string `json:"localProfileId,omitempty" tf:"local_profile_id,omitempty"`

	// Used for outbound requests to determine if a partner response for transfers is synchronous or asynchronous. The valid values are SYNC and NONE.
	MdnResponse *string `json:"mdnResponse,omitempty" tf:"mdn_response,omitempty"`

	// The signing algorithm for the Mdn response. The valid values are SHA256 | SHA384 | SHA512 | SHA1 | NONE | DEFAULT.
	MdnSigningAlgorithm *string `json:"mdnSigningAlgorithm,omitempty" tf:"mdn_signing_algorithm,omitempty"`

	// Used as the subject HTTP header attribute in AS2 messages that are being sent with the connector.
	MessageSubject *string `json:"messageSubject,omitempty" tf:"message_subject,omitempty"`

	// The unique identifier for the AS2 partner profile.
	PartnerProfileID *string `json:"partnerProfileId,omitempty" tf:"partner_profile_id,omitempty"`

	// The algorithm that is used to sign AS2 messages sent with the connector. The valid values are SHA256 | SHA384 | SHA512 | SHA1 | NONE .
	SigningAlgorithm *string `json:"signingAlgorithm,omitempty" tf:"signing_algorithm,omitempty"`
}

type As2ConfigObservation struct {

	// Specifies weather AS2 file is compressed. The valud values are ZLIB and  DISABLED.
	Compression *string `json:"compression,omitempty" tf:"compression,omitempty"`

	// The algorithm that is used to encrypt the file. The valid values are AES128_CBC | AES192_CBC | AES256_CBC | NONE.
	EncryptionAlgorithm *string `json:"encryptionAlgorithm,omitempty" tf:"encryption_algorithm,omitempty"`

	// The unique identifier for the AS2 local profile.
	LocalProfileID *string `json:"localProfileId,omitempty" tf:"local_profile_id,omitempty"`

	// Used for outbound requests to determine if a partner response for transfers is synchronous or asynchronous. The valid values are SYNC and NONE.
	MdnResponse *string `json:"mdnResponse,omitempty" tf:"mdn_response,omitempty"`

	// The signing algorithm for the Mdn response. The valid values are SHA256 | SHA384 | SHA512 | SHA1 | NONE | DEFAULT.
	MdnSigningAlgorithm *string `json:"mdnSigningAlgorithm,omitempty" tf:"mdn_signing_algorithm,omitempty"`

	// Used as the subject HTTP header attribute in AS2 messages that are being sent with the connector.
	MessageSubject *string `json:"messageSubject,omitempty" tf:"message_subject,omitempty"`

	// The unique identifier for the AS2 partner profile.
	PartnerProfileID *string `json:"partnerProfileId,omitempty" tf:"partner_profile_id,omitempty"`

	// The algorithm that is used to sign AS2 messages sent with the connector. The valid values are SHA256 | SHA384 | SHA512 | SHA1 | NONE .
	SigningAlgorithm *string `json:"signingAlgorithm,omitempty" tf:"signing_algorithm,omitempty"`
}

type As2ConfigParameters struct {

	// Specifies weather AS2 file is compressed. The valud values are ZLIB and  DISABLED.
	// +kubebuilder:validation:Optional
	Compression *string `json:"compression" tf:"compression,omitempty"`

	// The algorithm that is used to encrypt the file. The valid values are AES128_CBC | AES192_CBC | AES256_CBC | NONE.
	// +kubebuilder:validation:Optional
	EncryptionAlgorithm *string `json:"encryptionAlgorithm" tf:"encryption_algorithm,omitempty"`

	// The unique identifier for the AS2 local profile.
	// +kubebuilder:validation:Optional
	LocalProfileID *string `json:"localProfileId" tf:"local_profile_id,omitempty"`

	// Used for outbound requests to determine if a partner response for transfers is synchronous or asynchronous. The valid values are SYNC and NONE.
	// +kubebuilder:validation:Optional
	MdnResponse *string `json:"mdnResponse" tf:"mdn_response,omitempty"`

	// The signing algorithm for the Mdn response. The valid values are SHA256 | SHA384 | SHA512 | SHA1 | NONE | DEFAULT.
	// +kubebuilder:validation:Optional
	MdnSigningAlgorithm *string `json:"mdnSigningAlgorithm,omitempty" tf:"mdn_signing_algorithm,omitempty"`

	// Used as the subject HTTP header attribute in AS2 messages that are being sent with the connector.
	// +kubebuilder:validation:Optional
	MessageSubject *string `json:"messageSubject,omitempty" tf:"message_subject,omitempty"`

	// The unique identifier for the AS2 partner profile.
	// +kubebuilder:validation:Optional
	PartnerProfileID *string `json:"partnerProfileId" tf:"partner_profile_id,omitempty"`

	// The algorithm that is used to sign AS2 messages sent with the connector. The valid values are SHA256 | SHA384 | SHA512 | SHA1 | NONE .
	// +kubebuilder:validation:Optional
	SigningAlgorithm *string `json:"signingAlgorithm" tf:"signing_algorithm,omitempty"`
}

type ConnectorInitParameters struct {

	// The IAM Role which provides read and write access to the parent directory of the file location mentioned in the StartFileTransfer request.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("arn",true)
	AccessRole *string `json:"accessRole,omitempty" tf:"access_role,omitempty"`

	// Reference to a Role in iam to populate accessRole.
	// +kubebuilder:validation:Optional
	AccessRoleRef *v1.Reference `json:"accessRoleRef,omitempty" tf:"-"`

	// Selector for a Role in iam to populate accessRole.
	// +kubebuilder:validation:Optional
	AccessRoleSelector *v1.Selector `json:"accessRoleSelector,omitempty" tf:"-"`

	// Either SFTP or AS2 is configured.The parameters to configure for the connector object. Fields documented below.
	As2Config *As2ConfigInitParameters `json:"as2Config,omitempty" tf:"as2_config,omitempty"`

	// The IAM Role which is required for allowing the connector to turn on CloudWatch logging for Amazon S3 events.
	LoggingRole *string `json:"loggingRole,omitempty" tf:"logging_role,omitempty"`

	// Name of the security policy for the connector.
	SecurityPolicyName *string `json:"securityPolicyName,omitempty" tf:"security_policy_name,omitempty"`

	// Either SFTP or AS2 is configured.The parameters to configure for the connector object. Fields documented below.
	SftpConfig *SftpConfigInitParameters `json:"sftpConfig,omitempty" tf:"sftp_config,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The URL of the partners AS2 endpoint or SFTP endpoint.
	URL *string `json:"url,omitempty" tf:"url,omitempty"`
}

type ConnectorObservation struct {

	// The IAM Role which provides read and write access to the parent directory of the file location mentioned in the StartFileTransfer request.
	AccessRole *string `json:"accessRole,omitempty" tf:"access_role,omitempty"`

	// The ARN of the connector.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Either SFTP or AS2 is configured.The parameters to configure for the connector object. Fields documented below.
	As2Config *As2ConfigObservation `json:"as2Config,omitempty" tf:"as2_config,omitempty"`

	// The unique identifier for the AS2 profile or SFTP Profile.
	ConnectorID *string `json:"connectorId,omitempty" tf:"connector_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The IAM Role which is required for allowing the connector to turn on CloudWatch logging for Amazon S3 events.
	LoggingRole *string `json:"loggingRole,omitempty" tf:"logging_role,omitempty"`

	// Name of the security policy for the connector.
	SecurityPolicyName *string `json:"securityPolicyName,omitempty" tf:"security_policy_name,omitempty"`

	// Either SFTP or AS2 is configured.The parameters to configure for the connector object. Fields documented below.
	SftpConfig *SftpConfigObservation `json:"sftpConfig,omitempty" tf:"sftp_config,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +mapType=granular
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// The URL of the partners AS2 endpoint or SFTP endpoint.
	URL *string `json:"url,omitempty" tf:"url,omitempty"`
}

type ConnectorParameters struct {

	// The IAM Role which provides read and write access to the parent directory of the file location mentioned in the StartFileTransfer request.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("arn",true)
	// +kubebuilder:validation:Optional
	AccessRole *string `json:"accessRole,omitempty" tf:"access_role,omitempty"`

	// Reference to a Role in iam to populate accessRole.
	// +kubebuilder:validation:Optional
	AccessRoleRef *v1.Reference `json:"accessRoleRef,omitempty" tf:"-"`

	// Selector for a Role in iam to populate accessRole.
	// +kubebuilder:validation:Optional
	AccessRoleSelector *v1.Selector `json:"accessRoleSelector,omitempty" tf:"-"`

	// Either SFTP or AS2 is configured.The parameters to configure for the connector object. Fields documented below.
	// +kubebuilder:validation:Optional
	As2Config *As2ConfigParameters `json:"as2Config,omitempty" tf:"as2_config,omitempty"`

	// The IAM Role which is required for allowing the connector to turn on CloudWatch logging for Amazon S3 events.
	// +kubebuilder:validation:Optional
	LoggingRole *string `json:"loggingRole,omitempty" tf:"logging_role,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Name of the security policy for the connector.
	// +kubebuilder:validation:Optional
	SecurityPolicyName *string `json:"securityPolicyName,omitempty" tf:"security_policy_name,omitempty"`

	// Either SFTP or AS2 is configured.The parameters to configure for the connector object. Fields documented below.
	// +kubebuilder:validation:Optional
	SftpConfig *SftpConfigParameters `json:"sftpConfig,omitempty" tf:"sftp_config,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The URL of the partners AS2 endpoint or SFTP endpoint.
	// +kubebuilder:validation:Optional
	URL *string `json:"url,omitempty" tf:"url,omitempty"`
}

type SftpConfigInitParameters struct {

	// A list of public portion of the host key, or keys, that are used to authenticate the user to the external server to which you are connecting.(https://docs.aws.amazon.com/transfer/latest/userguide/API_SftpConnectorConfig.html)
	// +listType=set
	TrustedHostKeys []*string `json:"trustedHostKeys,omitempty" tf:"trusted_host_keys,omitempty"`

	// The identifier for the secret (in AWS Secrets Manager) that contains the SFTP user's private key, password, or both. The identifier can be either the Amazon Resource Name (ARN) or the name of the secret.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/secretsmanager/v1beta1.Secret
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	UserSecretID *string `json:"userSecretId,omitempty" tf:"user_secret_id,omitempty"`

	// Reference to a Secret in secretsmanager to populate userSecretId.
	// +kubebuilder:validation:Optional
	UserSecretIDRef *v1.Reference `json:"userSecretIdRef,omitempty" tf:"-"`

	// Selector for a Secret in secretsmanager to populate userSecretId.
	// +kubebuilder:validation:Optional
	UserSecretIDSelector *v1.Selector `json:"userSecretIdSelector,omitempty" tf:"-"`
}

type SftpConfigObservation struct {

	// A list of public portion of the host key, or keys, that are used to authenticate the user to the external server to which you are connecting.(https://docs.aws.amazon.com/transfer/latest/userguide/API_SftpConnectorConfig.html)
	// +listType=set
	TrustedHostKeys []*string `json:"trustedHostKeys,omitempty" tf:"trusted_host_keys,omitempty"`

	// The identifier for the secret (in AWS Secrets Manager) that contains the SFTP user's private key, password, or both. The identifier can be either the Amazon Resource Name (ARN) or the name of the secret.
	UserSecretID *string `json:"userSecretId,omitempty" tf:"user_secret_id,omitempty"`
}

type SftpConfigParameters struct {

	// A list of public portion of the host key, or keys, that are used to authenticate the user to the external server to which you are connecting.(https://docs.aws.amazon.com/transfer/latest/userguide/API_SftpConnectorConfig.html)
	// +kubebuilder:validation:Optional
	// +listType=set
	TrustedHostKeys []*string `json:"trustedHostKeys,omitempty" tf:"trusted_host_keys,omitempty"`

	// The identifier for the secret (in AWS Secrets Manager) that contains the SFTP user's private key, password, or both. The identifier can be either the Amazon Resource Name (ARN) or the name of the secret.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/secretsmanager/v1beta1.Secret
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	UserSecretID *string `json:"userSecretId,omitempty" tf:"user_secret_id,omitempty"`

	// Reference to a Secret in secretsmanager to populate userSecretId.
	// +kubebuilder:validation:Optional
	UserSecretIDRef *v1.Reference `json:"userSecretIdRef,omitempty" tf:"-"`

	// Selector for a Secret in secretsmanager to populate userSecretId.
	// +kubebuilder:validation:Optional
	UserSecretIDSelector *v1.Selector `json:"userSecretIdSelector,omitempty" tf:"-"`
}

// ConnectorSpec defines the desired state of Connector
type ConnectorSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ConnectorParameters `json:"forProvider"`
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
	InitProvider ConnectorInitParameters `json:"initProvider,omitempty"`
}

// ConnectorStatus defines the observed state of Connector.
type ConnectorStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ConnectorObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// Connector is the Schema for the Connectors API. Provides a AWS Transfer AS2 Connector Resource
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Connector struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.url) || (has(self.initProvider) && has(self.initProvider.url))",message="spec.forProvider.url is a required parameter"
	Spec   ConnectorSpec   `json:"spec"`
	Status ConnectorStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ConnectorList contains a list of Connectors
type ConnectorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Connector `json:"items"`
}

// Repository type metadata.
var (
	Connector_Kind             = "Connector"
	Connector_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Connector_Kind}.String()
	Connector_KindAPIVersion   = Connector_Kind + "." + CRDGroupVersion.String()
	Connector_GroupVersionKind = CRDGroupVersion.WithKind(Connector_Kind)
)

func init() {
	SchemeBuilder.Register(&Connector{}, &ConnectorList{})
}
