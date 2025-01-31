package cdklabscdkverifiedpermissions

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/cdklabs/cdk-verified-permissions-go/cdklabscdkverifiedpermissions/internal"
)

// Experimental.
type IIdentitySource interface {
	awscdk.IResource
	// Identity Source identifier.
	// Experimental.
	IdentitySourceId() *string
}

// The jsii proxy for IIdentitySource
type jsiiProxy_IIdentitySource struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IIdentitySource) IdentitySourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identitySourceId",
		&returns,
	)
	return returns
}

