package slim_test

import (
	"bytes"
	"testing"
	"time"

	"github.com/snple/slim"
	"github.com/snple/slim/parser"
	"github.com/snple/slim/require"
)

type srcfile struct {
	name string
	size int
}

func TestBytecode(t *testing.T) {
	testBytecodeSerialization(t, bytecode(concatInsts(), objectsArray()))

	testBytecodeSerialization(t, bytecode(
		concatInsts(), objectsArray(
			&slim.Char{Value: 'y'},
			&slim.Float{Value: 93.11},
			compiledFunction(1, 0,
				slim.MakeInstruction(parser.OpConstant, 3),
				slim.MakeInstruction(parser.OpSetLocal, 0),
				slim.MakeInstruction(parser.OpGetGlobal, 0),
				slim.MakeInstruction(parser.OpGetFree, 0)),
			&slim.Float{Value: 39.2},
			&slim.Int{Value: 192},
			&slim.String{Value: "bar"})))

	testBytecodeSerialization(t, bytecodeFileSet(
		concatInsts(
			slim.MakeInstruction(parser.OpConstant, 0),
			slim.MakeInstruction(parser.OpSetGlobal, 0),
			slim.MakeInstruction(parser.OpConstant, 6),
			slim.MakeInstruction(parser.OpPop)),
		objectsArray(
			&slim.Int{Value: 55},
			&slim.Int{Value: 66},
			&slim.Int{Value: 77},
			&slim.Int{Value: 88},
			&slim.ImmutableMap{
				Value: map[string]slim.Object{
					"array": &slim.ImmutableArray{
						Value: []slim.Object{
							&slim.Int{Value: 1},
							&slim.Int{Value: 2},
							&slim.Int{Value: 3},
							slim.TrueValue,
							slim.FalseValue,
							slim.UndefinedValue,
						},
					},
					"true":  slim.TrueValue,
					"false": slim.FalseValue,
					"bytes": &slim.Bytes{Value: make([]byte, 16)},
					"char":  &slim.Char{Value: 'Y'},
					"error": &slim.Error{Value: &slim.String{
						Value: "some error",
					}},
					"float": &slim.Float{Value: -19.84},
					"immutable_array": &slim.ImmutableArray{
						Value: []slim.Object{
							&slim.Int{Value: 1},
							&slim.Int{Value: 2},
							&slim.Int{Value: 3},
							slim.TrueValue,
							slim.FalseValue,
							slim.UndefinedValue,
						},
					},
					"immutable_map": &slim.ImmutableMap{
						Value: map[string]slim.Object{
							"a": &slim.Int{Value: 1},
							"b": &slim.Int{Value: 2},
							"c": &slim.Int{Value: 3},
							"d": slim.TrueValue,
							"e": slim.FalseValue,
							"f": slim.UndefinedValue,
						},
					},
					"int": &slim.Int{Value: 91},
					"map": &slim.Map{
						Value: map[string]slim.Object{
							"a": &slim.Int{Value: 1},
							"b": &slim.Int{Value: 2},
							"c": &slim.Int{Value: 3},
							"d": slim.TrueValue,
							"e": slim.FalseValue,
							"f": slim.UndefinedValue,
						},
					},
					"string":    &slim.String{Value: "foo bar"},
					"time":      &slim.Time{Value: time.Now()},
					"undefined": slim.UndefinedValue,
				},
			},
			compiledFunction(1, 0,
				slim.MakeInstruction(parser.OpConstant, 3),
				slim.MakeInstruction(parser.OpSetLocal, 0),
				slim.MakeInstruction(parser.OpGetGlobal, 0),
				slim.MakeInstruction(parser.OpGetFree, 0),
				slim.MakeInstruction(parser.OpBinaryOp, 11),
				slim.MakeInstruction(parser.OpGetFree, 1),
				slim.MakeInstruction(parser.OpBinaryOp, 11),
				slim.MakeInstruction(parser.OpGetLocal, 0),
				slim.MakeInstruction(parser.OpBinaryOp, 11),
				slim.MakeInstruction(parser.OpReturn, 1)),
			compiledFunction(1, 0,
				slim.MakeInstruction(parser.OpConstant, 2),
				slim.MakeInstruction(parser.OpSetLocal, 0),
				slim.MakeInstruction(parser.OpGetFree, 0),
				slim.MakeInstruction(parser.OpGetLocal, 0),
				slim.MakeInstruction(parser.OpClosure, 4, 2),
				slim.MakeInstruction(parser.OpReturn, 1)),
			compiledFunction(1, 0,
				slim.MakeInstruction(parser.OpConstant, 1),
				slim.MakeInstruction(parser.OpSetLocal, 0),
				slim.MakeInstruction(parser.OpGetLocal, 0),
				slim.MakeInstruction(parser.OpClosure, 5, 1),
				slim.MakeInstruction(parser.OpReturn, 1))),
		fileSet(srcfile{name: "file1", size: 100},
			srcfile{name: "file2", size: 200})))
}

