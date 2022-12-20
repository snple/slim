package stdlib

import (
	"regexp"

	"github.com/snple/slim"
)

func makeTextRegexp(re *regexp.Regexp) *slim.ImmutableMap {
	return &slim.ImmutableMap{
		Value: map[string]slim.Object{
			// match(text) => bool
			"match": &slim.UserFunction{
				Value: func(args ...slim.Object) (
					ret slim.Object,
					err error,
				) {
					if len(args) != 1 {
						err = slim.ErrWrongNumArguments
						return
					}

					s1, ok := slim.ToString(args[0])
					if !ok {
						err = slim.ErrInvalidArgumentType{
							Name:     "first",
							Expected: "string(compatible)",
							Found:    args[0].TypeName(),
						}
						return
					}

					if re.MatchString(s1) {
						ret = slim.TrueValue
					} else {
						ret = slim.FalseValue
					}

					return
				},
			},

			// find(text) 			=> array(array({text:,begin:,end:}))/undefined
			// find(text, maxCount) => array(array({text:,begin:,end:}))/undefined
			"find": &slim.UserFunction{
				Value: func(args ...slim.Object) (
					ret slim.Object,
					err error,
				) {
					numArgs := len(args)
					if numArgs != 1 && numArgs != 2 {
						err = slim.ErrWrongNumArguments
						return
					}

					s1, ok := slim.ToString(args[0])
					if !ok {
						err = slim.ErrInvalidArgumentType{
							Name:     "first",
							Expected: "string(compatible)",
							Found:    args[0].TypeName(),
						}
						return
					}

					if numArgs == 1 {
						m := re.FindStringSubmatchIndex(s1)
						if m == nil {
							ret = slim.UndefinedValue
							return
						}

						arr := &slim.Array{}
						for i := 0; i < len(m); i += 2 {
							arr.Value = append(arr.Value,
								&slim.ImmutableMap{
									Value: map[string]slim.Object{
										"text": &slim.String{
											Value: s1[m[i]:m[i+1]],
										},
										"begin": &slim.Int{
											Value: int64(m[i]),
										},
										"end": &slim.Int{
											Value: int64(m[i+1]),
										},
									}})
						}

						ret = &slim.Array{Value: []slim.Object{arr}}

						return
					}

					i2, ok := slim.ToInt(args[1])
					if !ok {
						err = slim.ErrInvalidArgumentType{
							Name:     "second",
							Expected: "int(compatible)",
							Found:    args[1].TypeName(),
						}
						return
					}
					m := re.FindAllStringSubmatchIndex(s1, i2)
					if m == nil {
						ret = slim.UndefinedValue
						return
					}

					arr := &slim.Array{}
					for _, m := range m {
						subMatch := &slim.Array{}
						for i := 0; i < len(m); i += 2 {
							subMatch.Value = append(subMatch.Value,
								&slim.ImmutableMap{
									Value: map[string]slim.Object{
										"text": &slim.String{
											Value: s1[m[i]:m[i+1]],
										},
										"begin": &slim.Int{
											Value: int64(m[i]),
										},
										"end": &slim.Int{
											Value: int64(m[i+1]),
										},
									}})
						}

						arr.Value = append(arr.Value, subMatch)
					}

					ret = arr

					return
				},
			},

			// replace(src, repl) => string
			"replace": &slim.UserFunction{
				Value: func(args ...slim.Object) (
					ret slim.Object,
					err error,
				) {
					if len(args) != 2 {
						err = slim.ErrWrongNumArguments
						return
					}

					s1, ok := slim.ToString(args[0])
					if !ok {
						err = slim.ErrInvalidArgumentType{
							Name:     "first",
							Expected: "string(compatible)",
							Found:    args[0].TypeName(),
						}
						return
					}

					s2, ok := slim.ToString(args[1])
					if !ok {
						err = slim.ErrInvalidArgumentType{
							Name:     "second",
							Expected: "string(compatible)",
							Found:    args[1].TypeName(),
						}
						return
					}

					s, ok := doTextRegexpReplace(re, s1, s2)
					if !ok {
						return nil, slim.ErrStringLimit
					}

					ret = &slim.String{Value: s}

					return
				},
			},

			// split(text) 			 => array(string)
			// split(text, maxCount) => array(string)
			"split": &slim.UserFunction{
				Value: func(args ...slim.Object) (
					ret slim.Object,
					err error,
				) {
					numArgs := len(args)
					if numArgs != 1 && numArgs != 2 {
						err = slim.ErrWrongNumArguments
						return
					}

					s1, ok := slim.ToString(args[0])
					if !ok {
						err = slim.ErrInvalidArgumentType{
							Name:     "first",
							Expected: "string(compatible)",
							Found:    args[0].TypeName(),
						}
						return
					}

					var i2 = -1
					if numArgs > 1 {
						i2, ok = slim.ToInt(args[1])
						if !ok {
							err = slim.ErrInvalidArgumentType{
								Name:     "second",
								Expected: "int(compatible)",
								Found:    args[1].TypeName(),
							}
							return
						}
					}

					arr := &slim.Array{}
					for _, s := range re.Split(s1, i2) {
						arr.Value = append(arr.Value,
							&slim.String{Value: s})
					}

					ret = arr

					return
				},
			},
		},
	}
}

// Size-limit checking implementation of regexp.ReplaceAllString.
func doTextRegexpReplace(re *regexp.Regexp, src, repl string) (string, bool) {
	idx := 0
	out := ""
	for _, m := range re.FindAllStringSubmatchIndex(src, -1) {
		var exp []byte
		exp = re.ExpandString(exp, repl, src, m)
		if len(out)+m[0]-idx+len(exp) > slim.MaxStringLen {
			return "", false
		}
		out += src[idx:m[0]] + string(exp)
		idx = m[1]
	}
	if idx < len(src) {
		if len(out)+len(src)-idx > slim.MaxStringLen {
			return "", false
		}
		out += src[idx:]
	}
	return out, true
}
