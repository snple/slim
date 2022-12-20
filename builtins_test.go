package slim_test

import (
	"errors"
	"reflect"
	"testing"

	"github.com/snple/slim"
)

func Test_builtinDelete(t *testing.T) {
	var builtinDelete func(args ...slim.Object) (slim.Object, error)
	for _, f := range slim.GetAllBuiltinFunctions() {
		if f.Name == "delete" {
			builtinDelete = f.Value
			break
		}
	}
	if builtinDelete == nil {
		t.Fatal("builtin delete not found")
	}
	type args struct {
		args []slim.Object
	}
	tests := []struct {
		name      string
		args      args
		want      slim.Object
		wantErr   bool
		wantedErr error
		target    interface{}
	}{
		{name: "invalid-arg", args: args{[]slim.Object{&slim.String{},
			&slim.String{}}}, wantErr: true,
			wantedErr: slim.ErrInvalidArgumentType{
				Name:     "first",
				Expected: "map",
				Found:    "string"},
		},
		{name: "no-args",
			wantErr: true, wantedErr: slim.ErrWrongNumArguments},
		{name: "empty-args", args: args{[]slim.Object{}}, wantErr: true,
			wantedErr: slim.ErrWrongNumArguments,
		},
		{name: "3-args", args: args{[]slim.Object{
			(*slim.Map)(nil), (*slim.String)(nil), (*slim.String)(nil)}},
			wantErr: true, wantedErr: slim.ErrWrongNumArguments,
		},
		{name: "nil-map-empty-key",
			args: args{[]slim.Object{&slim.Map{}, &slim.String{}}},
			want: slim.UndefinedValue,
		},
		{name: "nil-map-nonstr-key",
			args: args{[]slim.Object{
				&slim.Map{}, &slim.Int{}}}, wantErr: true,
			wantedErr: slim.ErrInvalidArgumentType{
				Name: "second", Expected: "string", Found: "int"},
		},
		{name: "nil-map-no-key",
			args: args{[]slim.Object{&slim.Map{}}}, wantErr: true,
			wantedErr: slim.ErrWrongNumArguments,
		},
		{name: "map-missing-key",
			args: args{
				[]slim.Object{
					&slim.Map{Value: map[string]slim.Object{
						"key": &slim.String{Value: "value"},
					}},
					&slim.String{Value: "key1"}}},
			want: slim.UndefinedValue,
			target: &slim.Map{
				Value: map[string]slim.Object{
					"key": &slim.String{
						Value: "value"}}},
		},
		{name: "map-emptied",
			args: args{
				[]slim.Object{
					&slim.Map{Value: map[string]slim.Object{
						"key": &slim.String{Value: "value"},
					}},
					&slim.String{Value: "key"}}},
			want:   slim.UndefinedValue,
			target: &slim.Map{Value: map[string]slim.Object{}},
		},
		{name: "map-multi-keys",
			args: args{
				[]slim.Object{
					&slim.Map{Value: map[string]slim.Object{
						"key1": &slim.String{Value: "value1"},
						"key2": &slim.Int{Value: 10},
					}},
					&slim.String{Value: "key1"}}},
			want: slim.UndefinedValue,
			target: &slim.Map{Value: map[string]slim.Object{
				"key2": &slim.Int{Value: 10}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := builtinDelete(tt.args.args...)
			if (err != nil) != tt.wantErr {
				t.Errorf("builtinDelete() error = %v, wantErr %v",
					err, tt.wantErr)
				return
			}
			if tt.wantErr && !errors.Is(err, tt.wantedErr) {
				if err.Error() != tt.wantedErr.Error() {
					t.Errorf("builtinDelete() error = %v, wantedErr %v",
						err, tt.wantedErr)
					return
				}
			}
			if got != tt.want {
				t.Errorf("builtinDelete() = %v, want %v", got, tt.want)
				return
			}
			if !tt.wantErr && tt.target != nil {
				switch v := tt.args.args[0].(type) {
				case *slim.Map, *slim.Array:
					if !reflect.DeepEqual(tt.target, tt.args.args[0]) {
						t.Errorf("builtinDelete() objects are not equal "+
							"got: %+v, want: %+v", tt.args.args[0], tt.target)
					}
				default:
					t.Errorf("builtinDelete() unsuporrted arg[0] type %s",
						v.TypeName())
					return
				}
			}
		})
	}
}

func Test_builtinSplice(t *testing.T) {
	var builtinSplice func(args ...slim.Object) (slim.Object, error)
	for _, f := range slim.GetAllBuiltinFunctions() {
		if f.Name == "splice" {
			builtinSplice = f.Value
			break
		}
	}
	if builtinSplice == nil {
		t.Fatal("builtin splice not found")
	}
	tests := []struct {
		name      string
		args      []slim.Object
		deleted   slim.Object
		Array     *slim.Array
		wantErr   bool
		wantedErr error
	}{
		{name: "no args", args: []slim.Object{}, wantErr: true,
			wantedErr: slim.ErrWrongNumArguments,
		},
		{name: "invalid args", args: []slim.Object{&slim.Map{}},
			wantErr: true,
			wantedErr: slim.ErrInvalidArgumentType{
				Name: "first", Expected: "array", Found: "map"},
		},
		{name: "invalid args",
			args:    []slim.Object{&slim.Array{}, &slim.String{}},
			wantErr: true,
			wantedErr: slim.ErrInvalidArgumentType{
				Name: "second", Expected: "int", Found: "string"},
		},
		{name: "negative index",
			args:      []slim.Object{&slim.Array{}, &slim.Int{Value: -1}},
			wantErr:   true,
			wantedErr: slim.ErrIndexOutOfBounds},
		{name: "non int count",
			args: []slim.Object{
				&slim.Array{}, &slim.Int{Value: 0},
				&slim.String{Value: ""}},
			wantErr: true,
			wantedErr: slim.ErrInvalidArgumentType{
				Name: "third", Expected: "int", Found: "string"},
		},
		{name: "negative count",
			args: []slim.Object{
				&slim.Array{Value: []slim.Object{
					&slim.Int{Value: 0},
					&slim.Int{Value: 1},
					&slim.Int{Value: 2}}},
				&slim.Int{Value: 0},
				&slim.Int{Value: -1}},
			wantErr:   true,
			wantedErr: slim.ErrIndexOutOfBounds,
		},
		{name: "insert with zero count",
			args: []slim.Object{
				&slim.Array{Value: []slim.Object{
					&slim.Int{Value: 0},
					&slim.Int{Value: 1},
					&slim.Int{Value: 2}}},
				&slim.Int{Value: 0},
				&slim.Int{Value: 0},
				&slim.String{Value: "b"}},
			deleted: &slim.Array{Value: []slim.Object{}},
			Array: &slim.Array{Value: []slim.Object{
				&slim.String{Value: "b"},
				&slim.Int{Value: 0},
				&slim.Int{Value: 1},
				&slim.Int{Value: 2}}},
		},
		{name: "insert",
			args: []slim.Object{
				&slim.Array{Value: []slim.Object{
					&slim.Int{Value: 0},
					&slim.Int{Value: 1},
					&slim.Int{Value: 2}}},
				&slim.Int{Value: 1},
				&slim.Int{Value: 0},
				&slim.String{Value: "c"},
				&slim.String{Value: "d"}},
			deleted: &slim.Array{Value: []slim.Object{}},
			Array: &slim.Array{Value: []slim.Object{
				&slim.Int{Value: 0},
				&slim.String{Value: "c"},
				&slim.String{Value: "d"},
				&slim.Int{Value: 1},
				&slim.Int{Value: 2}}},
		},
		{name: "insert with zero count",
			args: []slim.Object{
				&slim.Array{Value: []slim.Object{
					&slim.Int{Value: 0},
					&slim.Int{Value: 1},
					&slim.Int{Value: 2}}},
				&slim.Int{Value: 1},
				&slim.Int{Value: 0},
				&slim.String{Value: "c"},
				&slim.String{Value: "d"}},
			deleted: &slim.Array{Value: []slim.Object{}},
			Array: &slim.Array{Value: []slim.Object{
				&slim.Int{Value: 0},
				&slim.String{Value: "c"},
				&slim.String{Value: "d"},
				&slim.Int{Value: 1},
				&slim.Int{Value: 2}}},
		},
		{name: "insert with delete",
			args: []slim.Object{
				&slim.Array{Value: []slim.Object{
					&slim.Int{Value: 0},
					&slim.Int{Value: 1},
					&slim.Int{Value: 2}}},
				&slim.Int{Value: 1},
				&slim.Int{Value: 1},
				&slim.String{Value: "c"},
				&slim.String{Value: "d"}},
			deleted: &slim.Array{
				Value: []slim.Object{&slim.Int{Value: 1}}},
			Array: &slim.Array{Value: []slim.Object{
				&slim.Int{Value: 0},
				&slim.String{Value: "c"},
				&slim.String{Value: "d"},
				&slim.Int{Value: 2}}},
		},
		{name: "insert with delete multi",
			args: []slim.Object{
				&slim.Array{Value: []slim.Object{
					&slim.Int{Value: 0},
					&slim.Int{Value: 1},
					&slim.Int{Value: 2}}},
				&slim.Int{Value: 1},
				&slim.Int{Value: 2},
				&slim.String{Value: "c"},
				&slim.String{Value: "d"}},
			deleted: &slim.Array{Value: []slim.Object{
				&slim.Int{Value: 1},
				&slim.Int{Value: 2}}},
			Array: &slim.Array{
				Value: []slim.Object{
					&slim.Int{Value: 0},
					&slim.String{Value: "c"},
					&slim.String{Value: "d"}}},
		},
		{name: "delete all with positive count",
			args: []slim.Object{
				&slim.Array{Value: []slim.Object{
					&slim.Int{Value: 0},
					&slim.Int{Value: 1},
					&slim.Int{Value: 2}}},
				&slim.Int{Value: 0},
				&slim.Int{Value: 3}},
			deleted: &slim.Array{Value: []slim.Object{
				&slim.Int{Value: 0},
				&slim.Int{Value: 1},
				&slim.Int{Value: 2}}},
			Array: &slim.Array{Value: []slim.Object{}},
		},
		{name: "delete all with big count",
			args: []slim.Object{
				&slim.Array{Value: []slim.Object{
					&slim.Int{Value: 0},
					&slim.Int{Value: 1},
					&slim.Int{Value: 2}}},
				&slim.Int{Value: 0},
				&slim.Int{Value: 5}},
			deleted: &slim.Array{Value: []slim.Object{
				&slim.Int{Value: 0},
				&slim.Int{Value: 1},
				&slim.Int{Value: 2}}},
			Array: &slim.Array{Value: []slim.Object{}},
		},
		{name: "nothing2",
			args: []slim.Object{
				&slim.Array{Value: []slim.Object{
					&slim.Int{Value: 0},
					&slim.Int{Value: 1},
					&slim.Int{Value: 2}}}},
			Array: &slim.Array{Value: []slim.Object{}},
			deleted: &slim.Array{Value: []slim.Object{
				&slim.Int{Value: 0},
				&slim.Int{Value: 1},
				&slim.Int{Value: 2}}},
		},
		{name: "pop without count",
			args: []slim.Object{
				&slim.Array{Value: []slim.Object{
					&slim.Int{Value: 0},
					&slim.Int{Value: 1},
					&slim.Int{Value: 2}}},
				&slim.Int{Value: 2}},
			deleted: &slim.Array{Value: []slim.Object{&slim.Int{Value: 2}}},
			Array: &slim.Array{Value: []slim.Object{
				&slim.Int{Value: 0}, &slim.Int{Value: 1}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := builtinSplice(tt.args...)
			if (err != nil) != tt.wantErr {
				t.Errorf("builtinSplice() error = %v, wantErr %v",
					err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.deleted) {
				t.Errorf("builtinSplice() = %v, want %v", got, tt.deleted)
			}
			if tt.wantErr && tt.wantedErr.Error() != err.Error() {
				t.Errorf("builtinSplice() error = %v, wantedErr %v",
					err, tt.wantedErr)
			}
			if tt.Array != nil && !reflect.DeepEqual(tt.Array, tt.args[0]) {
				t.Errorf("builtinSplice() arrays are not equal expected"+
					" %s, got %s", tt.Array, tt.args[0].(*slim.Array))
			}
		})
	}
}

func Test_builtinRange(t *testing.T) {
	var builtinRange func(args ...slim.Object) (slim.Object, error)
	for _, f := range slim.GetAllBuiltinFunctions() {
		if f.Name == "range" {
			builtinRange = f.Value
			break
		}
	}
	if builtinRange == nil {
		t.Fatal("builtin range not found")
	}
	tests := []struct {
		name      string
		args      []slim.Object
		result    *slim.Array
		wantErr   bool
		wantedErr error
	}{
		{name: "no args", args: []slim.Object{}, wantErr: true,
			wantedErr: slim.ErrWrongNumArguments,
		},
		{name: "single args", args: []slim.Object{&slim.Map{}},
			wantErr:   true,
			wantedErr: slim.ErrWrongNumArguments,
		},
		{name: "4 args", args: []slim.Object{&slim.Map{}, &slim.String{}, &slim.String{}, &slim.String{}},
			wantErr:   true,
			wantedErr: slim.ErrWrongNumArguments,
		},
		{name: "invalid start",
			args:    []slim.Object{&slim.String{}, &slim.String{}},
			wantErr: true,
			wantedErr: slim.ErrInvalidArgumentType{
				Name: "start", Expected: "int", Found: "string"},
		},
		{name: "invalid stop",
			args:    []slim.Object{&slim.Int{}, &slim.String{}},
			wantErr: true,
			wantedErr: slim.ErrInvalidArgumentType{
				Name: "stop", Expected: "int", Found: "string"},
		},
		{name: "invalid step",
			args:    []slim.Object{&slim.Int{}, &slim.Int{}, &slim.String{}},
			wantErr: true,
			wantedErr: slim.ErrInvalidArgumentType{
				Name: "step", Expected: "int", Found: "string"},
		},
		{name: "zero step",
			args:      []slim.Object{&slim.Int{}, &slim.Int{}, &slim.Int{}}, //must greate than 0
			wantErr:   true,
			wantedErr: slim.ErrInvalidRangeStep,
		},
		{name: "negative step",
			args:      []slim.Object{&slim.Int{}, &slim.Int{}, intObject(-2)}, //must greate than 0
			wantErr:   true,
			wantedErr: slim.ErrInvalidRangeStep,
		},
		{name: "same bound",
			args:    []slim.Object{&slim.Int{}, &slim.Int{}},
			wantErr: false,
			result: &slim.Array{
				Value: nil,
			},
		},
		{name: "positive range",
			args:    []slim.Object{&slim.Int{}, &slim.Int{Value: 5}},
			wantErr: false,
			result: &slim.Array{
				Value: []slim.Object{
					intObject(0),
					intObject(1),
					intObject(2),
					intObject(3),
					intObject(4),
				},
			},
		},
		{name: "negative range",
			args:    []slim.Object{&slim.Int{}, &slim.Int{Value: -5}},
			wantErr: false,
			result: &slim.Array{
				Value: []slim.Object{
					intObject(0),
					intObject(-1),
					intObject(-2),
					intObject(-3),
					intObject(-4),
				},
			},
		},

		{name: "positive with step",
			args:    []slim.Object{&slim.Int{}, &slim.Int{Value: 5}, &slim.Int{Value: 2}},
			wantErr: false,
			result: &slim.Array{
				Value: []slim.Object{
					intObject(0),
					intObject(2),
					intObject(4),
				},
			},
		},

		{name: "negative with step",
			args:    []slim.Object{&slim.Int{}, &slim.Int{Value: -10}, &slim.Int{Value: 2}},
			wantErr: false,
			result: &slim.Array{
				Value: []slim.Object{
					intObject(0),
					intObject(-2),
					intObject(-4),
					intObject(-6),
					intObject(-8),
				},
			},
		},

		{name: "large range",
			args:    []slim.Object{intObject(-10), intObject(10), &slim.Int{Value: 3}},
			wantErr: false,
			result: &slim.Array{
				Value: []slim.Object{
					intObject(-10),
					intObject(-7),
					intObject(-4),
					intObject(-1),
					intObject(2),
					intObject(5),
					intObject(8),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := builtinRange(tt.args...)
			if (err != nil) != tt.wantErr {
				t.Errorf("builtinRange() error = %v, wantErr %v",
					err, tt.wantErr)
				return
			}
			if tt.wantErr && tt.wantedErr.Error() != err.Error() {
				t.Errorf("builtinRange() error = %v, wantedErr %v",
					err, tt.wantedErr)
			}
			if tt.result != nil && !reflect.DeepEqual(tt.result, got) {
				t.Errorf("builtinRange() arrays are not equal expected"+
					" %s, got %s", tt.result, got.(*slim.Array))
			}
		})
	}
}
