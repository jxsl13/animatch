package common

import "fmt"

// WrapErr returns a wrapper function for the provided error
// the returned function may take any argument as exra information in order to
// specify the error more concretely
func WrapErr(err error) func(interface{}) error {
	return func(info interface{}) error {
		return fmt.Errorf("%w: %v", err, info)
	}
}

type Error string

func (e Error) Error() string {
	return string(e)
}
