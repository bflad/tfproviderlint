package terraformtype

import (
	"go/ast"
	"go/types"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
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

// HelperSchemaResourceInfo represents all gathered Resource data for easier access
type HelperSchemaResourceInfo struct {
	AstCompositeLit *ast.CompositeLit
	Fields          map[string]*ast.KeyValueExpr
	Resource        *schema.Resource
	TypesInfo       *types.Info
}

// NewHelperSchemaResourceInfo instantiates a HelperSchemaResourceInfo
func NewHelperSchemaResourceInfo(cl *ast.CompositeLit, info *types.Info) *HelperSchemaResourceInfo {
	result := &HelperSchemaResourceInfo{
		AstCompositeLit: cl,
		Fields:          astCompositeLitFields(cl),
		Resource:        &schema.Resource{},
		TypesInfo:       info,
	}

	return result
}

// DeclaresField returns true if the field name is present in the AST
func (info *HelperSchemaResourceInfo) DeclaresField(fieldName string) bool {
	return info.Fields[fieldName] != nil
}

// HelperSchemaSchemaInfo represents all gathered Schema data for easier access
type HelperSchemaSchemaInfo struct {
	AstCompositeLit *ast.CompositeLit
	Fields          map[string]*ast.KeyValueExpr
	Schema          *schema.Schema
	SchemaValueType string
	TypesInfo       *types.Info
}

// NewHelperSchemaSchemaInfo instantiates a HelperSchemaSchemaInfo
func NewHelperSchemaSchemaInfo(cl *ast.CompositeLit, info *types.Info) *HelperSchemaSchemaInfo {
	result := &HelperSchemaSchemaInfo{
		AstCompositeLit: cl,
		Fields:          astCompositeLitFields(cl),
		Schema:          &schema.Schema{},
		SchemaValueType: helperSchemaTypeSchemaType(cl, info),
		TypesInfo:       info,
	}

	if kvExpr := result.Fields[SchemaFieldComputed]; kvExpr != nil {
		result.Schema.Computed = *astBoolValue(kvExpr.Value)
	}

	if kvExpr := result.Fields[SchemaFieldConflictsWith]; kvExpr != nil && astExprValue(kvExpr.Value) != nil {
		result.Schema.ConflictsWith = []string{}
	}

	if kvExpr := result.Fields[SchemaFieldDefault]; kvExpr != nil && astExprValue(kvExpr.Value) != nil {
		result.Schema.Default = func() (interface{}, error) { return nil, nil }
	}

	if kvExpr := result.Fields[SchemaFieldDiffSuppressFunc]; kvExpr != nil && astExprValue(kvExpr.Value) != nil {
		result.Schema.DiffSuppressFunc = func(k, old, new string, d *schema.ResourceData) bool { return false }
	}

	if kvExpr := result.Fields[SchemaFieldForceNew]; kvExpr != nil {
		result.Schema.ForceNew = *astBoolValue(kvExpr.Value)
	}

	if kvExpr := result.Fields[SchemaFieldMaxItems]; kvExpr != nil {
		result.Schema.MaxItems = *astIntValue(kvExpr.Value)
	}

	if kvExpr := result.Fields[SchemaFieldMinItems]; kvExpr != nil {
		result.Schema.MinItems = *astIntValue(kvExpr.Value)
	}

	if kvExpr := result.Fields[SchemaFieldOptional]; kvExpr != nil {
		result.Schema.Optional = *astBoolValue(kvExpr.Value)
	}

	if kvExpr := result.Fields[SchemaFieldRequired]; kvExpr != nil {
		result.Schema.Required = *astBoolValue(kvExpr.Value)
	}

	if kvExpr := result.Fields[SchemaFieldSensitive]; kvExpr != nil {
		result.Schema.Sensitive = *astBoolValue(kvExpr.Value)
	}

	if kvExpr := result.Fields[SchemaFieldValidateFunc]; kvExpr != nil && astExprValue(kvExpr.Value) != nil {
		result.Schema.ValidateFunc = func(interface{}, string) ([]string, []error) { return nil, nil }
	}

	return result
}

// DeclaresField returns true if the field name is present in the AST
func (info *HelperSchemaSchemaInfo) DeclaresField(fieldName string) bool {
	return info.Fields[fieldName] != nil
}

// DeclaresBoolFieldWithZeroValue returns true if the field name is present and is false
func (info *HelperSchemaSchemaInfo) DeclaresBoolFieldWithZeroValue(fieldName string) bool {
	kvExpr := info.Fields[fieldName]

	// Field not declared
	if kvExpr == nil {
		return false
	}

	valuePtr := astBoolValue(kvExpr.Value)

	// Value not readable
	if valuePtr == nil {
		return false
	}

	return !*valuePtr
}

// IsType returns true if the given input is equal to the Type
func (info *HelperSchemaSchemaInfo) IsType(valueType string) bool {
	return info.SchemaValueType == valueType
}

// IsOneOfTypes returns true if one of the given input is equal to the Type
func (info *HelperSchemaSchemaInfo) IsOneOfTypes(valueTypes ...string) bool {
	for _, valueType := range valueTypes {
		if info.SchemaValueType == valueType {
			return true
		}
	}

	return false
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

// helperSchemaTypeSchemaType extracts the string representation of a Schema Type value
func helperSchemaTypeSchemaType(schema *ast.CompositeLit, info *types.Info) string {
	kvExpr := astCompositeLitField(schema, SchemaFieldType)

	if kvExpr == nil {
		return ""
	}

	if !IsHelperSchemaValueType(kvExpr.Value, info) {
		return ""
	}

	return helperSchemaValueTypeString(kvExpr.Value)
}

// helperSchemaValueTypeString extracts the string representation of a Schema ValueType
func helperSchemaValueTypeString(e ast.Expr) string {
	switch e := e.(type) {
	case *ast.SelectorExpr:
		return e.Sel.Name
	default:
		return ""
	}
}
