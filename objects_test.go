package slim_test

import (
	"testing"

	"github.com/snple/slim"
	"github.com/snple/slim/require"
	"github.com/snple/slim/token"
)

func TestObject_TypeName(t *testing.T) {
	var o slim.Object = &slim.Int{}
	require.Equal(t, "int", o.TypeName())
	o = &slim.Float{}
	require.Equal(t, "float", o.TypeName())
	o = &slim.Char{}
	require.Equal(t, "char", o.TypeName())
	o = &slim.String{}
	require.Equal(t, "string", o.TypeName())
	o = &slim.Bool{}
	require.Equal(t, "bool", o.TypeName())
	o = &slim.Array{}
	require.Equal(t, "array", o.TypeName())
	o = &slim.Map{}
	require.Equal(t, "map", o.TypeName())
	o = &slim.ArrayIterator{}
	require.Equal(t, "array-iterator", o.TypeName())
	o = &slim.StringIterator{}
	require.Equal(t, "string-iterator", o.TypeName())
	o = &slim.MapIterator{}
	require.Equal(t, "map-iterator", o.TypeName())
	o = &slim.BuiltinFunction{Name: "fn"}
	require.Equal(t, "builtin-function:fn", o.TypeName())
	o = &slim.UserFunction{Name: "fn"}
	require.Equal(t, "user-function:fn", o.TypeName())
	o = &slim.CompiledFunction{}
	require.Equal(t, "compiled-function", o.TypeName())
	o = &slim.Undefined{}
	require.Equal(t, "undefined", o.TypeName())
	o = &slim.Error{}
	require.Equal(t, "error", o.TypeName())
	o = &slim.Bytes{}
	require.Equal(t, "bytes", o.TypeName())
}

func TestObject_IsFalsy(t *testing.T) {
	var o slim.Object = &slim.Int{Value: 0}
	require.True(t, o.IsFalsy())
	o = &slim.Int{Value: 1}
	require.False(t, o.IsFalsy())
	o = &slim.Float{Value: 0}
	require.False(t, o.IsFalsy())
	o = &slim.Float{Value: 1}
	require.False(t, o.IsFalsy())
	o = &slim.Char{Value: ' '}
	require.False(t, o.IsFalsy())
	o = &slim.Char{Value: 'T'}
	require.False(t, o.IsFalsy())
	o = &slim.String{Value: ""}
	require.True(t, o.IsFalsy())
	o = &slim.String{Value: " "}
	require.False(t, o.IsFalsy())
	o = &slim.Array{Value: nil}
	require.True(t, o.IsFalsy())
	o = &slim.Array{Value: []slim.Object{nil}} // nil is not valid but still count as 1 element
	require.False(t, o.IsFalsy())
	o = &slim.Map{Value: nil}
	require.True(t, o.IsFalsy())
	o = &slim.Map{Value: map[string]slim.Object{"a": nil}} // nil is not valid but still count as 1 element
	require.False(t, o.IsFalsy())
	o = &slim.StringIterator{}
	require.True(t, o.IsFalsy())
	o = &slim.ArrayIterator{}
	require.True(t, o.IsFalsy())
	o = &slim.MapIterator{}
	require.True(t, o.IsFalsy())
	o = &slim.BuiltinFunction{}
	require.False(t, o.IsFalsy())
	o = &slim.CompiledFunction{}
	require.False(t, o.IsFalsy())
	o = &slim.Undefined{}
	require.True(t, o.IsFalsy())
	o = &slim.Error{}
	require.True(t, o.IsFalsy())
	o = &slim.Bytes{}
	require.True(t, o.IsFalsy())
	o = &slim.Bytes{Value: []byte{1, 2}}
	require.False(t, o.IsFalsy())
}

