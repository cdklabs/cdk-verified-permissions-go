package cdklabscdkverifiedpermissions

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/cdklabs/cdk-verified-permissions-go/cdklabscdkverifiedpermissions/internal"
)

// Experimental.
type IPolicy interface {
	awscdk.IResource
	// The unique ID of the new or updated policy.
	// Experimental.
	PolicyId() *string
	// The type of the policy.
	//
	// This is one of the following values: Static or TemplateLinked.
	// Experimental.
	PolicyType() PolicyType
}

// The jsii proxy for IPolicy
type jsiiProxy_IPolicy struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IPolicy) PolicyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPolicy) PolicyType() PolicyType {
	var returns PolicyType
	_jsii_.Get(
		j,
		"policyType",
		&returns,
	)
	return returns
}

