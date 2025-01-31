package cdklabscdkverifiedpermissions

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/cdklabs/cdk-verified-permissions-go/cdklabscdkverifiedpermissions/internal"
)

// Experimental.
type IPolicyTemplate interface {
	awscdk.IResource
	// The ID of the policy template.
	// Experimental.
	PolicyTemplateId() *string
}

// The jsii proxy for IPolicyTemplate
type jsiiProxy_IPolicyTemplate struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IPolicyTemplate) PolicyTemplateId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyTemplateId",
		&returns,
	)
	return returns
}

