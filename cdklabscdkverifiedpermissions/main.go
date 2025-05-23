// L2 AWS CDK Constructs for Amazon Verified Permissions
package cdklabscdkverifiedpermissions

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-verified-permissions.AddPolicyOptions",
		reflect.TypeOf((*AddPolicyOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-verified-permissions.CognitoGroupConfiguration",
		reflect.TypeOf((*CognitoGroupConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-verified-permissions.CognitoUserPoolConfiguration",
		reflect.TypeOf((*CognitoUserPoolConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@cdklabs/cdk-verified-permissions.DeletionProtectionMode",
		reflect.TypeOf((*DeletionProtectionMode)(nil)).Elem(),
		map[string]interface{}{
			"ENABLED": DeletionProtectionMode_ENABLED,
			"DISABLED": DeletionProtectionMode_DISABLED,
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-verified-permissions.EntityIdentifierProperty",
		reflect.TypeOf((*EntityIdentifierProperty)(nil)).Elem(),
	)
	_jsii_.RegisterInterface(
		"@cdklabs/cdk-verified-permissions.IIdentitySource",
		reflect.TypeOf((*IIdentitySource)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "identitySourceId", GoGetter: "IdentitySourceId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_IIdentitySource{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@cdklabs/cdk-verified-permissions.IPolicy",
		reflect.TypeOf((*IPolicy)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "policyId", GoGetter: "PolicyId"},
			_jsii_.MemberProperty{JsiiProperty: "policyType", GoGetter: "PolicyType"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_IPolicy{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@cdklabs/cdk-verified-permissions.IPolicyStore",
		reflect.TypeOf((*IPolicyStore)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberMethod{JsiiMethod: "grantAuth", GoMethod: "GrantAuth"},
			_jsii_.MemberMethod{JsiiMethod: "grantRead", GoMethod: "GrantRead"},
			_jsii_.MemberMethod{JsiiMethod: "grantWrite", GoMethod: "GrantWrite"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "policyStoreArn", GoGetter: "PolicyStoreArn"},
			_jsii_.MemberProperty{JsiiProperty: "policyStoreId", GoGetter: "PolicyStoreId"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_IPolicyStore{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@cdklabs/cdk-verified-permissions.IPolicyTemplate",
		reflect.TypeOf((*IPolicyTemplate)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "policyTemplateId", GoGetter: "PolicyTemplateId"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_IPolicyTemplate{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdklabs/cdk-verified-permissions.IdentitySource",
		reflect.TypeOf((*IdentitySource)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addAudience", GoMethod: "AddAudience"},
			_jsii_.MemberMethod{JsiiMethod: "addClientId", GoMethod: "AddClientId"},
			_jsii_.MemberMethod{JsiiMethod: "addUserPoolClient", GoMethod: "AddUserPoolClient"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "audiencesOIDC", GoGetter: "AudiencesOIDC"},
			_jsii_.MemberProperty{JsiiProperty: "clientIds", GoGetter: "ClientIds"},
			_jsii_.MemberProperty{JsiiProperty: "cognitoGroupEntityType", GoGetter: "CognitoGroupEntityType"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "groupConfigGroupClaimOIDC", GoGetter: "GroupConfigGroupClaimOIDC"},
			_jsii_.MemberProperty{JsiiProperty: "groupConfigGroupEntityTypeOIDC", GoGetter: "GroupConfigGroupEntityTypeOIDC"},
			_jsii_.MemberProperty{JsiiProperty: "identitySourceId", GoGetter: "IdentitySourceId"},
			_jsii_.MemberProperty{JsiiProperty: "issuer", GoGetter: "Issuer"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "policyStore", GoGetter: "PolicyStore"},
			_jsii_.MemberProperty{JsiiProperty: "principalIdClaimOIDC", GoGetter: "PrincipalIdClaimOIDC"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "userPoolArn", GoGetter: "UserPoolArn"},
		},
		func() interface{} {
			j := jsiiProxy_IdentitySource{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IIdentitySource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-verified-permissions.IdentitySourceAttributes",
		reflect.TypeOf((*IdentitySourceAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-verified-permissions.IdentitySourceConfiguration",
		reflect.TypeOf((*IdentitySourceConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-verified-permissions.IdentitySourceProps",
		reflect.TypeOf((*IdentitySourceProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-verified-permissions.OpenIdConnectAccessTokenConfiguration",
		reflect.TypeOf((*OpenIdConnectAccessTokenConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-verified-permissions.OpenIdConnectConfiguration",
		reflect.TypeOf((*OpenIdConnectConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-verified-permissions.OpenIdConnectGroupConfiguration",
		reflect.TypeOf((*OpenIdConnectGroupConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-verified-permissions.OpenIdConnectIdentityTokenConfiguration",
		reflect.TypeOf((*OpenIdConnectIdentityTokenConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdklabs/cdk-verified-permissions.Policy",
		reflect.TypeOf((*Policy)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "definition", GoGetter: "Definition"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "policyId", GoGetter: "PolicyId"},
			_jsii_.MemberProperty{JsiiProperty: "policyStoreId", GoGetter: "PolicyStoreId"},
			_jsii_.MemberProperty{JsiiProperty: "policyType", GoGetter: "PolicyType"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Policy{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IPolicy)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-verified-permissions.PolicyAttributes",
		reflect.TypeOf((*PolicyAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-verified-permissions.PolicyDefinitionProperty",
		reflect.TypeOf((*PolicyDefinitionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-verified-permissions.PolicyProps",
		reflect.TypeOf((*PolicyProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdklabs/cdk-verified-permissions.PolicyStore",
		reflect.TypeOf((*PolicyStore)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addPolicies", GoMethod: "AddPolicies"},
			_jsii_.MemberMethod{JsiiMethod: "addPoliciesFromPath", GoMethod: "AddPoliciesFromPath"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "deletionProtection", GoGetter: "DeletionProtection"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberMethod{JsiiMethod: "grantAuth", GoMethod: "GrantAuth"},
			_jsii_.MemberMethod{JsiiMethod: "grantRead", GoMethod: "GrantRead"},
			_jsii_.MemberMethod{JsiiMethod: "grantWrite", GoMethod: "GrantWrite"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "policyStoreArn", GoGetter: "PolicyStoreArn"},
			_jsii_.MemberProperty{JsiiProperty: "policyStoreId", GoGetter: "PolicyStoreId"},
			_jsii_.MemberProperty{JsiiProperty: "policyStoreName", GoGetter: "PolicyStoreName"},
			_jsii_.MemberProperty{JsiiProperty: "schema", GoGetter: "Schema"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "validationSettings", GoGetter: "ValidationSettings"},
		},
		func() interface{} {
			j := jsiiProxy_PolicyStore{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IPolicyStore)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-verified-permissions.PolicyStoreAttributes",
		reflect.TypeOf((*PolicyStoreAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-verified-permissions.PolicyStoreProps",
		reflect.TypeOf((*PolicyStoreProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdklabs/cdk-verified-permissions.PolicyTemplate",
		reflect.TypeOf((*PolicyTemplate)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "policyStore", GoGetter: "PolicyStore"},
			_jsii_.MemberProperty{JsiiProperty: "policyTemplateId", GoGetter: "PolicyTemplateId"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "statement", GoGetter: "Statement"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_PolicyTemplate{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IPolicyTemplate)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-verified-permissions.PolicyTemplateAttributes",
		reflect.TypeOf((*PolicyTemplateAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-verified-permissions.PolicyTemplateProps",
		reflect.TypeOf((*PolicyTemplateProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@cdklabs/cdk-verified-permissions.PolicyType",
		reflect.TypeOf((*PolicyType)(nil)).Elem(),
		map[string]interface{}{
			"STATIC": PolicyType_STATIC,
			"TEMPLATELINKED": PolicyType_TEMPLATELINKED,
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-verified-permissions.Schema",
		reflect.TypeOf((*Schema)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-verified-permissions.StaticPolicyDefinitionProperty",
		reflect.TypeOf((*StaticPolicyDefinitionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-verified-permissions.StaticPolicyFromFileProps",
		reflect.TypeOf((*StaticPolicyFromFileProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-verified-permissions.Tag",
		reflect.TypeOf((*Tag)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-verified-permissions.TemplateFromFileProps",
		reflect.TypeOf((*TemplateFromFileProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-verified-permissions.TemplateLinkedPolicyDefinitionProperty",
		reflect.TypeOf((*TemplateLinkedPolicyDefinitionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-verified-permissions.ValidationSettings",
		reflect.TypeOf((*ValidationSettings)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@cdklabs/cdk-verified-permissions.ValidationSettingsMode",
		reflect.TypeOf((*ValidationSettingsMode)(nil)).Elem(),
		map[string]interface{}{
			"OFF": ValidationSettingsMode_OFF,
			"STRICT": ValidationSettingsMode_STRICT,
		},
	)
}
