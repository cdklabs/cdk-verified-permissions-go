package cdklabscdkverifiedpermissions


// Experimental.
type TemplateLinkedPolicyDefinitionProperty struct {
	// The unique identifier of the policy template used to create this policy.
	// Experimental.
	PolicyTemplate IPolicyTemplate `field:"required" json:"policyTemplate" yaml:"policyTemplate"`
	// The principal associated with this template-linked policy.
	// Default: - No Principal. It is set to unspecified.
	//
	// Experimental.
	Principal *EntityIdentifierProperty `field:"optional" json:"principal" yaml:"principal"`
	// The resource associated with this template-linked policy.
	// Default: - No Resource. It is set to unspecified.
	//
	// Experimental.
	Resource *EntityIdentifierProperty `field:"optional" json:"resource" yaml:"resource"`
}

