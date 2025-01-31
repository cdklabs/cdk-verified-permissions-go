package cdklabscdkverifiedpermissions


// Experimental.
type PolicyStoreProps struct {
	// The policy store's validation settings.
	// Default: - If not provided, the Policy store will be created with ValidationSettingsMode = "OFF".
	//
	// Experimental.
	ValidationSettings *ValidationSettings `field:"required" json:"validationSettings" yaml:"validationSettings"`
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
}

