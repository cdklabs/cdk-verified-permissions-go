package cdklabscdkverifiedpermissions


// Encryption settings for the policy store.
//
// This is a union type - provide either
// `awsOwnedKey` set to true for AWS owned encryption, or `customerManagedKey` for
// customer-managed KMS key encryption. These options are mutually exclusive.
// Experimental.
type EncryptionSettings struct {
	// Use an AWS owned key for encryption.
	//
	// Cannot be specified together with `customerManagedKey`.
	// Default: - false.
	//
	// Experimental.
	AwsOwnedKey *bool `field:"optional" json:"awsOwnedKey" yaml:"awsOwnedKey"`
	// Use a customer-managed KMS key for encryption.
	//
	// Cannot be specified together with `awsOwnedKey`.
	// Default: - undefined.
	//
	// Experimental.
	CustomerManagedKey *KmsEncryptionSettings `field:"optional" json:"customerManagedKey" yaml:"customerManagedKey"`
}