func TestObject_String(t *testing.T) {
	var o slim.Object = &slim.Int{Value: 0}
	require.Equal(t, "0", o.String())
	o = &slim.Int{Value: 1}
	require.Equal(t, "1", o.String())
	o = &slim.Float{Value: 0}
	require.Equal(t, "0", o.String())
	o = &slim.Float{Value: 1}
	require.Equal(t, "1", o.String())
	o = &slim.Char{Value: ' '}
	require.Equal(t, " ", o.String())
	o = &slim.Char{Value: 'T'}
	require.Equal(t, "T", o.String())
	o = &slim.String{Value: ""}
	require.Equal(t, `""`, o.String())
	o = &slim.String{Value: " "}
	require.Equal(t, `" "`, o.String())
	o = &slim.Array{Value: nil}
	require.Equal(t, "[]", o.String())
	o = &slim.Map{Value: nil}
	require.Equal(t, "{}", o.String())
	o = &slim.Error{Value: nil}
	require.Equal(t, "error", o.String())
	o = &slim.Error{Value: &slim.String{Value: "error 1"}}
	require.Equal(t, `error: "error 1"`, o.String())
	o = &slim.StringIterator{}
	require.Equal(t, "<string-iterator>", o.String())
	o = &slim.ArrayIterator{}
	require.Equal(t, "<array-iterator>", o.String())
	o = &slim.MapIterator{}
	require.Equal(t, "<map-iterator>", o.String())
	o = &slim.Undefined{}
	require.Equal(t, "<undefined>", o.String())
	o = &slim.Bytes{}
	require.Equal(t, "", o.String())
	o = &slim.Bytes{Value: []byte("foo")}
	require.Equal(t, "foo", o.String())
}

func TestObject_BinaryOp(t *testing.T) {
	var o slim.Object = &slim.Char{}
	_, err := o.BinaryOp(token.Add, slim.UndefinedValue)
	require.Error(t, err)
	o = &slim.Bool{}
	_, err = o.BinaryOp(token.Add, slim.UndefinedValue)
	require.Error(t, err)
	o = &slim.Map{}
	_, err = o.BinaryOp(token.Add, slim.UndefinedValue)
	require.Error(t, err)
	o = &slim.ArrayIterator{}
	_, err = o.BinaryOp(token.Add, slim.UndefinedValue)
	require.Error(t, err)
	o = &slim.StringIterator{}
	_, err = o.BinaryOp(token.Add, slim.UndefinedValue)
	require.Error(t, err)
	o = &slim.MapIterator{}
	_, err = o.BinaryOp(token.Add, slim.UndefinedValue)
	require.Error(t, err)
	o = &slim.BuiltinFunction{}
	_, err = o.BinaryOp(token.Add, slim.UndefinedValue)
	require.Error(t, err)
	o = &slim.CompiledFunction{}
	_, err = o.BinaryOp(token.Add, slim.UndefinedValue)
	require.Error(t, err)
	o = &slim.Undefined{}
	_, err = o.BinaryOp(token.Add, slim.UndefinedValue)
	require.Error(t, err)
	o = &slim.Error{}
	_, err = o.BinaryOp(token.Add, slim.UndefinedValue)
	require.Error(t, err)
}

