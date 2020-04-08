package schema

type ReadFunc func(*ResourceData, interface{}) error

type Resource struct {
	Read   ReadFunc
	Schema map[string]*Schema
}

type ResourceData struct{}

type Schema struct{}
