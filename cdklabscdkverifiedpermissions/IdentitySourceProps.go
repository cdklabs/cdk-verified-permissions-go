package cdklabscdkverifiedpermissions


// Experimental.
type IdentitySourceProps struct {
	// Identity Source configuration.
	// Experimental.
	Configuration *IdentitySourceConfiguration `field:"required" json:"configuration" yaml:"configuration"`
	// Policy Store in which you want to store this identity source.
	// Experimental.
	PolicyStore IPolicyStore `field:"required" json:"policyStore" yaml:"policyStore"`
	// Principal entity type.
	// Default: - No principal entity type for the identity source.
	//
	// Experimental.
	PrincipalEntityType *string `field:"optional" json:"principalEntityType" yaml:"principalEntityType"`
}

