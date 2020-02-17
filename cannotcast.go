package cast

import (
	"fmt"
)

// The CannotCastComplainer interface identifies an error due to a situation
// where converting from one type to another could not be done.
//
// Example:
//
//	s, err := cast.String(v)
//	if nil != err {
//		switch {
//		case cast.CannotCastComplainer:
//			//@TODO
//		default:
//			//@TODO
//		}
//	}
type CannotCastComplainer interface {
	error
	CannotCastComplainer()
}

type internalCannotCastComplainer struct{
	expectedType string
	actualType string
	reason string
}

func (r internalCannotCastComplainer) Error() string {
	reason := r.reason
	if reason == "" {
		reason = "incompatible types"
	}
	return fmt.Sprintf("cannot cast %q to %q [reason: %s]", r.actualType, r.expectedType, reason)
}

func (internalCannotCastComplainer) CannotCastComplainer() {
	// Nothing here.
}
