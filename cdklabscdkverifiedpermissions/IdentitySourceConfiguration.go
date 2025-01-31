package cdklabscdkverifiedpermissions


// Experimental.
type IdentitySourceConfiguration struct {
	// Cognito User Pool Configuration.
	// Default: - no Cognito User Pool Config.
	//
	// Experimental.
	CognitoUserPoolConfiguration *CognitoUserPoolConfiguration `field:"optional" json:"cognitoUserPoolConfiguration" yaml:"cognitoUserPoolConfiguration"`
	// OpenID Connect Idp configuration.
	// Default: - no OpenID Provider config.
	//
	// Experimental.
	OpenIdConnectConfiguration *OpenIdConnectConfiguration `field:"optional" json:"openIdConnectConfiguration" yaml:"openIdConnectConfiguration"`
}

