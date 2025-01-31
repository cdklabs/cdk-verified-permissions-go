package cdklabscdkverifiedpermissions


// Experimental.
type OpenIdConnectConfiguration struct {
	// The issuer URL of an OIDC identity provider.
	//
	// This URL must have an OIDC discovery endpoint at the path .well-known/openid-configuration
	// Experimental.
	Issuer *string `field:"required" json:"issuer" yaml:"issuer"`
	// The configuration for processing access tokens from your OIDC identity provider Exactly one between accessTokenOnly and identityTokenOnly must be defined.
	// Default: - no Access Token Config.
	//
	// Experimental.
	AccessTokenOnly *OpenIdConnectAccessTokenConfiguration `field:"optional" json:"accessTokenOnly" yaml:"accessTokenOnly"`
	// A descriptive string that you want to prefix to user entities from your OIDC identity provider.
	// Default: - no Entity ID Prefix.
	//
	// Experimental.
	EntityIdPrefix *string `field:"optional" json:"entityIdPrefix" yaml:"entityIdPrefix"`
	// The claim in OIDC identity provider tokens that indicates a user's group membership, and the entity type that you want to map it to.
	// Default: - no Group Config.
	//
	// Experimental.
	GroupConfiguration *OpenIdConnectGroupConfiguration `field:"optional" json:"groupConfiguration" yaml:"groupConfiguration"`
	// The configuration for processing identity (ID) tokens from your OIDC identity provider Exactly one between accessTokenOnly and identityTokenOnly must be defined.
	// Default: - no ID Token Config.
	//
	// Experimental.
	IdentityTokenOnly *OpenIdConnectIdentityTokenConfiguration `field:"optional" json:"identityTokenOnly" yaml:"identityTokenOnly"`
}

