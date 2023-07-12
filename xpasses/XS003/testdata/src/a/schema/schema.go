package schema

type Resource struct {
	Schema map[string]*Schema
}

type Schema struct {
	Type ValueType

	Optional bool
	Required bool

	Default     interface{}
	DefaultFunc func() (interface{}, error)

	ExactlyOneOf []string
	AtLeastOneOf []string

	Elem interface{}
}

type ValueType int

const (
	TypeInvalid ValueType = iota
	TypeBool
	TypeInt
	TypeFloat
	TypeString
	TypeList
	TypeMap
	TypeSet
	typeObject
)

