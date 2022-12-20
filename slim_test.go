package slim_test

import (
	"strings"
	"testing"
	"time"

	"github.com/snple/slim"
	"github.com/snple/slim/parser"
	"github.com/snple/slim/require"
)

func TestInstructions_String(t *testing.T) {
	assertInstructionString(t,
		[][]byte{
			slim.MakeInstruction(parser.OpConstant, 1),
			slim.MakeInstruction(parser.OpConstant, 2),
			slim.MakeInstruction(parser.OpConstant, 65535),
		},
		`0000 CONST   1    
0003 CONST   2    
0006 CONST   65535`)

	assertInstructionString(t,
		[][]byte{
			slim.MakeInstruction(parser.OpBinaryOp, 11),
			slim.MakeInstruction(parser.OpConstant, 2),
			slim.MakeInstruction(parser.OpConstant, 65535),
		},
		`0000 BINARYOP 11   
0002 CONST   2    
0005 CONST   65535`)

	assertInstructionString(t,
		[][]byte{
			slim.MakeInstruction(parser.OpBinaryOp, 11),
			slim.MakeInstruction(parser.OpGetLocal, 1),
			slim.MakeInstruction(parser.OpConstant, 2),
			slim.MakeInstruction(parser.OpConstant, 65535),
		},
		`0000 BINARYOP 11   
0002 GETL    1    
0004 CONST   2    
0007 CONST   65535`)
}

func TestMakeInstruction(t *testing.T) {
	makeInstruction(t, []byte{parser.OpConstant, 0, 0},
		parser.OpConstant, 0)
	makeInstruction(t, []byte{parser.OpConstant, 0, 1},
		parser.OpConstant, 1)
	makeInstruction(t, []byte{parser.OpConstant, 255, 254},
		parser.OpConstant, 65534)
	makeInstruction(t, []byte{parser.OpPop}, parser.OpPop)
	makeInstruction(t, []byte{parser.OpTrue}, parser.OpTrue)
	makeInstruction(t, []byte{parser.OpFalse}, parser.OpFalse)
}

func TestNumObjects(t *testing.T) {
	testCountObjects(t, &slim.Array{}, 1)
	testCountObjects(t, &slim.Array{Value: []slim.Object{
		&slim.Int{Value: 1},
		&slim.Int{Value: 2},
		&slim.Array{Value: []slim.Object{
			&slim.Int{Value: 3},
			&slim.Int{Value: 4},
			&slim.Int{Value: 5},
		}},
	}}, 7)
	testCountObjects(t, slim.TrueValue, 1)
	testCountObjects(t, slim.FalseValue, 1)
	testCountObjects(t, &slim.BuiltinFunction{}, 1)
	testCountObjects(t, &slim.Bytes{Value: []byte("foobar")}, 1)
	testCountObjects(t, &slim.Char{Value: '가'}, 1)
	testCountObjects(t, &slim.CompiledFunction{}, 1)
	testCountObjects(t, &slim.Error{Value: &slim.Int{Value: 5}}, 2)
	testCountObjects(t, &slim.Float{Value: 19.84}, 1)
	testCountObjects(t, &slim.ImmutableArray{Value: []slim.Object{
		&slim.Int{Value: 1},
		&slim.Int{Value: 2},
		&slim.ImmutableArray{Value: []slim.Object{
			&slim.Int{Value: 3},
			&slim.Int{Value: 4},
			&slim.Int{Value: 5},
		}},
	}}, 7)
	testCountObjects(t, &slim.ImmutableMap{
		Value: map[string]slim.Object{
			"k1": &slim.Int{Value: 1},
			"k2": &slim.Int{Value: 2},
			"k3": &slim.Array{Value: []slim.Object{
				&slim.Int{Value: 3},
				&slim.Int{Value: 4},
				&slim.Int{Value: 5},
			}},
		}}, 7)
	testCountObjects(t, &slim.Int{Value: 1984}, 1)
	testCountObjects(t, &slim.Map{Value: map[string]slim.Object{
		"k1": &slim.Int{Value: 1},
		"k2": &slim.Int{Value: 2},
		"k3": &slim.Array{Value: []slim.Object{
			&slim.Int{Value: 3},
			&slim.Int{Value: 4},
			&slim.Int{Value: 5},
		}},
	}}, 7)
	testCountObjects(t, &slim.String{Value: "foo bar"}, 1)
	testCountObjects(t, &slim.Time{Value: time.Now()}, 1)
	testCountObjects(t, slim.UndefinedValue, 1)
}

func testCountObjects(t *testing.T, o slim.Object, expected int) {
	require.Equal(t, expected, slim.CountObjects(o))
}

func assertInstructionString(
	t *testing.T,
	instructions [][]byte,
	expected string,
) {
	concatted := make([]byte, 0)
	for _, e := range instructions {
		concatted = append(concatted, e...)
	}
	require.Equal(t, expected, strings.Join(
		slim.FormatInstructions(concatted, 0), "\n"))
}

func makeInstruction(
	t *testing.T,
	expected []byte,
	opcode parser.Opcode,
	operands ...int,
) {
	inst := slim.MakeInstruction(opcode, operands...)
	require.Equal(t, expected, inst)
}
