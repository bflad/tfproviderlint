package schema

type Resource struct {
	Exists ExistsFunc
}

type ResourceData struct{}

type ExistsFunc func(*ResourceData, interface{}) (bool, error)
