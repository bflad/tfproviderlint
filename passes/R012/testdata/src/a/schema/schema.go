package schema

type CreateFunc func(*ResourceData, interface{}) error

type CustomizeDiffFunc func(*ResourceDiff, interface{}) error

type ReadFunc func(*ResourceData, interface{}) error

type Resource struct {
	Create        CreateFunc
	CustomizeDiff CustomizeDiffFunc
	Read          ReadFunc
}

type ResourceData struct{}

type ResourceDiff struct{}