func TestArray_BinaryOp(t *testing.T) {
	testBinaryOp(t, &slim.Array{Value: nil}, token.Add,
		&slim.Array{Value: nil}, &slim.Array{Value: nil})
	testBinaryOp(t, &slim.Array{Value: nil}, token.Add,
		&slim.Array{Value: []slim.Object{}}, &slim.Array{Value: nil})
	testBinaryOp(t, &slim.Array{Value: []slim.Object{}}, token.Add,
		&slim.Array{Value: nil}, &slim.Array{Value: []slim.Object{}})
	testBinaryOp(t, &slim.Array{Value: []slim.Object{}}, token.Add,
		&slim.Array{Value: []slim.Object{}},
		&slim.Array{Value: []slim.Object{}})
	testBinaryOp(t, &slim.Array{Value: nil}, token.Add,
		&slim.Array{Value: []slim.Object{
			&slim.Int{Value: 1},
		}}, &slim.Array{Value: []slim.Object{
			&slim.Int{Value: 1},
		}})
	testBinaryOp(t, &slim.Array{Value: nil}, token.Add,
		&slim.Array{Value: []slim.Object{
			&slim.Int{Value: 1},
			&slim.Int{Value: 2},
			&slim.Int{Value: 3},
		}}, &slim.Array{Value: []slim.Object{
			&slim.Int{Value: 1},
			&slim.Int{Value: 2},
			&slim.Int{Value: 3},
		}})
	testBinaryOp(t, &slim.Array{Value: []slim.Object{
		&slim.Int{Value: 1},
		&slim.Int{Value: 2},
		&slim.Int{Value: 3},
	}}, token.Add, &slim.Array{Value: nil},
		&slim.Array{Value: []slim.Object{
			&slim.Int{Value: 1},
			&slim.Int{Value: 2},
			&slim.Int{Value: 3},
		}})
	testBinaryOp(t, &slim.Array{Value: []slim.Object{
		&slim.Int{Value: 1},
		&slim.Int{Value: 2},
		&slim.Int{Value: 3},
	}}, token.Add, &slim.Array{Value: []slim.Object{
		&slim.Int{Value: 4},
		&slim.Int{Value: 5},
		&slim.Int{Value: 6},
	}}, &slim.Array{Value: []slim.Object{
		&slim.Int{Value: 1},
		&slim.Int{Value: 2},
		&slim.Int{Value: 3},
		&slim.Int{Value: 4},
		&slim.Int{Value: 5},
		&slim.Int{Value: 6},
	}})
}

func TestError_Equals(t *testing.T) {
	err1 := &slim.Error{Value: &slim.String{Value: "some error"}}
	err2 := err1
	require.True(t, err1.Equals(err2))
	require.True(t, err2.Equals(err1))

	err2 = &slim.Error{Value: &slim.String{Value: "some error"}}
	require.False(t, err1.Equals(err2))
	require.False(t, err2.Equals(err1))
}

