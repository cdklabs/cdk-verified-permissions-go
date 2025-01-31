# Amazon Verified Permissions L2 CDK Construct

This repo contains the implementation of an L2 CDK Construct for Amazon Verified Permissions

# Project Stability

This construct is still versioned with alpha/v0 major version and we could introduce breaking changes even without a major version bump. Our goal is to keep the API stable & backwards compatible as much as possible but we currently cannot guarantee that. Once we'll publish v1.0.0 the breaking changes will be introduced via major version bumps.

# Getting Started

## Policy Store

Define a Policy Store with defaults (No description, No schema & Validation Settings Mode set to OFF):

```go
test := cdklabscdkverifiedpermissions.NewPolicyStore(*scope, jsii.String("PolicyStore"))
```

Define a Policy Store without Schema definition (Validation Settings Mode must be set to OFF):

```go
validationSettingsOff := map[string]validationSettingsMode{
	"mode": cdklabscdkverifiedpermissions.ValidationSettingsMode_OFF,
}
test := cdklabscdkverifiedpermissions.NewPolicyStore(*scope, jsii.String("PolicyStore"), &PolicyStoreProps{
	ValidationSettings: validationSettingsOff,
})
```

Define a Policy Store with Description and Schema definition (a STRICT Validation Settings Mode is strongly suggested for Policy Stores with schemas):

```go
validationSettingsStrict := map[string]validationSettingsMode{
	"mode": cdklabscdkverifiedpermissions.ValidationSettingsMode_STRICT,
}
cedarJsonSchema := map[string]map[string]map[string]map[string]interface{}{
	"PhotoApp": map[string]map[string]map[string]interface{}{
		"entityTypes": map[string]map[string]interface{}{
			"User": map[string]interface{}{
			},
			"Photo": map[string]interface{}{
			},
		},
		"actions": map[string]map[string]map[string][]*string{
			"viewPhoto": map[string]map[string][]*string{
				"appliesTo": map[string][]*string{
					"principalTypes": []*string{
						jsii.String("User"),
					},
					"resourceTypes": []*string{
						jsii.String("Photo"),
					},
				},
			},
		},
	},
}
cedarSchema := map[string]*string{
	"cedarJson": JSON.stringify(cedarJsonSchema),
}
policyStore := cdklabscdkverifiedpermissions.NewPolicyStore(*scope, jsii.String("PolicyStore"), &PolicyStoreProps{
	Schema: cedarSchema,
	ValidationSettings: validationSettingsStrict,
	Description: jsii.String("PolicyStore description"),
})
```

## Schemas

If you want to have type safety when defining a schema, you can accomplish this **<ins>only</ins>** in typescript. Simply use the `Schema` type exported by the `@cedar-policy/cedar-wasm`.

You can also generate simple schemas using the static functions `schemaFromOpenApiSpec` or `schemaFromRestApi` in the PolicyStore construct. This functionality replicates what you can find in the AWS Verified Permissions console.

Generate a schema from an OpenAPI spec:

```go
validationSettingsStrict := map[string]validationSettingsMode{
	"mode": cdklabscdkverifiedpermissions.ValidationSettingsMode_STRICT,
}
cedarJsonSchema := cdklabscdkverifiedpermissions.PolicyStore_SchemaFromOpenApiSpec(jsii.String("path/to/swaggerfile.json"), jsii.String("UserGroup"))
cedarSchema := map[string]*string{
	"cedarJson": JSON.stringify(cedarJsonSchema),
}
policyStore := cdklabscdkverifiedpermissions.NewPolicyStore(*scope, jsii.String("PolicyStore"), &PolicyStoreProps{
	Schema: cedarSchema,
	ValidationSettings: validationSettingsStrict,
	Description: jsii.String("Policy store with schema generated from API Gateway"),
})
```

Generate a schema from a RestApi construct:

```go
validationSettingsStrict := map[string]validationSettingsMode{
	"mode": cdklabscdkverifiedpermissions.ValidationSettingsMode_STRICT,
}
cedarJsonSchema := cdklabscdkverifiedpermissions.PolicyStore_SchemaFromRestApi(
awscdk.NewRestApi(*scope, jsii.String("RestApi")), jsii.String("UserGroup"))
cedarSchema := map[string]*string{
	"cedarJson": JSON.stringify(cedarJsonSchema),
}
policyStore := cdklabscdkverifiedpermissions.NewPolicyStore(*scope, jsii.String("PolicyStore"), &PolicyStoreProps{
	Schema: cedarSchema,
	ValidationSettings: validationSettingsStrict,
	Description: jsii.String("Policy store with schema generated from RestApi construct"),
})
```

