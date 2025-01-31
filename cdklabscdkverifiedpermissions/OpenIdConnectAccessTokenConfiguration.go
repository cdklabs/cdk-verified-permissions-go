package cdklabscdkverifiedpermissions


// Experimental.
type OpenIdConnectAccessTokenConfiguration struct {
	// The access token aud claim values that you want to accept in your policy store.
	// Default: - no audiences.
	//
	// Experimental.
	Audiences *[]*string `field:"optional" json:"audiences" yaml:"audiences"`
	// The claim that determines the principal in OIDC access tokens.
	// Default: - no principal claim.
	//
	// Experimental.
	PrincipalIdClaim *string `field:"optional" json:"principalIdClaim" yaml:"principalIdClaim"`
}

