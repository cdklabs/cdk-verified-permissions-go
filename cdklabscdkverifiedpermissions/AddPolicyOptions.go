package cdklabscdkverifiedpermissions


// Experimental.
type AddPolicyOptions struct {
	// The configuration of the Policy.
	// Experimental.
	PolicyConfiguration *PolicyDefinitionProperty `field:"required" json:"policyConfiguration" yaml:"policyConfiguration"`
	// The id of the Policy.
	// Experimental.
	PolicyId *string `field:"required" json:"policyId" yaml:"policyId"`
}

