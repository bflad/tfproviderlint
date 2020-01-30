package schema

type Resource struct {
	Create CreateFunc
}

type ResourceData struct{}

type CreateFunc func(*ResourceData, interface{}) error
