package stdlib

import (
	"os/exec"

	"github.com/snple/slim"
)

func makeOSExecCommand(cmd *exec.Cmd) *slim.ImmutableMap {
	return &slim.ImmutableMap{
		Value: map[string]slim.Object{
			// combined_output() => bytes/error
			"combined_output": &slim.UserFunction{
				Name:  "combined_output",
				Value: FuncARYE(cmd.CombinedOutput),
			},
			// output() => bytes/error
			"output": &slim.UserFunction{
				Name:  "output",
				Value: FuncARYE(cmd.Output),
			}, //
			// run() => error
			"run": &slim.UserFunction{
				Name:  "run",
				Value: FuncARE(cmd.Run),
			}, //
			// start() => error
			"start": &slim.UserFunction{
				Name:  "start",
				Value: FuncARE(cmd.Start),
			}, //
			// wait() => error
			"wait": &slim.UserFunction{
				Name:  "wait",
				Value: FuncARE(cmd.Wait),
			}, //
			// set_path(path string)
			"set_path": &slim.UserFunction{
				Name: "set_path",
				Value: func(args ...slim.Object) (slim.Object, error) {
					if len(args) != 1 {
						return nil, slim.ErrWrongNumArguments
					}
					s1, ok := slim.ToString(args[0])
					if !ok {
						return nil, slim.ErrInvalidArgumentType{
							Name:     "first",
							Expected: "string(compatible)",
							Found:    args[0].TypeName(),
						}
					}
					cmd.Path = s1
					return slim.UndefinedValue, nil
				},
			},
			// set_dir(dir string)
			"set_dir": &slim.UserFunction{
				Name: "set_dir",
				Value: func(args ...slim.Object) (slim.Object, error) {
					if len(args) != 1 {
						return nil, slim.ErrWrongNumArguments
					}
					s1, ok := slim.ToString(args[0])
					if !ok {
						return nil, slim.ErrInvalidArgumentType{
							Name:     "first",
							Expected: "string(compatible)",
							Found:    args[0].TypeName(),
						}
					}
					cmd.Dir = s1
					return slim.UndefinedValue, nil
				},
			},
			// set_env(env array(string))
			"set_env": &slim.UserFunction{
				Name: "set_env",
				Value: func(args ...slim.Object) (slim.Object, error) {
					if len(args) != 1 {
						return nil, slim.ErrWrongNumArguments
					}

					var env []string
					var err error
					switch arg0 := args[0].(type) {
					case *slim.Array:
						env, err = stringArray(arg0.Value, "first")
						if err != nil {
							return nil, err
						}
					case *slim.ImmutableArray:
						env, err = stringArray(arg0.Value, "first")
						if err != nil {
							return nil, err
						}
					default:
						return nil, slim.ErrInvalidArgumentType{
							Name:     "first",
							Expected: "array",
							Found:    arg0.TypeName(),
						}
					}
					cmd.Env = env
					return slim.UndefinedValue, nil
				},
			},
			// process() => imap(process)
			"process": &slim.UserFunction{
				Name: "process",
				Value: func(args ...slim.Object) (slim.Object, error) {
					if len(args) != 0 {
						return nil, slim.ErrWrongNumArguments
					}
					return makeOSProcess(cmd.Process), nil
				},
			},
		},
	}
}
