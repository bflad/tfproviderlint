package schema

import "time"

type Resource struct {
	Create CreateFunc
	Timeouts *ResourceTimeout
}

type ResourceTimeout struct {
	Create *time.Duration
}

type ResourceData struct{}

type CreateFunc func(*ResourceData, interface{}) error