## Identity Source

Define Identity Source with Cognito Configuration and required properties:

```go
userPool := awscdk.NewUserPool(*scope, jsii.String("UserPool")) // Creating a new Cognito UserPool
validationSettingsStrict := map[string]validationSettingsMode{
	"mode": cdklabscdkverifiedpermissions.ValidationSettingsMode_STRICT,
}
cedarJsonSchema := map[string]map[string]map[string]map[string]interface{}{
	"PhotoApp": map[string]map[string]map[string]interface{}{
		"entityTypes": map[string]map[string]interface{}{
			"User": map[string]interface{}{
			},
			"Photo": map[string]interface{}{
			},
		},
		"actions": map[string]map[string]map[string][]*string{
			"viewPhoto": map[string]map[string][]*string{
				"appliesTo": map[string][]*string{
					"principalTypes": []*string{
						jsii.String("User"),
					},
					"resourceTypes": []*string{
						jsii.String("Photo"),
					},
				},
			},
		},
	},
}
cedarSchema := map[string]*string{
	"cedarJson": JSON.stringify(cedarJsonSchema),
}
policyStore := cdklabscdkverifiedpermissions.NewPolicyStore(*scope, jsii.String("PolicyStore"), &PolicyStoreProps{
	Schema: cedarSchema,
	ValidationSettings: validationSettingsStrict,
})
cdklabscdkverifiedpermissions.NewIdentitySource(*scope, jsii.String("IdentitySource"), &IdentitySourceProps{
	Configuration: &IdentitySourceConfiguration{
		CognitoUserPoolConfiguration: &CognitoUserPoolConfiguration{
			UserPool: userPool,
		},
	},
	PolicyStore: policyStore,
})
```

Define Identity Source with Cognito Configuration and all properties:

```go
validationSettingsStrict := map[string]validationSettingsMode{
	"mode": cdklabscdkverifiedpermissions.ValidationSettingsMode_STRICT,
}
cedarJsonSchema := map[string]map[string]map[string]map[string]interface{}{
	"PhotoApp": map[string]map[string]map[string]interface{}{
		"entityTypes": map[string]map[string]interface{}{
			"User": map[string]interface{}{
			},
			"Photo": map[string]interface{}{
			},
		},
		"actions": map[string]map[string]map[string][]*string{
			"viewPhoto": map[string]map[string][]*string{
				"appliesTo": map[string][]*string{
					"principalTypes": []*string{
						jsii.String("User"),
					},
					"resourceTypes": []*string{
						jsii.String("Photo"),
					},
				},
			},
		},
	},
}
cedarSchema := map[string]*string{
	"cedarJson": JSON.stringify(cedarJsonSchema),
}
policyStore := cdklabscdkverifiedpermissions.NewPolicyStore(*scope, jsii.String("PolicyStore"), &PolicyStoreProps{
	Schema: cedarSchema,
	ValidationSettings: validationSettingsStrict,
})
cognitoGroupEntityType := "test"
userPool := awscdk.NewUserPool(*scope, jsii.String("UserPool")) // Creating a new Cognito UserPool
 // Creating a new Cognito UserPool
cdklabscdkverifiedpermissions.NewIdentitySource(*scope, jsii.String("IdentitySource"), &IdentitySourceProps{
	Configuration: &IdentitySourceConfiguration{
		CognitoUserPoolConfiguration: &CognitoUserPoolConfiguration{
			ClientIds: []*string{
				jsii.String("&ExampleCogClientId;"),
			},
			UserPool: userPool,
			GroupConfiguration: &CognitoGroupConfiguration{
				GroupEntityType: cognitoGroupEntityType,
			},
		},
	},
	PolicyStore: policyStore,
	PrincipalEntityType: jsii.String("PETEXAMPLEabcdefg111111"),
})
```

Define Identity Source with OIDC Configuration and Access Token selection config:

