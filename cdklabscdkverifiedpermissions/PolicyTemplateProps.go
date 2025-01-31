package cdklabscdkverifiedpermissions


// Experimental.
type PolicyTemplateProps struct {
	// The policy store that contains the template.
	// Default: - The policy store to attach the new or updated policy template.
	//
	// Experimental.
	PolicyStore IPolicyStore `field:"required" json:"policyStore" yaml:"policyStore"`
	// Specifies the content that you want to use for the new policy template, written in the Cedar policy language.
	// Default: - The statement to attach to the new or updated policy template.
	//
	// Experimental.
	Statement *string `field:"required" json:"statement" yaml:"statement"`
	// The description to attach to the new or updated policy template.
	// Default: - No description.
	//
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

