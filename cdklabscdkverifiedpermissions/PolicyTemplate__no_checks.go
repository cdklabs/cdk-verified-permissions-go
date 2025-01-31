//go:build no_runtime_type_checking

package cdklabscdkverifiedpermissions

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_PolicyTemplate) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (p *jsiiProxy_PolicyTemplate) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (p *jsiiProxy_PolicyTemplate) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validatePolicyTemplate_FromFileParameters(scope constructs.Construct, id *string, props *TemplateFromFileProps) error {
	return nil
}

func validatePolicyTemplate_FromPolicyTemplateAttributesParameters(scope constructs.Construct, id *string, attrs *PolicyTemplateAttributes) error {
	return nil
}

func validatePolicyTemplate_FromPolicyTemplateIdParameters(scope constructs.Construct, id *string, policyTemplateId *string) error {
	return nil
}

func validatePolicyTemplate_IsConstructParameters(x interface{}) error {
	return nil
}

func validatePolicyTemplate_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validatePolicyTemplate_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewPolicyTemplateParameters(scope constructs.Construct, id *string, props *PolicyTemplateProps) error {
	return nil
}

