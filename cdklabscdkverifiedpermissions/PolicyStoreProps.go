package cdklabscdkverifiedpermissions


// Experimental.
type PolicyStoreProps struct {
	// The policy store's validation settings.
	// Experimental.
	ValidationSettings *ValidationSettings `field:"required" json:"validationSettings" yaml:"validationSettings"`
	// The policy store's deletion protection.
	// Default: - If not provided, the Policy store will be created with deletionProtection = "DISABLED".
	//
	// Experimental.
	DeletionProtection DeletionProtectionMode `field:"optional" json:"deletionProtection" yaml:"deletionProtection"`
	// The policy store's description.
	// Default: - No description.
	//
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// This attribute is not required from an API point of view.
	//
	// It represents the schema (in Cedar) to be applied to the PolicyStore.
	// Default: - No schema.
	//
	// Experimental.
	Schema *Schema `field:"optional" json:"schema" yaml:"schema"`
	// The tags assigned to the policy store.
	// Default: - none.
	//
	// Experimental.
	Tags *[]*Tag `field:"optional" json:"tags" yaml:"tags"`
}

