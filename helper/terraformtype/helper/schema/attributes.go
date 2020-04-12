package schema

import (
	"fmt"
	tfschema "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"math"
	"regexp"
	"strconv"
	"strings"
)

const (
	// Pattern for schema attribute names
	AttributeNameRegexpPattern = `^[a-z0-9_]+$`

	// Pattern for schema references to attributes, such as ConflictsWith values
	AttributeReferenceRegexpPattern = `^[a-z0-9_]+(\.[a-z0-9_]+)*$`
)

var (
	AttributeNameRegexp      = regexp.MustCompile(AttributeNameRegexpPattern)
	AttributeReferenceRegexp = regexp.MustCompile(AttributeReferenceRegexpPattern)
)

// ParseAttributeReference validates and returns the split representation of schema attribute reference.
// Attribute references are used in Schema fields such as AtLeastOneOf, ConflictsWith, and ExactlyOneOf.
func ParseAttributeReference(reference string) ([]string, error) {
	if !AttributeReferenceRegexp.MatchString(reference) {
		return nil, fmt.Errorf("%q must contain only valid attribute names, separated by periods", reference)
	}

	attributeReferenceParts := strings.Split(reference, ".")

	if len(attributeReferenceParts) == 1 {
		return attributeReferenceParts, nil
	}

	configurationBlockReferenceErr := fmt.Errorf("%q configuration block attribute references are only valid for TypeList and MaxItems: 1 attributes and nested attributes must be separated by .0.", reference)

	if math.Mod(float64(len(attributeReferenceParts)), 2) == 0 {
		return attributeReferenceParts, configurationBlockReferenceErr
	}

	// All even parts of an attribute reference must be 0
	for idx, attributeReferencePart := range attributeReferenceParts {
		if math.Mod(float64(idx), 2) == 0 {
			continue
		}

		attributeReferencePartInt, err := strconv.Atoi(attributeReferencePart)

		if err != nil {
			return attributeReferenceParts, configurationBlockReferenceErr
		}

		if attributeReferencePartInt != 0 {
			return attributeReferenceParts, configurationBlockReferenceErr
		}
	}

	return attributeReferenceParts, nil
}

// ValidateAttributeReference validates schema attribute reference.
// Attribute references are used in Schema fields such as AtLeastOneOf, ConflictsWith, and ExactlyOneOf.
func ValidateAttributeReference(tfresource *tfschema.Resource, reference string) ([]string, error) {
	var curSchemaOrResource interface{} = tfresource
	attributeReferenceParts := strings.Split(reference, ".")
	for idx, attributeReferencePart := range attributeReferenceParts {
		attributeReference := strings.Join(attributeReferenceParts[:idx+1], ".")
		if math.Mod(float64(idx), 2) == 1 {
			// For odd part, ensure it is a 0 and the containing schema has `MaxItems` set to `1`. This is because
			// reference among multiple instances of the same nested block is not supported in current plugin SDK.
			attributeReferencePartInt, err := strconv.Atoi(attributeReferencePart)

			configurationBlockReferenceErr := fmt.Errorf("%q configuration block attribute references must be separated by .0", attributeReference)
			if err != nil {
				return nil, configurationBlockReferenceErr
			}
			if attributeReferencePartInt != 0 {
				return nil, configurationBlockReferenceErr
			}

			curSchema := curSchemaOrResource.(*tfschema.Schema)
			if curSchema.MaxItems != 1 || curSchema.Type != tfschema.TypeList {
				return nil, fmt.Errorf("%q configuration block attribute references are only valid for TypeList and MaxItems: 1 attributes", attributeReference)
			}
			curSchemaOrResource = curSchema.Elem
		} else {
			// For even part, ensure it references to defined attribute
			schema := curSchemaOrResource.(*tfschema.Resource).Schema[attributeReferencePart]
			if schema == nil {
				return nil, fmt.Errorf("%q references to unknown attribute", attributeReference)
			}
			curSchemaOrResource = schema
		}
	}

	return attributeReferenceParts, nil
}

