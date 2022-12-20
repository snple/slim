package stdlib

import "github.com/snple/slim"

func wrapError(err error) slim.Object {
	if err == nil {
		return slim.TrueValue
	}
	return &slim.Error{Value: &slim.String{Value: err.Error()}}
}
