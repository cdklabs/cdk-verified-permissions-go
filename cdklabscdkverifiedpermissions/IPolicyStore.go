package cdklabscdkverifiedpermissions

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/cdklabs/cdk-verified-permissions-go/cdklabscdkverifiedpermissions/internal"
)

// Experimental.
type IPolicyStore interface {
	awscdk.IResource
	// Adds an IAM policy statement associated with this policy store to an IAM principal's policy.
	// Experimental.
	Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant
	// Permits an IAM principal all auth operations on the policy store: IsAuthorized, IsAuthorizedWithToken.
	// Experimental.
	GrantAuth(grantee awsiam.IGrantable) awsiam.Grant
	// Permits an IAM principal all read operations on the policy store: GetIdentitySource, GetPolicy, GetPolicyStore, GetPolicyTemplate, GetSchema, ListIdentitySources, ListPolicies, ListPolicyTemplates.
	// Experimental.
	GrantRead(grantee awsiam.IGrantable) awsiam.Grant
	// Permits an IAM principal all write & read operations on the policy store: CreateIdentitySource, CreatePolicy,CreatePolicyTemplate, DeleteIdentitySource, DeletePolicy, DeletePolicyTemplate, PutSchema, UpdateIdentitySource, UpdatePolicy, UpdatePolicyTemplate.
	// Experimental.
	GrantWrite(grantee awsiam.IGrantable) awsiam.Grant
	// ARN of the Policy Store.
	// Experimental.
	PolicyStoreArn() *string
	// ID of the Policy Store.
	// Experimental.
	PolicyStoreId() *string
}

// The jsii proxy for IPolicyStore
type jsiiProxy_IPolicyStore struct {
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IPolicyStore) Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant {
	if err := i.validateGrantParameters(grantee); err != nil {
		panic(err)
	}
	args := []interface{}{grantee}
	for _, a := range actions {
		args = append(args, a)
	}

	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grant",
		args,
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IPolicyStore) GrantAuth(grantee awsiam.IGrantable) awsiam.Grant {
	if err := i.validateGrantAuthParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantAuth",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IPolicyStore) GrantRead(grantee awsiam.IGrantable) awsiam.Grant {
	if err := i.validateGrantReadParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantRead",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IPolicyStore) GrantWrite(grantee awsiam.IGrantable) awsiam.Grant {
	if err := i.validateGrantWriteParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantWrite",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IPolicyStore) PolicyStoreArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyStoreArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPolicyStore) PolicyStoreId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyStoreId",
		&returns,
	)
	return returns
}

