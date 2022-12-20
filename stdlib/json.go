package stdlib

import (
	"bytes"
	gojson "encoding/json"

	"github.com/snple/slim"
	"github.com/snple/slim/stdlib/json"
)

var jsonModule = map[string]slim.Object{
	"decode": &slim.UserFunction{
		Name:  "decode",
		Value: jsonDecode,
	},
	"encode": &slim.UserFunction{
		Name:  "encode",
		Value: jsonEncode,
	},
	"indent": &slim.UserFunction{
		Name:  "encode",
		Value: jsonIndent,
	},
	"html_escape": &slim.UserFunction{
		Name:  "html_escape",
		Value: jsonHTMLEscape,
	},
}

func jsonDecode(args ...slim.Object) (ret slim.Object, err error) {
	if len(args) != 1 {
		return nil, slim.ErrWrongNumArguments
	}

	switch o := args[0].(type) {
	case *slim.Bytes:
		v, err := json.Decode(o.Value)
		if err != nil {
			return &slim.Error{
				Value: &slim.String{Value: err.Error()},
			}, nil
		}
		return v, nil
	case *slim.String:
		v, err := json.Decode([]byte(o.Value))
		if err != nil {
			return &slim.Error{
				Value: &slim.String{Value: err.Error()},
			}, nil
		}
		return v, nil
	default:
		return nil, slim.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "bytes/string",
			Found:    args[0].TypeName(),
		}
	}
}

func jsonEncode(args ...slim.Object) (ret slim.Object, err error) {
	if len(args) != 1 {
		return nil, slim.ErrWrongNumArguments
	}

	b, err := json.Encode(args[0])
	if err != nil {
		return &slim.Error{Value: &slim.String{Value: err.Error()}}, nil
	}

	return &slim.Bytes{Value: b}, nil
}

func jsonIndent(args ...slim.Object) (ret slim.Object, err error) {
	if len(args) != 3 {
		return nil, slim.ErrWrongNumArguments
	}

	prefix, ok := slim.ToString(args[1])
	if !ok {
		return nil, slim.ErrInvalidArgumentType{
			Name:     "prefix",
			Expected: "string(compatible)",
			Found:    args[1].TypeName(),
		}
	}

	indent, ok := slim.ToString(args[2])
	if !ok {
		return nil, slim.ErrInvalidArgumentType{
			Name:     "indent",
			Expected: "string(compatible)",
			Found:    args[2].TypeName(),
		}
	}

	switch o := args[0].(type) {
	case *slim.Bytes:
		var dst bytes.Buffer
		err := gojson.Indent(&dst, o.Value, prefix, indent)
		if err != nil {
			return &slim.Error{
				Value: &slim.String{Value: err.Error()},
			}, nil
		}
		return &slim.Bytes{Value: dst.Bytes()}, nil
	case *slim.String:
		var dst bytes.Buffer
		err := gojson.Indent(&dst, []byte(o.Value), prefix, indent)
		if err != nil {
			return &slim.Error{
				Value: &slim.String{Value: err.Error()},
			}, nil
		}
		return &slim.Bytes{Value: dst.Bytes()}, nil
	default:
		return nil, slim.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "bytes/string",
			Found:    args[0].TypeName(),
		}
	}
}

func jsonHTMLEscape(args ...slim.Object) (ret slim.Object, err error) {
	if len(args) != 1 {
		return nil, slim.ErrWrongNumArguments
	}

	switch o := args[0].(type) {
	case *slim.Bytes:
		var dst bytes.Buffer
		gojson.HTMLEscape(&dst, o.Value)
		return &slim.Bytes{Value: dst.Bytes()}, nil
	case *slim.String:
		var dst bytes.Buffer
		gojson.HTMLEscape(&dst, []byte(o.Value))
		return &slim.Bytes{Value: dst.Bytes()}, nil
	default:
		return nil, slim.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "bytes/string",
			Found:    args[0].TypeName(),
		}
	}
}
