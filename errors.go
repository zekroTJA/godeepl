package godeepl

import "fmt"

type ResponseError struct {
	Code int
}

func (err ResponseError) Error() string {
	return fmt.Sprintf("Request Error %d", err.Code)
}
