package schema

type ResourceData struct{}

func (d *ResourceData) HasChanges(keys ...string) bool {
	return false
}
