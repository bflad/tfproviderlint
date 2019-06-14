package schema

type Schema struct {
	Computed     bool
	ValidateFunc func()
}
