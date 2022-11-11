package errors

import (
	"fmt"
)

type ResourceNotFoundError struct {
	Resource string
	ID       string
}

func (e *ResourceNotFoundError) Error() string {
	return fmt.Sprintf("Resource %v with ID:%v not found", e.Resource, e.ID)
}

type RequestError struct {
	StatusCode int
	Err        error
}