func TestBytecode_RemoveDuplicates(t *testing.T) {
	testBytecodeRemoveDuplicates(t,
		bytecode(
			concatInsts(), objectsArray(
				&slim.Char{Value: 'y'},
				&slim.Float{Value: 93.11},
				compiledFunction(1, 0,
					slim.MakeInstruction(parser.OpConstant, 3),
					slim.MakeInstruction(parser.OpSetLocal, 0),
					slim.MakeInstruction(parser.OpGetGlobal, 0),
					slim.MakeInstruction(parser.OpGetFree, 0)),
				&slim.Float{Value: 39.2},
				&slim.Int{Value: 192},
				&slim.String{Value: "bar"})),
		bytecode(
			concatInsts(), objectsArray(
				&slim.Char{Value: 'y'},
				&slim.Float{Value: 93.11},
				compiledFunction(1, 0,
					slim.MakeInstruction(parser.OpConstant, 3),
					slim.MakeInstruction(parser.OpSetLocal, 0),
					slim.MakeInstruction(parser.OpGetGlobal, 0),
					slim.MakeInstruction(parser.OpGetFree, 0)),
				&slim.Float{Value: 39.2},
				&slim.Int{Value: 192},
				&slim.String{Value: "bar"})))

	testBytecodeRemoveDuplicates(t,
		bytecode(
			concatInsts(
				slim.MakeInstruction(parser.OpConstant, 0),
				slim.MakeInstruction(parser.OpConstant, 1),
				slim.MakeInstruction(parser.OpConstant, 2),
				slim.MakeInstruction(parser.OpConstant, 3),
				slim.MakeInstruction(parser.OpConstant, 4),
				slim.MakeInstruction(parser.OpConstant, 5),
				slim.MakeInstruction(parser.OpConstant, 6),
				slim.MakeInstruction(parser.OpConstant, 7),
				slim.MakeInstruction(parser.OpConstant, 8),
				slim.MakeInstruction(parser.OpClosure, 4, 1)),
			objectsArray(
				&slim.Int{Value: 1},
				&slim.Float{Value: 2.0},
				&slim.Char{Value: '3'},
				&slim.String{Value: "four"},
				compiledFunction(1, 0,
					slim.MakeInstruction(parser.OpConstant, 3),
					slim.MakeInstruction(parser.OpConstant, 7),
					slim.MakeInstruction(parser.OpSetLocal, 0),
					slim.MakeInstruction(parser.OpGetGlobal, 0),
					slim.MakeInstruction(parser.OpGetFree, 0)),
				&slim.Int{Value: 1},
				&slim.Float{Value: 2.0},
				&slim.Char{Value: '3'},
				&slim.String{Value: "four"})),
		bytecode(
			concatInsts(
				slim.MakeInstruction(parser.OpConstant, 0),
				slim.MakeInstruction(parser.OpConstant, 1),
				slim.MakeInstruction(parser.OpConstant, 2),
				slim.MakeInstruction(parser.OpConstant, 3),
				slim.MakeInstruction(parser.OpConstant, 4),
				slim.MakeInstruction(parser.OpConstant, 0),
				slim.MakeInstruction(parser.OpConstant, 1),
				slim.MakeInstruction(parser.OpConstant, 2),
				slim.MakeInstruction(parser.OpConstant, 3),
				slim.MakeInstruction(parser.OpClosure, 4, 1)),
			objectsArray(
				&slim.Int{Value: 1},
				&slim.Float{Value: 2.0},
				&slim.Char{Value: '3'},
				&slim.String{Value: "four"},
				compiledFunction(1, 0,
					slim.MakeInstruction(parser.OpConstant, 3),
					slim.MakeInstruction(parser.OpConstant, 2),
					slim.MakeInstruction(parser.OpSetLocal, 0),
					slim.MakeInstruction(parser.OpGetGlobal, 0),
					slim.MakeInstruction(parser.OpGetFree, 0)))))

	testBytecodeRemoveDuplicates(t,
		bytecode(
			concatInsts(
				slim.MakeInstruction(parser.OpConstant, 0),
				slim.MakeInstruction(parser.OpConstant, 1),
				slim.MakeInstruction(parser.OpConstant, 2),
				slim.MakeInstruction(parser.OpConstant, 3),
				slim.MakeInstruction(parser.OpConstant, 4)),
			objectsArray(
				&slim.Int{Value: 1},
				&slim.Int{Value: 2},
				&slim.Int{Value: 3},
				&slim.Int{Value: 1},
				&slim.Int{Value: 3})),
		bytecode(
			concatInsts(
				slim.MakeInstruction(parser.OpConstant, 0),
				slim.MakeInstruction(parser.OpConstant, 1),
				slim.MakeInstruction(parser.OpConstant, 2),
				slim.MakeInstruction(parser.OpConstant, 0),
				slim.MakeInstruction(parser.OpConstant, 2)),
			objectsArray(
				&slim.Int{Value: 1},
				&slim.Int{Value: 2},
				&slim.Int{Value: 3})))
}