```go
validationSettingsStrict := map[string]validationSettingsMode{
	"mode": cdklabscdkverifiedpermissions.ValidationSettingsMode_STRICT,
}
cedarJsonSchema := map[string]map[string]map[string]map[string]interface{}{
	"PhotoApp": map[string]map[string]map[string]interface{}{
		"entityTypes": map[string]map[string]interface{}{
			"User": map[string]interface{}{
			},
			"Photo": map[string]interface{}{
			},
		},
		"actions": map[string]map[string]map[string][]*string{
			"viewPhoto": map[string]map[string][]*string{
				"appliesTo": map[string][]*string{
					"principalTypes": []*string{
						jsii.String("User"),
					},
					"resourceTypes": []*string{
						jsii.String("Photo"),
					},
				},
			},
		},
	},
}
cedarSchema := map[string]*string{
	"cedarJson": JSON.stringify(cedarJsonSchema),
}
policyStore := cdklabscdkverifiedpermissions.NewPolicyStore(*scope, jsii.String("PolicyStore"), &PolicyStoreProps{
	Schema: cedarSchema,
	ValidationSettings: validationSettingsStrict,
})
issuer := "https://iamanidp.com"
principalIdClaim := "sub"
entityIdPrefix := "prefix"
groupClaim := "group"
groupEntityType := "GroupType"
cdklabscdkverifiedpermissions.NewIdentitySource(*scope, jsii.String("IdentitySource"), &IdentitySourceProps{
	Configuration: &IdentitySourceConfiguration{
		OpenIdConnectConfiguration: &OpenIdConnectConfiguration{
			Issuer: issuer,
			EntityIdPrefix: entityIdPrefix,
			GroupConfiguration: &OpenIdConnectGroupConfiguration{
				GroupClaim: groupClaim,
				GroupEntityType: groupEntityType,
			},
			AccessTokenOnly: &OpenIdConnectAccessTokenConfiguration{
				Audiences: []*string{
					jsii.String("testAudience"),
				},
				PrincipalIdClaim: principalIdClaim,
			},
		},
	},
	PolicyStore: policyStore,
	PrincipalEntityType: jsii.String("TestType"),
})
```

Define Identity Source with OIDC Configuration and Identity Token selection config:

```go
validationSettingsStrict := map[string]validationSettingsMode{
	"mode": cdklabscdkverifiedpermissions.ValidationSettingsMode_STRICT,
}
cedarJsonSchema := map[string]map[string]map[string]map[string]interface{}{
	"PhotoApp": map[string]map[string]map[string]interface{}{
		"entityTypes": map[string]map[string]interface{}{
			"User": map[string]interface{}{
			},
			"Photo": map[string]interface{}{
			},
		},
		"actions": map[string]map[string]map[string][]*string{
			"viewPhoto": map[string]map[string][]*string{
				"appliesTo": map[string][]*string{
					"principalTypes": []*string{
						jsii.String("User"),
					},
					"resourceTypes": []*string{
						jsii.String("Photo"),
					},
				},
			},
		},
	},
}
cedarSchema := map[string]*string{
	"cedarJson": JSON.stringify(cedarJsonSchema),
}
policyStore := cdklabscdkverifiedpermissions.NewPolicyStore(*scope, jsii.String("PolicyStore"), &PolicyStoreProps{
	Schema: cedarSchema,
	ValidationSettings: validationSettingsStrict,
})
issuer := "https://iamanidp.com"
entityIdPrefix := "prefix"
groupClaim := "group"
groupEntityType := "UserGroup"
principalIdClaim := "sub"
cdklabscdkverifiedpermissions.NewIdentitySource(*scope, jsii.String("IdentitySource"), &IdentitySourceProps{
	Configuration: &IdentitySourceConfiguration{
		OpenIdConnectConfiguration: &OpenIdConnectConfiguration{
			Issuer: issuer,
			EntityIdPrefix: entityIdPrefix,
			GroupConfiguration: &OpenIdConnectGroupConfiguration{
				GroupClaim: groupClaim,
				GroupEntityType: groupEntityType,
			},
			IdentityTokenOnly: &OpenIdConnectIdentityTokenConfiguration{
				ClientIds: []*string{
				},
				PrincipalIdClaim: principalIdClaim,
			},
		},
	},
	PolicyStore: policyStore,
})
```

## Policy

Load all the `.cedar` files in a given folder and define Policy objects for each of them. All policies will be associated with the same policy store.

```go
validationSettingsStrict := map[string]validationSettingsMode{
	"mode": cdklabscdkverifiedpermissions.ValidationSettingsMode_STRICT,
}
policyStore := cdklabscdkverifiedpermissions.NewPolicyStore(*scope, jsii.String("PolicyStore"), &PolicyStoreProps{
	ValidationSettings: validationSettingsStrict,
})
policyStore.AddPoliciesFromPath(jsii.String("/path/to/my-policies"))
```

Define a Policy and add it to a specific Policy Store:

