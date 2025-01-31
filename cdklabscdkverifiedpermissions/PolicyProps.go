package cdklabscdkverifiedpermissions


// Experimental.
type PolicyProps struct {
	// Specifies the policy type and content to use for the new or updated policy.
	//
	// The definition structure must include either a Static or a TemplateLinked element.
	// Experimental.
	Definition *PolicyDefinitionProperty `field:"required" json:"definition" yaml:"definition"`
	// The policy store that contains the policy.
	// Experimental.
	PolicyStore IPolicyStore `field:"required" json:"policyStore" yaml:"policyStore"`
}

