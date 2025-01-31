package cdklabscdkverifiedpermissions


// Experimental.
type PolicyDefinitionProperty struct {
	// A structure that describes a static policy.
	// Default: - Static must be set for policies created from a static definition. Otherwise, use template linked definitions.
	//
	// Experimental.
	Static *StaticPolicyDefinitionProperty `field:"optional" json:"static" yaml:"static"`
	// A structure that describes a policy that was instantiated from a template.
	// Default: - Template linked must be set for policies created from a static definition. Otherwise, use static definitions.
	//
	// Experimental.
	TemplateLinked *TemplateLinkedPolicyDefinitionProperty `field:"optional" json:"templateLinked" yaml:"templateLinked"`
}

