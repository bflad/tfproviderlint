package schema

type Schema struct {
	Computed    bool
	DefaultFunc func() (interface{}, error)
}