func TestBytecode_CountObjects(t *testing.T) {
	b := bytecode(
		concatInsts(),
		objectsArray(
			&slim.Int{Value: 55},
			&slim.Int{Value: 66},
			&slim.Int{Value: 77},
			&slim.Int{Value: 88},
			compiledFunction(1, 0,
				slim.MakeInstruction(parser.OpConstant, 3),
				slim.MakeInstruction(parser.OpReturn, 1)),
			compiledFunction(1, 0,
				slim.MakeInstruction(parser.OpConstant, 2),
				slim.MakeInstruction(parser.OpReturn, 1)),
			compiledFunction(1, 0,
				slim.MakeInstruction(parser.OpConstant, 1),
				slim.MakeInstruction(parser.OpReturn, 1))))
	require.Equal(t, 7, b.CountObjects())
}

func fileSet(files ...srcfile) *parser.SourceFileSet {
	fileSet := parser.NewFileSet()
	for _, f := range files {
		fileSet.AddFile(f.name, -1, f.size)
	}
	return fileSet
}

func bytecodeFileSet(
	instructions []byte,
	constants []slim.Object,
	fileSet *parser.SourceFileSet,
) *slim.Bytecode {
	return &slim.Bytecode{
		FileSet:      fileSet,
		MainFunction: &slim.CompiledFunction{Instructions: instructions},
		Constants:    constants,
	}
}

func testBytecodeRemoveDuplicates(
	t *testing.T,
	input, expected *slim.Bytecode,
) {
	input.RemoveDuplicates()

	require.Equal(t, expected.FileSet, input.FileSet)
	require.Equal(t, expected.MainFunction, input.MainFunction)
	require.Equal(t, expected.Constants, input.Constants)
}

func testBytecodeSerialization(t *testing.T, b *slim.Bytecode) {
	var buf bytes.Buffer
	err := b.Encode(&buf)
	require.NoError(t, err)

	r := &slim.Bytecode{}
	err = r.Decode(bytes.NewReader(buf.Bytes()), nil)
	require.NoError(t, err)

	require.Equal(t, b.FileSet, r.FileSet)
	require.Equal(t, b.MainFunction, r.MainFunction)
	require.Equal(t, b.Constants, r.Constants)
}
