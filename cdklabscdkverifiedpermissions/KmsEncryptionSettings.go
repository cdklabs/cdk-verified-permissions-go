package cdklabscdkverifiedpermissions

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

// Encryption settings using a customer-managed KMS key.
// Experimental.
type KmsEncryptionSettings struct {
	// The KMS key to use for encryption.
	//
	// This can be either a Key construct or an IKey reference.
	// Experimental.
	Key awskms.IKey `field:"required" json:"key" yaml:"key"`
	// Additional encryption context key-value pairs.
	// Default: - No additional encryption context.
	//
	// Experimental.
	EncryptionContext *map[string]*string `field:"optional" json:"encryptionContext" yaml:"encryptionContext"`
}

