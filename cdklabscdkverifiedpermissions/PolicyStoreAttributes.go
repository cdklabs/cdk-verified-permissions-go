package cdklabscdkverifiedpermissions


// Experimental.
type PolicyStoreAttributes struct {
	// The ARN of the Amazon Verified Permissions Policy Store.
	//
	// One of this, or `policyStoreId`, is required.
	// Default: - no PolicyStore arn.
	//
	// Experimental.
	PolicyStoreArn *string `field:"optional" json:"policyStoreArn" yaml:"policyStoreArn"`
	// The id of the Amazon Verified Permissions PolicyStore.
	//
	// One of this, or `policyStoreArn`, is required.
	// Default: - no PolicyStore id.
	//
	// Experimental.
	PolicyStoreId *string `field:"optional" json:"policyStoreId" yaml:"policyStoreId"`
}

