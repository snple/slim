package stdlib

import (
	"fmt"

	"github.com/snple/slim"
)

var fmtModule = map[string]slim.Object{
	"print":   &slim.UserFunction{Name: "print", Value: fmtPrint},
	"printf":  &slim.UserFunction{Name: "printf", Value: fmtPrintf},
	"println": &slim.UserFunction{Name: "println", Value: fmtPrintln},
	"sprintf": &slim.UserFunction{Name: "sprintf", Value: fmtSprintf},
}

func fmtPrint(args ...slim.Object) (ret slim.Object, err error) {
	printArgs, err := getPrintArgs(args...)
	if err != nil {
		return nil, err
	}
	_, _ = fmt.Print(printArgs...)
	return nil, nil
}

func fmtPrintf(args ...slim.Object) (ret slim.Object, err error) {
	numArgs := len(args)
	if numArgs == 0 {
		return nil, slim.ErrWrongNumArguments
	}

	format, ok := args[0].(*slim.String)
	if !ok {
		return nil, slim.ErrInvalidArgumentType{
			Name:     "format",
			Expected: "string",
			Found:    args[0].TypeName(),
		}
	}
	if numArgs == 1 {
		fmt.Print(format)
		return nil, nil
	}

	s, err := slim.Format(format.Value, args[1:]...)
	if err != nil {
		return nil, err
	}
	fmt.Print(s)
	return nil, nil
}

func fmtPrintln(args ...slim.Object) (ret slim.Object, err error) {
	printArgs, err := getPrintArgs(args...)
	if err != nil {
		return nil, err
	}
	printArgs = append(printArgs, "\n")
	_, _ = fmt.Print(printArgs...)
	return nil, nil
}

func fmtSprintf(args ...slim.Object) (ret slim.Object, err error) {
	numArgs := len(args)
	if numArgs == 0 {
		return nil, slim.ErrWrongNumArguments
	}

	format, ok := args[0].(*slim.String)
	if !ok {
		return nil, slim.ErrInvalidArgumentType{
			Name:     "format",
			Expected: "string",
			Found:    args[0].TypeName(),
		}
	}
	if numArgs == 1 {
		// okay to return 'format' directly as String is immutable
		return format, nil
	}
	s, err := slim.Format(format.Value, args[1:]...)
	if err != nil {
		return nil, err
	}
	return &slim.String{Value: s}, nil
}

func getPrintArgs(args ...slim.Object) ([]interface{}, error) {
	var printArgs []interface{}
	l := 0
	for _, arg := range args {
		s, _ := slim.ToString(arg)
		slen := len(s)
		// make sure length does not exceed the limit
		if l+slen > slim.MaxStringLen {
			return nil, slim.ErrStringLimit
		}
		l += slen
		printArgs = append(printArgs, s)
	}
	return printArgs, nil
}
