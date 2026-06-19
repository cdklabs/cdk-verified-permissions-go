package cdklabscdkverifiedpermissions


// Experimental.
type PolicyStoreAliasProps struct {
	// The name of the policy store alias.
	// Experimental.
	AliasName *string `field:"required" json:"aliasName" yaml:"aliasName"`
	// The policy store to associate with this alias.
	// Experimental.
	PolicyStore IPolicyStore `field:"required" json:"policyStore" yaml:"policyStore"`
}

