package cdklabscdkverifiedpermissions


// Experimental.
type StaticPolicyFromFileProps struct {
	// The path to the file to be read which contains a single cedar statement representing a policy.
	// Experimental.
	Path *string `field:"required" json:"path" yaml:"path"`
	// The policy store that the policy will be created under.
	// Experimental.
	PolicyStore IPolicyStore `field:"required" json:"policyStore" yaml:"policyStore"`
	// The default description of static policies, this will be applied to every policy if the description is not retrieved via the.
	// See: getPolicyDescription method in cedar-helpers.
	//
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Boolean flag to activate policy validation against Cedar Language Syntax & Rules.
	// Default: - true.
	//
	// Experimental.
	EnablePolicyValidation *bool `field:"optional" json:"enablePolicyValidation" yaml:"enablePolicyValidation"`
}

