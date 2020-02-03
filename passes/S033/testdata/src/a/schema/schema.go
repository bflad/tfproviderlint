package schema

type Schema struct {
	Computed  bool
	StateFunc func(interface{}) string
}