func TestFloat_BinaryOp(t *testing.T) {
	// float + float
	for l := float64(-2); l <= 2.1; l += 0.4 {
		for r := float64(-2); r <= 2.1; r += 0.4 {
			testBinaryOp(t, &slim.Float{Value: l}, token.Add,
				&slim.Float{Value: r}, &slim.Float{Value: l + r})
		}
	}

	// float - float
	for l := float64(-2); l <= 2.1; l += 0.4 {
		for r := float64(-2); r <= 2.1; r += 0.4 {
			testBinaryOp(t, &slim.Float{Value: l}, token.Sub,
				&slim.Float{Value: r}, &slim.Float{Value: l - r})
		}
	}

	// float * float
	for l := float64(-2); l <= 2.1; l += 0.4 {
		for r := float64(-2); r <= 2.1; r += 0.4 {
			testBinaryOp(t, &slim.Float{Value: l}, token.Mul,
				&slim.Float{Value: r}, &slim.Float{Value: l * r})
		}
	}

	// float / float
	for l := float64(-2); l <= 2.1; l += 0.4 {
		for r := float64(-2); r <= 2.1; r += 0.4 {
			if r != 0 {
				testBinaryOp(t, &slim.Float{Value: l}, token.Quo,
					&slim.Float{Value: r}, &slim.Float{Value: l / r})
			}
		}
	}

	// float < float
	for l := float64(-2); l <= 2.1; l += 0.4 {
		for r := float64(-2); r <= 2.1; r += 0.4 {
			testBinaryOp(t, &slim.Float{Value: l}, token.Less,
				&slim.Float{Value: r}, boolValue(l < r))
		}
	}

	// float > float
	for l := float64(-2); l <= 2.1; l += 0.4 {
		for r := float64(-2); r <= 2.1; r += 0.4 {
			testBinaryOp(t, &slim.Float{Value: l}, token.Greater,
				&slim.Float{Value: r}, boolValue(l > r))
		}
	}

	// float <= float
	for l := float64(-2); l <= 2.1; l += 0.4 {
		for r := float64(-2); r <= 2.1; r += 0.4 {
			testBinaryOp(t, &slim.Float{Value: l}, token.LessEq,
				&slim.Float{Value: r}, boolValue(l <= r))
		}
	}

	// float >= float
	for l := float64(-2); l <= 2.1; l += 0.4 {
		for r := float64(-2); r <= 2.1; r += 0.4 {
			testBinaryOp(t, &slim.Float{Value: l}, token.GreaterEq,
				&slim.Float{Value: r}, boolValue(l >= r))
		}
	}

	// float + int
	for l := float64(-2); l <= 2.1; l += 0.4 {
		for r := int64(-2); r <= 2; r++ {
			testBinaryOp(t, &slim.Float{Value: l}, token.Add,
				&slim.Int{Value: r}, &slim.Float{Value: l + float64(r)})
		}
	}

	// float - int
	for l := float64(-2); l <= 2.1; l += 0.4 {
		for r := int64(-2); r <= 2; r++ {
			testBinaryOp(t, &slim.Float{Value: l}, token.Sub,
				&slim.Int{Value: r}, &slim.Float{Value: l - float64(r)})
		}
	}

	// float * int
	for l := float64(-2); l <= 2.1; l += 0.4 {
		for r := int64(-2); r <= 2; r++ {
			testBinaryOp(t, &slim.Float{Value: l}, token.Mul,
				&slim.Int{Value: r}, &slim.Float{Value: l * float64(r)})
		}
	}

	// float / int
	for l := float64(-2); l <= 2.1; l += 0.4 {
		for r := int64(-2); r <= 2; r++ {
			if r != 0 {
				testBinaryOp(t, &slim.Float{Value: l}, token.Quo,
					&slim.Int{Value: r},
					&slim.Float{Value: l / float64(r)})
			}
		}
	}

	// float < int
	for l := float64(-2); l <= 2.1; l += 0.4 {
		for r := int64(-2); r <= 2; r++ {
			testBinaryOp(t, &slim.Float{Value: l}, token.Less,
				&slim.Int{Value: r}, boolValue(l < float64(r)))
		}
	}

	// float > int
	for l := float64(-2); l <= 2.1; l += 0.4 {
		for r := int64(-2); r <= 2; r++ {
			testBinaryOp(t, &slim.Float{Value: l}, token.Greater,
				&slim.Int{Value: r}, boolValue(l > float64(r)))
		}
	}

	// float <= int
	for l := float64(-2); l <= 2.1; l += 0.4 {
		for r := int64(-2); r <= 2; r++ {
			testBinaryOp(t, &slim.Float{Value: l}, token.LessEq,
				&slim.Int{Value: r}, boolValue(l <= float64(r)))
		}
	}

	// float >= int
	for l := float64(-2); l <= 2.1; l += 0.4 {
		for r := int64(-2); r <= 2; r++ {
			testBinaryOp(t, &slim.Float{Value: l}, token.GreaterEq,
				&slim.Int{Value: r}, boolValue(l >= float64(r)))
		}
	}
}

