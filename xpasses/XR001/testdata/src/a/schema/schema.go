package schema

type ResourceData struct{}

func (d *ResourceData) GetOkExists(key string) (interface{}, bool) {
	return nil, false
}
