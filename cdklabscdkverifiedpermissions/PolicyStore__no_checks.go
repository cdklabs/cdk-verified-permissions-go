//go:build no_runtime_type_checking

package cdklabscdkverifiedpermissions

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_PolicyStore) validateAddPoliciesParameters(policyDefinitions *[]*AddPolicyOptions) error {
	return nil
}

func (p *jsiiProxy_PolicyStore) validateAddPoliciesFromPathParameters(absolutePath *string) error {
	return nil
}

func (p *jsiiProxy_PolicyStore) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (p *jsiiProxy_PolicyStore) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (p *jsiiProxy_PolicyStore) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (p *jsiiProxy_PolicyStore) validateGrantParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (p *jsiiProxy_PolicyStore) validateGrantAuthParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (p *jsiiProxy_PolicyStore) validateGrantReadParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (p *jsiiProxy_PolicyStore) validateGrantWriteParameters(grantee awsiam.IGrantable) error {
	return nil
}

func validatePolicyStore_FromPolicyStoreArnParameters(scope constructs.Construct, id *string, policyStoreArn *string) error {
	return nil
}

func validatePolicyStore_FromPolicyStoreAttributesParameters(scope constructs.Construct, id *string, attrs *PolicyStoreAttributes) error {
	return nil
}

func validatePolicyStore_FromPolicyStoreIdParameters(scope constructs.Construct, id *string, policyStoreId *string) error {
	return nil
}

func validatePolicyStore_IsConstructParameters(x interface{}) error {
	return nil
}

func validatePolicyStore_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validatePolicyStore_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validatePolicyStore_SchemaFromOpenApiSpecParameters(swaggerFilePath *string) error {
	return nil
}

func validatePolicyStore_SchemaFromRestApiParameters(restApi awsapigateway.RestApi) error {
	return nil
}

func validateNewPolicyStoreParameters(scope constructs.Construct, id *string, props *PolicyStoreProps) error {
	return nil
}

