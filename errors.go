package cloudns

import (
	"fmt"
	"strings"
)

// Constant errors which can be returned by cloudns-go when something goes wrong
const (
	ErrHTTPRequest         = constError("http request failed")
	ErrAPIInvocation       = constError("api invocation failed")
	ErrIllegalArgument     = constError("illegal argument provided")
	ErrInvalidOptions      = constError("invalid options provided")
	ErrMultipleCredentials = constError("more than one kind of credentials specified")
)

type constError string

func (err constError) wrap(inner error) error {
	return wrapError{outer: err, inner: inner}
}

func (err constError) Error() string {
	return string(err)
}

func (err constError) Is(target error) bool {
	errMsg := string(err)
	targetMsg := target.Error()
	return targetMsg == errMsg || strings.HasPrefix(targetMsg, errMsg+": ")
}

type wrapError struct {
	outer constError
	inner error
}

func (err wrapError) Error() string {
	if err.inner != nil {
		return fmt.Sprintf("%s: %v", err.outer.Error(), err.inner)
	}
	return err.outer.Error()
}

func (err wrapError) Is(target error) bool {
	return err.outer.Is(target)
}

func (err wrapError) Unwrap() error {
	return err.inner
}