```go
statement := `permit(
    principal,
    action in [MyFirstApp::Action::"Read"],
    resource
) when {
    true
};`

description := "Test policy assigned to the test store"
validationSettingsOff := map[string]validationSettingsMode{
	"mode": cdklabscdkverifiedpermissions.ValidationSettingsMode_OFF,
}
policyStore := cdklabscdkverifiedpermissions.NewPolicyStore(*scope, jsii.String("PolicyStore"), &PolicyStoreProps{
	ValidationSettings: validationSettingsOff,
})

// Create a policy and add it to the policy store
policy := cdklabscdkverifiedpermissions.NewPolicy(*scope, jsii.String("MyTestPolicy"), &PolicyProps{
	Definition: &PolicyDefinitionProperty{
		Static: &StaticPolicyDefinitionProperty{
			Statement: jsii.String(*Statement),
			Description: jsii.String(*Description),
		},
	},
	PolicyStore: policyStore,
})
```

Define a policy with a template linked definition:

```go
validationSettingsOff := map[string]validationSettingsMode{
	"mode": cdklabscdkverifiedpermissions.ValidationSettingsMode_OFF,
}
policyStore := cdklabscdkverifiedpermissions.NewPolicyStore(*scope, jsii.String("PolicyStore"), &PolicyStoreProps{
	ValidationSettings: validationSettingsOff,
})
policyTemplateStatement := `
permit (
  principal == ?principal,
  action in [TinyTodo::Action::"ReadList", TinyTodo::Action::"ListTasks"],
  resource == ?resource
);`
template := cdklabscdkverifiedpermissions.NewPolicyTemplate(*scope, jsii.String("PolicyTemplate"), &PolicyTemplateProps{
	Statement: policyTemplateStatement,
	PolicyStore: policyStore,
})

policy := cdklabscdkverifiedpermissions.NewPolicy(*scope, jsii.String("MyTestPolicy"), &PolicyProps{
	Definition: &PolicyDefinitionProperty{
		TemplateLinked: &TemplateLinkedPolicyDefinitionProperty{
			PolicyTemplate: template,
			Principal: &EntityIdentifierProperty{
				EntityId: jsii.String("exampleId"),
				EntityType: jsii.String("exampleType"),
			},
			Resource: &EntityIdentifierProperty{
				EntityId: jsii.String("exampleId"),
				EntityType: jsii.String("exampleType"),
			},
		},
	},
	PolicyStore: policyStore,
})
```

Define a Policy with a statement from file:
**PLEASE NOTE:** You can specify the description of the policy directly inside the Policy file, using the annotation `@cdkDescription`

```go
description := "Test policy assigned to the test store"
validationSettingsOff := map[string]validationSettingsMode{
	"mode": cdklabscdkverifiedpermissions.ValidationSettingsMode_OFF,
}
policyStore := cdklabscdkverifiedpermissions.NewPolicyStore(*scope, jsii.String("PolicyStore"), &PolicyStoreProps{
	ValidationSettings: validationSettingsOff,
})

// Create a policy and add it to the policy store
policyFromFileProps := map[string]interface{}{
	"policyStore": policyStore,
	"path": jsii.String("/path/to/policy-statement.cedar"),
	"description": jsii.String("the policy description"),
}
policy := cdklabscdkverifiedpermissions.Policy_FromFile(*scope, jsii.String("MyTestPolicy"), policyFromFileProps)
```

## Policy Template

Define a Policy Template referring to a Cedar Statement in local file:

```go
validationSettingsOff := map[string]validationSettingsMode{
	"mode": cdklabscdkverifiedpermissions.ValidationSettingsMode_OFF,
}
policyStore := cdklabscdkverifiedpermissions.NewPolicyStore(*scope, jsii.String("PolicyStore"), &PolicyStoreProps{
	ValidationSettings: validationSettingsOff,
})
templateFromFileProps := map[string]interface{}{
	"policyStore": policyStore,
	"path": jsii.String("/path/to/template-statement.cedar"),
	"description": jsii.String("Allows sharing photos in full access mode"),
}
template := cdklabscdkverifiedpermissions.PolicyTemplate_FromFile(*scope, jsii.String("PolicyTemplate"), templateFromFileProps)
```

# Notes

* This project is following the AWS CDK Official Design Guidelines (see https://github.com/aws/aws-cdk/blob/main/docs/DESIGN_GUIDELINES.md) and the AWS CDK New Constructs Creation Guide (see here https://github.com/aws/aws-cdk/blob/main/docs/NEW_CONSTRUCTS_GUIDE.md).
* Feedback is a gift: if you find something wrong or you've ideas to improve please open an issue or a pull request
