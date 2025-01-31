package cdklabscdkverifiedpermissions


// Experimental.
type StaticPolicyDefinitionProperty struct {
	// The policy content of the static policy, written in the Cedar policy language.
	//
	// You can specify a description of the policy
	// directly inside the policy statement, using the Cedar annotation '@cdkDescription'.
	// Experimental.
	Statement *string `field:"required" json:"statement" yaml:"statement"`
	// The description of the static policy.
	//
	// If this is set, it has always precedence over description defined in policy statement
	// through '@cdkDescription' annotation.
	// Default: - Empty description.
	//
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

