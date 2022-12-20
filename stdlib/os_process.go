package stdlib

import (
	"os"
	"syscall"

	"github.com/snple/slim"
)

func makeOSProcessState(state *os.ProcessState) *slim.ImmutableMap {
	return &slim.ImmutableMap{
		Value: map[string]slim.Object{
			"exited": &slim.UserFunction{
				Name:  "exited",
				Value: FuncARB(state.Exited),
			},
			"pid": &slim.UserFunction{
				Name:  "pid",
				Value: FuncARI(state.Pid),
			},
			"string": &slim.UserFunction{
				Name:  "string",
				Value: FuncARS(state.String),
			},
			"success": &slim.UserFunction{
				Name:  "success",
				Value: FuncARB(state.Success),
			},
		},
	}
}

func makeOSProcess(proc *os.Process) *slim.ImmutableMap {
	return &slim.ImmutableMap{
		Value: map[string]slim.Object{
			"kill": &slim.UserFunction{
				Name:  "kill",
				Value: FuncARE(proc.Kill),
			},
			"release": &slim.UserFunction{
				Name:  "release",
				Value: FuncARE(proc.Release),
			},
			"signal": &slim.UserFunction{
				Name: "signal",
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
					return wrapError(proc.Signal(syscall.Signal(i1))), nil
				},
			},
			"wait": &slim.UserFunction{
				Name: "wait",
				Value: func(args ...slim.Object) (slim.Object, error) {
					if len(args) != 0 {
						return nil, slim.ErrWrongNumArguments
					}
					state, err := proc.Wait()
					if err != nil {
						return wrapError(err), nil
					}
					return makeOSProcessState(state), nil
				},
			},
		},
	}
}
