package error

import (
	"github.com/bdlm/std/v2/caller"
)

// Caller is the interface implemented by error types that can expose
// runtime caller data.
type Caller interface {
	// Caller returns the associated Caller instance.
	Caller() caller.Caller
}

// Error defines a robust error stack interface.
type Error interface {
	// Error implements error.
	Error() string

	// Is tests to see if the test error matches any error in the stack via
	// equality comparison.
	Is(error) bool

	// Unwrap returns the wrapped error, if any, otherwise nil.
	Unwrap() error
}

// Wrapper is an interface implemented by error types that have wrapped
// a previous error.
type Wrapper interface {
	// Unwrap returns the next error in the error stack, if any, otherwise
	// nil.
	Unwrap() error
}