func TestInt_BinaryOp(t *testing.T) {
	// int + int
	for l := int64(-2); l <= 2; l++ {
		for r := int64(-2); r <= 2; r++ {
			testBinaryOp(t, &slim.Int{Value: l}, token.Add,
				&slim.Int{Value: r}, &slim.Int{Value: l + r})
		}
	}

	// int - int
	for l := int64(-2); l <= 2; l++ {
		for r := int64(-2); r <= 2; r++ {
			testBinaryOp(t, &slim.Int{Value: l}, token.Sub,
				&slim.Int{Value: r}, &slim.Int{Value: l - r})
		}
	}

	// int * int
	for l := int64(-2); l <= 2; l++ {
		for r := int64(-2); r <= 2; r++ {
			testBinaryOp(t, &slim.Int{Value: l}, token.Mul,
				&slim.Int{Value: r}, &slim.Int{Value: l * r})
		}
	}

	// int / int
	for l := int64(-2); l <= 2; l++ {
		for r := int64(-2); r <= 2; r++ {
			if r != 0 {
				testBinaryOp(t, &slim.Int{Value: l}, token.Quo,
					&slim.Int{Value: r}, &slim.Int{Value: l / r})
			}
		}
	}

	// int % int
	for l := int64(-4); l <= 4; l++ {
		for r := -int64(-4); r <= 4; r++ {
			if r == 0 {
				testBinaryOp(t, &slim.Int{Value: l}, token.Rem,
					&slim.Int{Value: r}, &slim.Int{Value: l % r})
			}
		}
	}

	// int & int
	testBinaryOp(t,
		&slim.Int{Value: 0}, token.And, &slim.Int{Value: 0},
		&slim.Int{Value: int64(0)})
	testBinaryOp(t,
		&slim.Int{Value: 1}, token.And, &slim.Int{Value: 0},
		&slim.Int{Value: int64(1) & int64(0)})
	testBinaryOp(t,
		&slim.Int{Value: 0}, token.And, &slim.Int{Value: 1},
		&slim.Int{Value: int64(0) & int64(1)})
	testBinaryOp(t,
		&slim.Int{Value: 1}, token.And, &slim.Int{Value: 1},
		&slim.Int{Value: int64(1)})
	testBinaryOp(t,
		&slim.Int{Value: 0}, token.And, &slim.Int{Value: int64(0xffffffff)},
		&slim.Int{Value: int64(0) & int64(0xffffffff)})
	testBinaryOp(t,
		&slim.Int{Value: 1}, token.And, &slim.Int{Value: int64(0xffffffff)},
		&slim.Int{Value: int64(1) & int64(0xffffffff)})
	testBinaryOp(t,
		&slim.Int{Value: int64(0xffffffff)}, token.And,
		&slim.Int{Value: int64(0xffffffff)},
		&slim.Int{Value: int64(0xffffffff)})
	testBinaryOp(t,
		&slim.Int{Value: 1984}, token.And,
		&slim.Int{Value: int64(0xffffffff)},
		&slim.Int{Value: int64(1984) & int64(0xffffffff)})
	testBinaryOp(t, &slim.Int{Value: -1984}, token.And,
		&slim.Int{Value: int64(0xffffffff)},
		&slim.Int{Value: int64(-1984) & int64(0xffffffff)})

	// int | int
	testBinaryOp(t,
		&slim.Int{Value: 0}, token.Or, &slim.Int{Value: 0},
		&slim.Int{Value: int64(0)})
	testBinaryOp(t,
		&slim.Int{Value: 1}, token.Or, &slim.Int{Value: 0},
		&slim.Int{Value: int64(1) | int64(0)})
	testBinaryOp(t,
		&slim.Int{Value: 0}, token.Or, &slim.Int{Value: 1},
		&slim.Int{Value: int64(0) | int64(1)})
	testBinaryOp(t,
		&slim.Int{Value: 1}, token.Or, &slim.Int{Value: 1},
		&slim.Int{Value: int64(1)})
	testBinaryOp(t,
		&slim.Int{Value: 0}, token.Or, &slim.Int{Value: int64(0xffffffff)},
		&slim.Int{Value: int64(0) | int64(0xffffffff)})
	testBinaryOp(t,
		&slim.Int{Value: 1}, token.Or, &slim.Int{Value: int64(0xffffffff)},
		&slim.Int{Value: int64(1) | int64(0xffffffff)})
	testBinaryOp(t,
		&slim.Int{Value: int64(0xffffffff)}, token.Or,
		&slim.Int{Value: int64(0xffffffff)},
		&slim.Int{Value: int64(0xffffffff)})
	testBinaryOp(t,
		&slim.Int{Value: 1984}, token.Or,
		&slim.Int{Value: int64(0xffffffff)},
		&slim.Int{Value: int64(1984) | int64(0xffffffff)})
	testBinaryOp(t,
		&slim.Int{Value: -1984}, token.Or,
		&slim.Int{Value: int64(0xffffffff)},
		&slim.Int{Value: int64(-1984) | int64(0xffffffff)})

	// int ^ int
	testBinaryOp(t,
		&slim.Int{Value: 0}, token.Xor, &slim.Int{Value: 0},
		&slim.Int{Value: int64(0)})
	testBinaryOp(t,
		&slim.Int{Value: 1}, token.Xor, &slim.Int{Value: 0},
		&slim.Int{Value: int64(1) ^ int64(0)})
	testBinaryOp(t,
		&slim.Int{Value: 0}, token.Xor, &slim.Int{Value: 1},
		&slim.Int{Value: int64(0) ^ int64(1)})
	testBinaryOp(t,
		&slim.Int{Value: 1}, token.Xor, &slim.Int{Value: 1},
		&slim.Int{Value: int64(0)})
	testBinaryOp(t,
		&slim.Int{Value: 0}, token.Xor, &slim.Int{Value: int64(0xffffffff)},
		&slim.Int{Value: int64(0) ^ int64(0xffffffff)})
	testBinaryOp(t,
		&slim.Int{Value: 1}, token.Xor, &slim.Int{Value: int64(0xffffffff)},
		&slim.Int{Value: int64(1) ^ int64(0xffffffff)})
	testBinaryOp(t,
		&slim.Int{Value: int64(0xffffffff)}, token.Xor,
		&slim.Int{Value: int64(0xffffffff)},
		&slim.Int{Value: int64(0)})
	testBinaryOp(t,
		&slim.Int{Value: 1984}, token.Xor,
		&slim.Int{Value: int64(0xffffffff)},
		&slim.Int{Value: int64(1984) ^ int64(0xffffffff)})
	testBinaryOp(t,
		&slim.Int{Value: -1984}, token.Xor,
		&slim.Int{Value: int64(0xffffffff)},
		&slim.Int{Value: int64(-1984) ^ int64(0xffffffff)})

	// int &^ int
	testBinaryOp(t,
		&slim.Int{Value: 0}, token.AndNot, &slim.Int{Value: 0},
		&slim.Int{Value: int64(0)})
	testBinaryOp(t,
		&slim.Int{Value: 1}, token.AndNot, &slim.Int{Value: 0},
		&slim.Int{Value: int64(1) &^ int64(0)})
	testBinaryOp(t,
		&slim.Int{Value: 0}, token.AndNot,
		&slim.Int{Value: 1}, &slim.Int{Value: int64(0) &^ int64(1)})
	testBinaryOp(t,
		&slim.Int{Value: 1}, token.AndNot, &slim.Int{Value: 1},
		&slim.Int{Value: int64(0)})
	testBinaryOp(t,
		&slim.Int{Value: 0}, token.AndNot,
		&slim.Int{Value: int64(0xffffffff)},
		&slim.Int{Value: int64(0) &^ int64(0xffffffff)})
	testBinaryOp(t,
		&slim.Int{Value: 1}, token.AndNot,
		&slim.Int{Value: int64(0xffffffff)},
		&slim.Int{Value: int64(1) &^ int64(0xffffffff)})
	testBinaryOp(t,
		&slim.Int{Value: int64(0xffffffff)}, token.AndNot,
		&slim.Int{Value: int64(0xffffffff)},
		&slim.Int{Value: int64(0)})
	testBinaryOp(t,
		&slim.Int{Value: 1984}, token.AndNot,
		&slim.Int{Value: int64(0xffffffff)},
		&slim.Int{Value: int64(1984) &^ int64(0xffffffff)})
	testBinaryOp(t,
		&slim.Int{Value: -1984}, token.AndNot,
		&slim.Int{Value: int64(0xffffffff)},
		&slim.Int{Value: int64(-1984) &^ int64(0xffffffff)})

	// int << int
	for s := int64(0); s < 64; s++ {
		testBinaryOp(t,
			&slim.Int{Value: 0}, token.Shl, &slim.Int{Value: s},
			&slim.Int{Value: int64(0) << uint(s)})
		testBinaryOp(t,
			&slim.Int{Value: 1}, token.Shl, &slim.Int{Value: s},
			&slim.Int{Value: int64(1) << uint(s)})
		testBinaryOp(t,
			&slim.Int{Value: 2}, token.Shl, &slim.Int{Value: s},
			&slim.Int{Value: int64(2) << uint(s)})
		testBinaryOp(t,
			&slim.Int{Value: -1}, token.Shl, &slim.Int{Value: s},
			&slim.Int{Value: int64(-1) << uint(s)})
		testBinaryOp(t,
			&slim.Int{Value: -2}, token.Shl, &slim.Int{Value: s},
			&slim.Int{Value: int64(-2) << uint(s)})
		testBinaryOp(t,
			&slim.Int{Value: int64(0xffffffff)}, token.Shl,
			&slim.Int{Value: s},
			&slim.Int{Value: int64(0xffffffff) << uint(s)})
	}

	// int >> int
	for s := int64(0); s < 64; s++ {
		testBinaryOp(t,
			&slim.Int{Value: 0}, token.Shr, &slim.Int{Value: s},
			&slim.Int{Value: int64(0) >> uint(s)})
		testBinaryOp(t,
			&slim.Int{Value: 1}, token.Shr, &slim.Int{Value: s},
			&slim.Int{Value: int64(1) >> uint(s)})
		testBinaryOp(t,
			&slim.Int{Value: 2}, token.Shr, &slim.Int{Value: s},
			&slim.Int{Value: int64(2) >> uint(s)})
		testBinaryOp(t,
			&slim.Int{Value: -1}, token.Shr, &slim.Int{Value: s},
			&slim.Int{Value: int64(-1) >> uint(s)})
		testBinaryOp(t,
			&slim.Int{Value: -2}, token.Shr, &slim.Int{Value: s},
			&slim.Int{Value: int64(-2) >> uint(s)})
		testBinaryOp(t,
			&slim.Int{Value: int64(0xffffffff)}, token.Shr,
			&slim.Int{Value: s},
			&slim.Int{Value: int64(0xffffffff) >> uint(s)})
	}

	// int < int
	for l := int64(-2); l <= 2; l++ {
		for r := int64(-2); r <= 2; r++ {
			testBinaryOp(t, &slim.Int{Value: l}, token.Less,
				&slim.Int{Value: r}, boolValue(l < r))
		}
	}

	// int > int
	for l := int64(-2); l <= 2; l++ {
		for r := int64(-2); r <= 2; r++ {
			testBinaryOp(t, &slim.Int{Value: l}, token.Greater,
				&slim.Int{Value: r}, boolValue(l > r))
		}
	}

	// int <= int
	for l := int64(-2); l <= 2; l++ {
		for r := int64(-2); r <= 2; r++ {
			testBinaryOp(t, &slim.Int{Value: l}, token.LessEq,
				&slim.Int{Value: r}, boolValue(l <= r))
		}
	}

	// int >= int
	for l := int64(-2); l <= 2; l++ {
		for r := int64(-2); r <= 2; r++ {
			testBinaryOp(t, &slim.Int{Value: l}, token.GreaterEq,
				&slim.Int{Value: r}, boolValue(l >= r))
		}
	}

	// int + float
	for l := int64(-2); l <= 2; l++ {
		for r := float64(-2); r <= 2.1; r += 0.5 {
			testBinaryOp(t, &slim.Int{Value: l}, token.Add,
				&slim.Float{Value: r},
				&slim.Float{Value: float64(l) + r})
		}
	}

	// int - float
	for l := int64(-2); l <= 2; l++ {
		for r := float64(-2); r <= 2.1; r += 0.5 {
			testBinaryOp(t, &slim.Int{Value: l}, token.Sub,
				&slim.Float{Value: r},
				&slim.Float{Value: float64(l) - r})
		}
	}

	// int * float
	for l := int64(-2); l <= 2; l++ {
		for r := float64(-2); r <= 2.1; r += 0.5 {
			testBinaryOp(t, &slim.Int{Value: l}, token.Mul,
				&slim.Float{Value: r},
				&slim.Float{Value: float64(l) * r})
		}
	}

	// int / float
	for l := int64(-2); l <= 2; l++ {
		for r := float64(-2); r <= 2.1; r += 0.5 {
			if r != 0 {
				testBinaryOp(t, &slim.Int{Value: l}, token.Quo,
					&slim.Float{Value: r},
					&slim.Float{Value: float64(l) / r})
			}
		}
	}

	// int < float
	for l := int64(-2); l <= 2; l++ {
		for r := float64(-2); r <= 2.1; r += 0.5 {
			testBinaryOp(t, &slim.Int{Value: l}, token.Less,
				&slim.Float{Value: r}, boolValue(float64(l) < r))
		}
	}

	// int > float
	for l := int64(-2); l <= 2; l++ {
		for r := float64(-2); r <= 2.1; r += 0.5 {
			testBinaryOp(t, &slim.Int{Value: l}, token.Greater,
				&slim.Float{Value: r}, boolValue(float64(l) > r))
		}
	}

	// int <= float
	for l := int64(-2); l <= 2; l++ {
		for r := float64(-2); r <= 2.1; r += 0.5 {
			testBinaryOp(t, &slim.Int{Value: l}, token.LessEq,
				&slim.Float{Value: r}, boolValue(float64(l) <= r))
		}
	}

	// int >= float
	for l := int64(-2); l <= 2; l++ {
		for r := float64(-2); r <= 2.1; r += 0.5 {
			testBinaryOp(t, &slim.Int{Value: l}, token.GreaterEq,
				&slim.Float{Value: r}, boolValue(float64(l) >= r))
		}
	}
}

