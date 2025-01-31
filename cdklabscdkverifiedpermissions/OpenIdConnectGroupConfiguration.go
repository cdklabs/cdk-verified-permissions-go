package cdklabscdkverifiedpermissions


// Experimental.
type OpenIdConnectGroupConfiguration struct {
	// The token claim that you want Verified Permissions to interpret as group membership.
	// Experimental.
	GroupClaim *string `field:"required" json:"groupClaim" yaml:"groupClaim"`
	// The policy store entity type that you want to map your users' group claim to.
	// Experimental.
	GroupEntityType *string `field:"required" json:"groupEntityType" yaml:"groupEntityType"`
}

