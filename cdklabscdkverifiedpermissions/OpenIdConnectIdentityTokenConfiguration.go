package cdklabscdkverifiedpermissions


// Experimental.
type OpenIdConnectIdentityTokenConfiguration struct {
	// The ID token audience, or client ID, claim values that you want to accept in your policy store from an OIDC identity provider.
	// Default: - no client IDs.
	//
	// Experimental.
	ClientIds *[]*string `field:"optional" json:"clientIds" yaml:"clientIds"`
	// The claim that determines the principal in OIDC access tokens.
	// Default: - no principal claim.
	//
	// Experimental.
	PrincipalIdClaim *string `field:"optional" json:"principalIdClaim" yaml:"principalIdClaim"`
}

