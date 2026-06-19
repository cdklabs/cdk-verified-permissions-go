//go:build no_runtime_type_checking

package cdklabscdkverifiedpermissions

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_PolicyStoreAlias) validateApplyCrossStackReferenceStrengthParameters(strength awscdk.ReferenceStrength) error {
	return nil
}

func (p *jsiiProxy_PolicyStoreAlias) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (p *jsiiProxy_PolicyStoreAlias) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (p *jsiiProxy_PolicyStoreAlias) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validatePolicyStoreAlias_FromAliasNameParameters(scope constructs.Construct, id *string, aliasName *string) error {
	return nil
}

func validatePolicyStoreAlias_FromPolicyStoreAliasAttributesParameters(scope constructs.Construct, id *string, attrs *PolicyStoreAliasAttributes) error {
	return nil
}

func validatePolicyStoreAlias_IsConstructParameters(x interface{}) error {
	return nil
}

func validatePolicyStoreAlias_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validatePolicyStoreAlias_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewPolicyStoreAliasParameters(scope constructs.Construct, id *string, props *PolicyStoreAliasProps) error {
	return nil
}

