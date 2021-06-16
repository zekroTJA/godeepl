package godeepl

import "fmt"

// ResponseError is returned when a request returns an
// erroneous response code.
type ResponseError struct {
	Code int
}

func (err ResponseError) Error() string {
	return fmt.Sprintf("Request Error %d", err.Code)
}
