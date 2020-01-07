package terraformtype

import (
	"go/ast"
	"go/types"
	"strings"
)

const (
	ResourceFieldCreate             = `Create`
	ResourceFieldCustomizeDiff      = `CustomizeDiff`
	ResourceFieldDelete             = `Delete`
	ResourceFieldDeprecationMessage = `DeprecationMessage`
	ResourceFieldExists             = `Exists`
	ResourceFieldImporter           = `Importer`
	ResourceFieldMigrateState       = `MigrateState`
	ResourceFieldRead               = `Read`
	ResourceFieldSchema             = `Schema`
	ResourceFieldSchemaVersion      = `SchemaVersion`
	ResourceFieldStateUpgraders     = `StateUpgraders`
	ResourceFieldTimeouts           = `Timeouts`
	ResourceFieldUpdate             = `Update`

	SchemaFieldAtLeastOneOf     = `AtLeastOneOf`
	SchemaFieldComputed         = `Computed`
	SchemaFieldComputedWhen     = `ComputedWhen`
	SchemaFieldConfigMode       = `ConfigMode`
	SchemaFieldConflictsWith    = `ConflictsWith`
	SchemaFieldDefault          = `Default`
	SchemaFieldDefaultFunc      = `DefaultFunc`
	SchemaFieldDeprecated       = `Deprecated`
	SchemaFieldDescription      = `Description`
	SchemaFieldDiffSuppressFunc = `DiffSuppressFunc`
	SchemaFieldElem             = `Elem`
	SchemaFieldExactlyOneOf     = `ExactlyOneOf`
	SchemaFieldForceNew         = `ForceNew`
	SchemaFieldInputDefault     = `InputDefault`
	SchemaFieldMaxItems         = `MaxItems`
	SchemaFieldMinItems         = `MinItems`
	SchemaFieldOptional         = `Optional`
	SchemaFieldPromoteSingle    = `PromoteSingle`
	SchemaFieldRemoved          = `Removed`
	SchemaFieldRequired         = `Required`
	SchemaFieldSensitive        = `Sensitive`
	SchemaFieldSet              = `Set`
	SchemaFieldStateFunc        = `StateFunc`
	SchemaFieldType             = `Type`
	SchemaFieldValidateFunc     = `ValidateFunc`

	SchemaValueTypeBool   = `TypeBool`
	SchemaValueTypeFloat  = `TypeFloat`
	SchemaValueTypeInt    = `TypeInt`
	SchemaValueTypeList   = `TypeList`
	SchemaValueTypeMap    = `TypeMap`
	SchemaValueTypeSet    = `TypeSet`
	SchemaValueTypeString = `TypeString`

	TypeNameResource     = `Resource`
	TypeNameResourceData = `ResourceData`
	TypeNameSchema       = `Schema`
	TypeNameSet          = `Set`
	TypeNameValueType    = `ValueType`

	TypePackagePathHelperSchema = `github.com/hashicorp/terraform-plugin-sdk/helper/schema`
)

// HelperSchemaTypeResourceContainsFields checks if any fields are present in the Resource
func HelperSchemaTypeResourceContainsFields(resource *ast.CompositeLit, fields ...string) bool {
	return astCompositeLitContainsAnyField(resource, fields...)
}

// HelperSchemaTypeSchemaContainsFields checks if any fields are present in the Schema
func HelperSchemaTypeSchemaContainsFields(schema *ast.CompositeLit, fields ...string) bool {
	return astCompositeLitContainsAnyField(schema, fields...)
}

// HelperSchemaTypeSchemaContainsType checks if any ValueType are present in the Schema
func HelperSchemaTypeSchemaContainsTypes(schema *ast.CompositeLit, info *types.Info, valueTypes ...string) bool {
	schemaType := HelperSchemaTypeSchemaType(schema, info)

	for _, valueType := range valueTypes {
		if schemaType == valueType {
			return true
		}
	}

	return false
}

// HelperSchemaTypeSchemaComputed extracts the Schema Computed value
func HelperSchemaTypeSchemaComputed(schema *ast.CompositeLit) *bool {
	return astCompositeLitFieldBoolValue(schema, SchemaFieldComputed)
}

// HelperSchemaTypeSchemaConflictsWith extracts the Schema ConflictsWith value
// NOTE: The current function signature output type should not be considered stable
//       except for being a pointer.
func HelperSchemaTypeSchemaConflictsWith(schema *ast.CompositeLit) *ast.Expr {
	return astCompositeLitFieldExprValue(schema, SchemaFieldConflictsWith)
}

