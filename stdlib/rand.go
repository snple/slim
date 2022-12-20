package stdlib

import (
	"math/rand"

	"github.com/snple/slim"
)

var randModule = map[string]slim.Object{
	"int": &slim.UserFunction{
		Name:  "int",
		Value: FuncARI64(rand.Int63),
	},
	"float": &slim.UserFunction{
		Name:  "float",
		Value: FuncARF(rand.Float64),
	},
	"intn": &slim.UserFunction{
		Name:  "intn",
		Value: FuncAI64RI64(rand.Int63n),
	},
	"exp_float": &slim.UserFunction{
		Name:  "exp_float",
		Value: FuncARF(rand.ExpFloat64),
	},
	"norm_float": &slim.UserFunction{
		Name:  "norm_float",
		Value: FuncARF(rand.NormFloat64),
	},
	"perm": &slim.UserFunction{
		Name:  "perm",
		Value: FuncAIRIs(rand.Perm),
	},
	"seed": &slim.UserFunction{
		Name:  "seed",
		Value: FuncAI64R(rand.Seed),
	},
	"read": &slim.UserFunction{
		Name: "read",
		Value: func(args ...slim.Object) (ret slim.Object, err error) {
			if len(args) != 1 {
				return nil, slim.ErrWrongNumArguments
			}
			y1, ok := args[0].(*slim.Bytes)
			if !ok {
				return nil, slim.ErrInvalidArgumentType{
					Name:     "first",
					Expected: "bytes",
					Found:    args[0].TypeName(),
				}
			}
			res, err := rand.Read(y1.Value)
			if err != nil {
				ret = wrapError(err)
				return
			}
			return &slim.Int{Value: int64(res)}, nil
		},
	},
	"rand": &slim.UserFunction{
		Name: "rand",
		Value: func(args ...slim.Object) (slim.Object, error) {
			if len(args) != 1 {
				return nil, slim.ErrWrongNumArguments
			}
			i1, ok := slim.ToInt64(args[0])
			if !ok {
				return nil, slim.ErrInvalidArgumentType{
					Name:     "first",
					Expected: "int(compatible)",
					Found:    args[0].TypeName(),
				}
			}
			src := rand.NewSource(i1)
			return randRand(rand.New(src)), nil
		},
	},
}

func randRand(r *rand.Rand) *slim.ImmutableMap {
	return &slim.ImmutableMap{
		Value: map[string]slim.Object{
			"int": &slim.UserFunction{
				Name:  "int",
				Value: FuncARI64(r.Int63),
			},
			"float": &slim.UserFunction{
				Name:  "float",
				Value: FuncARF(r.Float64),
			},
			"intn": &slim.UserFunction{
				Name:  "intn",
				Value: FuncAI64RI64(r.Int63n),
			},
			"exp_float": &slim.UserFunction{
				Name:  "exp_float",
				Value: FuncARF(r.ExpFloat64),
			},
			"norm_float": &slim.UserFunction{
				Name:  "norm_float",
				Value: FuncARF(r.NormFloat64),
			},
			"perm": &slim.UserFunction{
				Name:  "perm",
				Value: FuncAIRIs(r.Perm),
			},
			"seed": &slim.UserFunction{
				Name:  "seed",
				Value: FuncAI64R(r.Seed),
			},
			"read": &slim.UserFunction{
				Name: "read",
				Value: func(args ...slim.Object) (
					ret slim.Object,
					err error,
				) {
					if len(args) != 1 {
						return nil, slim.ErrWrongNumArguments
					}
					y1, ok := args[0].(*slim.Bytes)
					if !ok {
						return nil, slim.ErrInvalidArgumentType{
							Name:     "first",
							Expected: "bytes",
							Found:    args[0].TypeName(),
						}
					}
					res, err := r.Read(y1.Value)
					if err != nil {
						ret = wrapError(err)
						return
					}
					return &slim.Int{Value: int64(res)}, nil
				},
			},
		},
	}
}
