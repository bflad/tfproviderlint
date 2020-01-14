package schema

type ResourceData struct{}

func (d *ResourceData) HasChange(key string) bool {
	return false
}
