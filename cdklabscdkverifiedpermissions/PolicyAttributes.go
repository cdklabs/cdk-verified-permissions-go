package cdklabscdkverifiedpermissions


// Experimental.
type PolicyAttributes struct {
	// The unique ID of the new or updated policy.
	// Experimental.
	PolicyId *string `field:"required" json:"policyId" yaml:"policyId"`
	// The type of the policy.
	//
	// This is one of the following values: Static or TemplateLinked.
	// Default: - Static.
	//
	// Experimental.
	PolicyType PolicyType `field:"optional" json:"policyType" yaml:"policyType"`
}

