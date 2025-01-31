//go:build no_runtime_type_checking

package cdklabscdkverifiedpermissions

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IPolicyStore) validateGrantParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_IPolicyStore) validateGrantAuthParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_IPolicyStore) validateGrantReadParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_IPolicyStore) validateGrantWriteParameters(grantee awsiam.IGrantable) error {
	return nil
}

