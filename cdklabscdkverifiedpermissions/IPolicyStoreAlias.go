package cdklabscdkverifiedpermissions

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/cdklabs/cdk-verified-permissions-go/cdklabscdkverifiedpermissions/internal"
)

// Experimental.
type IPolicyStoreAlias interface {
	awscdk.IResource
	// The name of the policy store alias.
	// Experimental.
	AliasName() *string
}

// The jsii proxy for IPolicyStoreAlias
type jsiiProxy_IPolicyStoreAlias struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IPolicyStoreAlias) AliasName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aliasName",
		&returns,
	)
	return returns
}