// HelperSchemaTypeSchemaDefault extracts the Schema Default value
// NOTE: The current function signature output type should not be considered stable
//       except for being a pointer.
func HelperSchemaTypeSchemaDefault(schema *ast.CompositeLit) *ast.Expr {
	return astCompositeLitFieldExprValue(schema, SchemaFieldDefault)
}

// HelperSchemaTypeSchemaDiffSuppressFunc extracts the Schema DiffSuppressFunc value
// NOTE: The current function signature output type should not be considered stable
//       except for being a pointer.
func HelperSchemaTypeSchemaDiffSuppressFunc(schema *ast.CompositeLit) *ast.Expr {
	return astCompositeLitFieldExprValue(schema, SchemaFieldDiffSuppressFunc)
}

// HelperSchemaTypeSchemaMaxItems extracts the Schema MaxItems value
func HelperSchemaTypeSchemaMaxItems(schema *ast.CompositeLit) *int {
	return astCompositeLitFieldIntValue(schema, SchemaFieldMaxItems)
}

// HelperSchemaTypeSchemaOptional extracts the Schema Optional value
func HelperSchemaTypeSchemaOptional(schema *ast.CompositeLit) *bool {
	return astCompositeLitFieldBoolValue(schema, SchemaFieldOptional)
}

// HelperSchemaTypeSchemaRequired extracts the Schema Required value
func HelperSchemaTypeSchemaRequired(schema *ast.CompositeLit) *bool {
	return astCompositeLitFieldBoolValue(schema, SchemaFieldRequired)
}

// HelperSchemaTypeSchemaType extracts the string representation of a Schema Type value
func HelperSchemaTypeSchemaType(schema *ast.CompositeLit, info *types.Info) string {
	kvExpr := astCompositeLitField(schema, SchemaFieldType)

	if kvExpr == nil {
		return ""
	}

	if !IsHelperSchemaValueType(kvExpr.Value, info) {
		return ""
	}

	return HelperSchemaValueTypeString(kvExpr.Value)
}

// HelperSchemaTypeSchemaValidateFunc extracts the Schema ValidateFunc value
// NOTE: The current function signature output type should not be considered stable
//       except for being a pointer.
func HelperSchemaTypeSchemaValidateFunc(schema *ast.CompositeLit) *ast.Expr {
	return astCompositeLitFieldExprValue(schema, SchemaFieldValidateFunc)
}

// HelperSchemaValueTypeString extracts the string representation of a Schema ValueType
func HelperSchemaValueTypeString(e ast.Expr) string {
	switch e := e.(type) {
	case *ast.SelectorExpr:
		return e.Sel.Name
	default:
		return ""
	}
}

// IsHelperSchemaTypeResource returns if the type is Resource from the helper/schema package
func IsHelperSchemaTypeResource(t types.Type) bool {
	switch t := t.(type) {
	case *types.Named:
		return IsHelperSchemaNamedType(t, TypeNameResource)
	default:
		return false
	}
}

// IsHelperSchemaTypeResourceData returns if the type is ResourceData from the helper/schema package
func IsHelperSchemaTypeResourceData(t types.Type) bool {
	switch t := t.(type) {
	case *types.Named:
		return IsHelperSchemaNamedType(t, TypeNameResourceData)
	default:
		return false
	}
}

// IsHelperSchemaTypeSchema returns if the type is Schema from the helper/schema package
func IsHelperSchemaTypeSchema(t types.Type) bool {
	switch t := t.(type) {
	case *types.Named:
		return IsHelperSchemaNamedType(t, TypeNameSchema)
	default:
		return false
	}
}

// IsHelperSchemaValueType returns if the Schema field Type matches
func IsHelperSchemaValueType(e ast.Expr, info *types.Info) bool {
	switch e := e.(type) {
	case *ast.SelectorExpr:
		switch t := info.TypeOf(e).(type) {
		case *types.Named:
			return IsHelperSchemaNamedType(t, TypeNameValueType)
		default:
			return false
		}
	default:
		return false
	}
}

// IsHelperSchemaTypeSet returns if the type is Set from the helper/schema package
// Use IsHelperSchemaTypeSchemaFieldType for verifying Type: schema.TypeSet ValueType
func IsHelperSchemaTypeSet(t types.Type) bool {
	switch t := t.(type) {
	case *types.Named:
		return IsHelperSchemaNamedType(t, TypeNameSet)
	default:
		return false
	}
}

// IsHelperSchemaNamedType returns if the type name matches and is from the helper/schema package
func IsHelperSchemaNamedType(t *types.Named, typeName string) bool {
	if t.Obj().Name() != typeName {
		return false
	}

	// HasSuffix here due to vendoring
	return strings.HasSuffix(t.Obj().Pkg().Path(), TypePackagePathHelperSchema)
}
