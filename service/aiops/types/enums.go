// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type EncryptionConfigurationType string

// Enum values for EncryptionConfigurationType
const (
	EncryptionConfigurationTypeAwsOwnedKey           EncryptionConfigurationType = "AWS_OWNED_KEY"
	EncryptionConfigurationTypeCustomerManagedKmsKey EncryptionConfigurationType = "CUSTOMER_MANAGED_KMS_KEY"
)

// Values returns all known values for EncryptionConfigurationType. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (EncryptionConfigurationType) Values() []EncryptionConfigurationType {
	return []EncryptionConfigurationType{
		"AWS_OWNED_KEY",
		"CUSTOMER_MANAGED_KMS_KEY",
	}
}
