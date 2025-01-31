//go:build !no_runtime_type_checking

package cdklabscdkverifiedpermissions

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/constructs-go/constructs/v10"
)

func (p *jsiiProxy_PolicyStore) validateAddPoliciesParameters(policyDefinitions *[]*AddPolicyOptions) error {
	if policyDefinitions == nil {
		return fmt.Errorf("parameter policyDefinitions is required, but nil was provided")
	}
	for idx_830d5d, v := range *policyDefinitions {
		if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter policyDefinitions[%#v]", idx_830d5d) }); err != nil {
			return err
		}
	}

	return nil
}

func (p *jsiiProxy_PolicyStore) validateAddPoliciesFromPathParameters(absolutePath *string) error {
	if absolutePath == nil {
		return fmt.Errorf("parameter absolutePath is required, but nil was provided")
	}

	return nil
}

func (p *jsiiProxy_PolicyStore) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	if policy == "" {
		return fmt.Errorf("parameter policy is required, but nil was provided")
	}

	return nil
}

func (p *jsiiProxy_PolicyStore) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	if arnAttr == nil {
		return fmt.Errorf("parameter arnAttr is required, but nil was provided")
	}

	if arnComponents == nil {
		return fmt.Errorf("parameter arnComponents is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(arnComponents, func() string { return "parameter arnComponents" }); err != nil {
		return err
	}

	return nil
}

func (p *jsiiProxy_PolicyStore) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	if nameAttr == nil {
		return fmt.Errorf("parameter nameAttr is required, but nil was provided")
	}

	return nil
}

func (p *jsiiProxy_PolicyStore) validateGrantParameters(grantee awsiam.IGrantable) error {
	if grantee == nil {
		return fmt.Errorf("parameter grantee is required, but nil was provided")
	}

	return nil
}

func (p *jsiiProxy_PolicyStore) validateGrantAuthParameters(grantee awsiam.IGrantable) error {
	if grantee == nil {
		return fmt.Errorf("parameter grantee is required, but nil was provided")
	}

	return nil
}

func (p *jsiiProxy_PolicyStore) validateGrantReadParameters(grantee awsiam.IGrantable) error {
	if grantee == nil {
		return fmt.Errorf("parameter grantee is required, but nil was provided")
	}

	return nil
}

func (p *jsiiProxy_PolicyStore) validateGrantWriteParameters(grantee awsiam.IGrantable) error {
	if grantee == nil {
		return fmt.Errorf("parameter grantee is required, but nil was provided")
	}

	return nil
}

func validatePolicyStore_FromPolicyStoreArnParameters(scope constructs.Construct, id *string, policyStoreArn *string) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if policyStoreArn == nil {
		return fmt.Errorf("parameter policyStoreArn is required, but nil was provided")
	}

	return nil
}

func validatePolicyStore_FromPolicyStoreAttributesParameters(scope constructs.Construct, id *string, attrs *PolicyStoreAttributes) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if attrs == nil {
		return fmt.Errorf("parameter attrs is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(attrs, func() string { return "parameter attrs" }); err != nil {
		return err
	}

	return nil
}

func validatePolicyStore_FromPolicyStoreIdParameters(scope constructs.Construct, id *string, policyStoreId *string) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if policyStoreId == nil {
		return fmt.Errorf("parameter policyStoreId is required, but nil was provided")
	}

	return nil
}

func validatePolicyStore_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validatePolicyStore_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	if construct == nil {
		return fmt.Errorf("parameter construct is required, but nil was provided")
	}

	return nil
}

func validatePolicyStore_IsResourceParameters(construct constructs.IConstruct) error {
	if construct == nil {
		return fmt.Errorf("parameter construct is required, but nil was provided")
	}

	return nil
}

func validatePolicyStore_SchemaFromOpenApiSpecParameters(swaggerFilePath *string) error {
	if swaggerFilePath == nil {
		return fmt.Errorf("parameter swaggerFilePath is required, but nil was provided")
	}

	return nil
}

func validatePolicyStore_SchemaFromRestApiParameters(restApi awsapigateway.RestApi) error {
	if restApi == nil {
		return fmt.Errorf("parameter restApi is required, but nil was provided")
	}

	return nil
}

func validateNewPolicyStoreParameters(scope constructs.Construct, id *string, props *PolicyStoreProps) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

