//go:build no_runtime_type_checking

package cdklabscdkverifiedpermissions

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_Policy) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (p *jsiiProxy_Policy) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (p *jsiiProxy_Policy) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validatePolicy_FromFileParameters(scope constructs.Construct, id *string, props *StaticPolicyFromFileProps) error {
	return nil
}

func validatePolicy_FromPolicyAttributesParameters(scope constructs.Construct, id *string, attrs *PolicyAttributes) error {
	return nil
}

func validatePolicy_FromPolicyIdParameters(scope constructs.Construct, id *string, policyId *string) error {
	return nil
}

func validatePolicy_IsConstructParameters(x interface{}) error {
	return nil
}

func validatePolicy_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validatePolicy_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewPolicyParameters(scope constructs.Construct, id *string, props *PolicyProps) error {
	return nil
}

