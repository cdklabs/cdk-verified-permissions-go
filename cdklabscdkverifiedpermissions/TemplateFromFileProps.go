package cdklabscdkverifiedpermissions


// Experimental.
type TemplateFromFileProps struct {
	// The path to the file to be read which contains a single cedar statement representing a policy template.
	// Experimental.
	Path *string `field:"required" json:"path" yaml:"path"`
	// The policy store that the policy template will be created under.
	// Experimental.
	PolicyStore IPolicyStore `field:"required" json:"policyStore" yaml:"policyStore"`
	// The description of the plicy template.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

