package cdklabscdkverifiedpermissions

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscognito"
)

// Experimental.
type CognitoUserPoolConfiguration struct {
	// Cognito User Pool.
	// Default: - no Cognito User Pool.
	//
	// Experimental.
	UserPool awscognito.IUserPool `field:"required" json:"userPool" yaml:"userPool"`
	// Client identifiers.
	// Default: - empty list.
	//
	// Experimental.
	ClientIds *[]*string `field:"optional" json:"clientIds" yaml:"clientIds"`
	// Cognito Group Configuration.
	// Default: - no Cognito Group configuration provided.
	//
	// Experimental.
	GroupConfiguration *CognitoGroupConfiguration `field:"optional" json:"groupConfiguration" yaml:"groupConfiguration"`
}

