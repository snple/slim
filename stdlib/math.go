package stdlib

import (
	"math"

	"github.com/snple/slim"
)

var mathModule = map[string]slim.Object{
	"e":                      &slim.Float{Value: math.E},
	"pi":                     &slim.Float{Value: math.Pi},
	"phi":                    &slim.Float{Value: math.Phi},
	"sqrt2":                  &slim.Float{Value: math.Sqrt2},
	"sqrtE":                  &slim.Float{Value: math.SqrtE},
	"sqrtPi":                 &slim.Float{Value: math.SqrtPi},
	"sqrtPhi":                &slim.Float{Value: math.SqrtPhi},
	"ln2":                    &slim.Float{Value: math.Ln2},
	"log2E":                  &slim.Float{Value: math.Log2E},
	"ln10":                   &slim.Float{Value: math.Ln10},
	"log10E":                 &slim.Float{Value: math.Log10E},
	"maxFloat32":             &slim.Float{Value: math.MaxFloat32},
	"smallestNonzeroFloat32": &slim.Float{Value: math.SmallestNonzeroFloat32},
	"maxFloat64":             &slim.Float{Value: math.MaxFloat64},
	"smallestNonzeroFloat64": &slim.Float{Value: math.SmallestNonzeroFloat64},
	"maxInt":                 &slim.Int{Value: math.MaxInt},
	"minInt":                 &slim.Int{Value: math.MinInt},
	"maxInt8":                &slim.Int{Value: math.MaxInt8},
	"minInt8":                &slim.Int{Value: math.MinInt8},
	"maxInt16":               &slim.Int{Value: math.MaxInt16},
	"minInt16":               &slim.Int{Value: math.MinInt16},
	"maxInt32":               &slim.Int{Value: math.MaxInt32},
	"minInt32":               &slim.Int{Value: math.MinInt32},
	"maxInt64":               &slim.Int{Value: math.MaxInt64},
	"minInt64":               &slim.Int{Value: math.MinInt64},
	"abs": &slim.UserFunction{
		Name:  "abs",
		Value: FuncAFRF(math.Abs),
	},
	"acos": &slim.UserFunction{
		Name:  "acos",
		Value: FuncAFRF(math.Acos),
	},
	"acosh": &slim.UserFunction{
		Name:  "acosh",
		Value: FuncAFRF(math.Acosh),
	},
	"asin": &slim.UserFunction{
		Name:  "asin",
		Value: FuncAFRF(math.Asin),
	},
	"asinh": &slim.UserFunction{
		Name:  "asinh",
		Value: FuncAFRF(math.Asinh),
	},
	"atan": &slim.UserFunction{
		Name:  "atan",
		Value: FuncAFRF(math.Atan),
	},
	"atan2": &slim.UserFunction{
		Name:  "atan2",
		Value: FuncAFFRF(math.Atan2),
	},
	"atanh": &slim.UserFunction{
		Name:  "atanh",
		Value: FuncAFRF(math.Atanh),
	},
	"cbrt": &slim.UserFunction{
		Name:  "cbrt",
		Value: FuncAFRF(math.Cbrt),
	},
	"ceil": &slim.UserFunction{
		Name:  "ceil",
		Value: FuncAFRF(math.Ceil),
	},
	"copysign": &slim.UserFunction{
		Name:  "copysign",
		Value: FuncAFFRF(math.Copysign),
	},
	"cos": &slim.UserFunction{
		Name:  "cos",
		Value: FuncAFRF(math.Cos),
	},
	"cosh": &slim.UserFunction{
		Name:  "cosh",
		Value: FuncAFRF(math.Cosh),
	},
	"dim": &slim.UserFunction{
		Name:  "dim",
		Value: FuncAFFRF(math.Dim),
	},
	"erf": &slim.UserFunction{
		Name:  "erf",
		Value: FuncAFRF(math.Erf),
	},
	"erfc": &slim.UserFunction{
		Name:  "erfc",
		Value: FuncAFRF(math.Erfc),
	},
	"exp": &slim.UserFunction{
		Name:  "exp",
		Value: FuncAFRF(math.Exp),
	},
	"exp2": &slim.UserFunction{
		Name:  "exp2",
		Value: FuncAFRF(math.Exp2),
	},
	"expm1": &slim.UserFunction{
		Name:  "expm1",
		Value: FuncAFRF(math.Expm1),
	},
	"floor": &slim.UserFunction{
		Name:  "floor",
		Value: FuncAFRF(math.Floor),
	},
	"gamma": &slim.UserFunction{
		Name:  "gamma",
		Value: FuncAFRF(math.Gamma),
	},
	"hypot": &slim.UserFunction{
		Name:  "hypot",
		Value: FuncAFFRF(math.Hypot),
	},
	"ilogb": &slim.UserFunction{
		Name:  "ilogb",
		Value: FuncAFRI(math.Ilogb),
	},
	"inf": &slim.UserFunction{
		Name:  "inf",
		Value: FuncAIRF(math.Inf),
	},
	"is_inf": &slim.UserFunction{
		Name:  "is_inf",
		Value: FuncAFIRB(math.IsInf),
	},
	"is_nan": &slim.UserFunction{
		Name:  "is_nan",
		Value: FuncAFRB(math.IsNaN),
	},
	"j0": &slim.UserFunction{
		Name:  "j0",
		Value: FuncAFRF(math.J0),
	},
	"j1": &slim.UserFunction{
		Name:  "j1",
		Value: FuncAFRF(math.J1),
	},
	"jn": &slim.UserFunction{
		Name:  "jn",
		Value: FuncAIFRF(math.Jn),
	},
	"ldexp": &slim.UserFunction{
		Name:  "ldexp",
		Value: FuncAFIRF(math.Ldexp),
	},
	"log": &slim.UserFunction{
		Name:  "log",
		Value: FuncAFRF(math.Log),
	},
	"log10": &slim.UserFunction{
		Name:  "log10",
		Value: FuncAFRF(math.Log10),
	},
	"log1p": &slim.UserFunction{
		Name:  "log1p",
		Value: FuncAFRF(math.Log1p),
	},
	"log2": &slim.UserFunction{
		Name:  "log2",
		Value: FuncAFRF(math.Log2),
	},
	"logb": &slim.UserFunction{
		Name:  "logb",
		Value: FuncAFRF(math.Logb),
	},
	"max": &slim.UserFunction{
		Name:  "max",
		Value: FuncAFFRF(math.Max),
	},
	"min": &slim.UserFunction{
		Name:  "min",
		Value: FuncAFFRF(math.Min),
	},
	"mod": &slim.UserFunction{
		Name:  "mod",
		Value: FuncAFFRF(math.Mod),
	},
	"nan": &slim.UserFunction{
		Name:  "nan",
		Value: FuncARF(math.NaN),
	},
	"nextafter": &slim.UserFunction{
		Name:  "nextafter",
		Value: FuncAFFRF(math.Nextafter),
	},
	"pow": &slim.UserFunction{
		Name:  "pow",
		Value: FuncAFFRF(math.Pow),
	},
	"pow10": &slim.UserFunction{
		Name:  "pow10",
		Value: FuncAIRF(math.Pow10),
	},
	"remainder": &slim.UserFunction{
		Name:  "remainder",
		Value: FuncAFFRF(math.Remainder),
	},
	"signbit": &slim.UserFunction{
		Name:  "signbit",
		Value: FuncAFRB(math.Signbit),
	},
	"sin": &slim.UserFunction{
		Name:  "sin",
		Value: FuncAFRF(math.Sin),
	},
	"sinh": &slim.UserFunction{
		Name:  "sinh",
		Value: FuncAFRF(math.Sinh),
	},
	"sqrt": &slim.UserFunction{
		Name:  "sqrt",
		Value: FuncAFRF(math.Sqrt),
	},
	"tan": &slim.UserFunction{
		Name:  "tan",
		Value: FuncAFRF(math.Tan),
	},
	"tanh": &slim.UserFunction{
		Name:  "tanh",
		Value: FuncAFRF(math.Tanh),
	},
	"trunc": &slim.UserFunction{
		Name:  "trunc",
		Value: FuncAFRF(math.Trunc),
	},
	"y0": &slim.UserFunction{
		Name:  "y0",
		Value: FuncAFRF(math.Y0),
	},
	"y1": &slim.UserFunction{
		Name:  "y1",
		Value: FuncAFRF(math.Y1),
	},
	"yn": &slim.UserFunction{
		Name:  "yn",
		Value: FuncAIFRF(math.Yn),
	},
}