func TestMap_Index(t *testing.T) {
	m := &slim.Map{Value: make(map[string]slim.Object)}
	k := &slim.Int{Value: 1}
	v := &slim.String{Value: "abcdef"}
	err := m.IndexSet(k, v)

	require.NoError(t, err)

	res, err := m.IndexGet(k)
	require.NoError(t, err)
	require.Equal(t, v, res)
}

func TestString_BinaryOp(t *testing.T) {
	lstr := "abcde"
	rstr := "01234"
	for l := 0; l < len(lstr); l++ {
		for r := 0; r < len(rstr); r++ {
			ls := lstr[l:]
			rs := rstr[r:]
			testBinaryOp(t, &slim.String{Value: ls}, token.Add,
				&slim.String{Value: rs},
				&slim.String{Value: ls + rs})

			rc := []rune(rstr)[r]
			testBinaryOp(t, &slim.String{Value: ls}, token.Add,
				&slim.Char{Value: rc},
				&slim.String{Value: ls + string(rc)})
		}
	}
}

func testBinaryOp(
	t *testing.T,
	lhs slim.Object,
	op token.Token,
	rhs slim.Object,
	expected slim.Object,
) {
	t.Helper()
	actual, err := lhs.BinaryOp(op, rhs)
	require.NoError(t, err)
	require.Equal(t, expected, actual)
}

func boolValue(b bool) slim.Object {
	if b {
		return slim.TrueValue
	}
	return slim.FalseValue
}
