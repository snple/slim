package stdlib

import (
	"os"

	"github.com/snple/slim"
)

func makeOSFile(file *os.File) *slim.ImmutableMap {
	return &slim.ImmutableMap{
		Value: map[string]slim.Object{
			// chdir() => true/error
			"chdir": &slim.UserFunction{
				Name:  "chdir",
				Value: FuncARE(file.Chdir),
			}, //
			// chown(uid int, gid int) => true/error
			"chown": &slim.UserFunction{
				Name:  "chown",
				Value: FuncAIIRE(file.Chown),
			}, //
			// close() => error
			"close": &slim.UserFunction{
				Name:  "close",
				Value: FuncARE(file.Close),
			}, //
			// name() => string
			"name": &slim.UserFunction{
				Name:  "name",
				Value: FuncARS(file.Name),
			}, //
			// readdirnames(n int) => array(string)/error
			"readdirnames": &slim.UserFunction{
				Name:  "readdirnames",
				Value: FuncAIRSsE(file.Readdirnames),
			}, //
			// sync() => error
			"sync": &slim.UserFunction{
				Name:  "sync",
				Value: FuncARE(file.Sync),
			}, //
			// write(bytes) => int/error
			"write": &slim.UserFunction{
				Name:  "write",
				Value: FuncAYRIE(file.Write),
			}, //
			// write(string) => int/error
			"write_string": &slim.UserFunction{
				Name:  "write_string",
				Value: FuncASRIE(file.WriteString),
			}, //
			// read(bytes) => int/error
			"read": &slim.UserFunction{
				Name:  "read",
				Value: FuncAYRIE(file.Read),
			}, //
			// chmod(mode int) => error
			"chmod": &slim.UserFunction{
				Name: "chmod",
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
					return wrapError(file.Chmod(os.FileMode(i1))), nil
				},
			},
			// seek(offset int, whence int) => int/error
			"seek": &slim.UserFunction{
				Name: "seek",
				Value: func(args ...slim.Object) (slim.Object, error) {
					if len(args) != 2 {
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
					i2, ok := slim.ToInt(args[1])
					if !ok {
						return nil, slim.ErrInvalidArgumentType{
							Name:     "second",
							Expected: "int(compatible)",
							Found:    args[1].TypeName(),
						}
					}
					res, err := file.Seek(i1, i2)
					if err != nil {
						return wrapError(err), nil
					}
					return &slim.Int{Value: res}, nil
				},
			},
			// stat() => imap(fileinfo)/error
			"stat": &slim.UserFunction{
				Name: "stat",
				Value: func(args ...slim.Object) (slim.Object, error) {
					if len(args) != 0 {
						return nil, slim.ErrWrongNumArguments
					}
					return osStat(&slim.String{Value: file.Name()})
				},
			},
		},
	}
}
