package steamweb

import (
	"errors"
)

var (
	// ErrRequestNoInterface is returned when trying to execute a `Request` without an `Interface` set.
	ErrRequestNoInterface = errors.New("steamweb: no Interface set in Request")
	// ErrRequestNoMethod is returned when trying to execute a `Request` without an `Method` set.
	ErrRequestNoMethod = errors.New("steamweb: no Method set in Request")
)
