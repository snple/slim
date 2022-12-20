package stdlib

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"

	"github.com/snple/slim"
)

var osModule = map[string]slim.Object{
	"o_rdonly":            &slim.Int{Value: int64(os.O_RDONLY)},
	"o_wronly":            &slim.Int{Value: int64(os.O_WRONLY)},
	"o_rdwr":              &slim.Int{Value: int64(os.O_RDWR)},
	"o_append":            &slim.Int{Value: int64(os.O_APPEND)},
	"o_create":            &slim.Int{Value: int64(os.O_CREATE)},
	"o_excl":              &slim.Int{Value: int64(os.O_EXCL)},
	"o_sync":              &slim.Int{Value: int64(os.O_SYNC)},
	"o_trunc":             &slim.Int{Value: int64(os.O_TRUNC)},
	"mode_dir":            &slim.Int{Value: int64(os.ModeDir)},
	"mode_append":         &slim.Int{Value: int64(os.ModeAppend)},
	"mode_exclusive":      &slim.Int{Value: int64(os.ModeExclusive)},
	"mode_temporary":      &slim.Int{Value: int64(os.ModeTemporary)},
	"mode_symlink":        &slim.Int{Value: int64(os.ModeSymlink)},
	"mode_device":         &slim.Int{Value: int64(os.ModeDevice)},
	"mode_named_pipe":     &slim.Int{Value: int64(os.ModeNamedPipe)},
	"mode_socket":         &slim.Int{Value: int64(os.ModeSocket)},
	"mode_setuid":         &slim.Int{Value: int64(os.ModeSetuid)},
	"mode_setgui":         &slim.Int{Value: int64(os.ModeSetgid)},
	"mode_char_device":    &slim.Int{Value: int64(os.ModeCharDevice)},
	"mode_sticky":         &slim.Int{Value: int64(os.ModeSticky)},
	"mode_type":           &slim.Int{Value: int64(os.ModeType)},
	"mode_perm":           &slim.Int{Value: int64(os.ModePerm)},
	"path_separator":      &slim.Char{Value: os.PathSeparator},
	"path_list_separator": &slim.Char{Value: os.PathListSeparator},
	"dev_null":            &slim.String{Value: os.DevNull},
	"seek_set":            &slim.Int{Value: int64(io.SeekStart)},
	"seek_cur":            &slim.Int{Value: int64(io.SeekCurrent)},
	"seek_end":            &slim.Int{Value: int64(io.SeekEnd)},
	"args": &slim.UserFunction{
		Name:  "args",
		Value: osArgs,
	}, // args() => array(string)
	"chdir": &slim.UserFunction{
		Name:  "chdir",
		Value: FuncASRE(os.Chdir),
	}, // chdir(dir string) => error
	"chmod": osFuncASFmRE("chmod", os.Chmod), // chmod(name string, mode int) => error
	"chown": &slim.UserFunction{
		Name:  "chown",
		Value: FuncASIIRE(os.Chown),
	}, // chown(name string, uid int, gid int) => error
	"clearenv": &slim.UserFunction{
		Name:  "clearenv",
		Value: FuncAR(os.Clearenv),
	}, // clearenv()
	"environ": &slim.UserFunction{
		Name:  "environ",
		Value: FuncARSs(os.Environ),
	}, // environ() => array(string)
	"exit": &slim.UserFunction{
		Name:  "exit",
		Value: FuncAIR(os.Exit),
	}, // exit(code int)
	"expand_env": &slim.UserFunction{
		Name:  "expand_env",
		Value: osExpandEnv,
	}, // expand_env(s string) => string
	"getegid": &slim.UserFunction{
		Name:  "getegid",
		Value: FuncARI(os.Getegid),
	}, // getegid() => int
	"getenv": &slim.UserFunction{
		Name:  "getenv",
		Value: FuncASRS(os.Getenv),
	}, // getenv(s string) => string
	"geteuid": &slim.UserFunction{
		Name:  "geteuid",
		Value: FuncARI(os.Geteuid),
	}, // geteuid() => int
	"getgid": &slim.UserFunction{
		Name:  "getgid",
		Value: FuncARI(os.Getgid),
	}, // getgid() => int
	"getgroups": &slim.UserFunction{
		Name:  "getgroups",
		Value: FuncARIsE(os.Getgroups),
	}, // getgroups() => array(string)/error
	"getpagesize": &slim.UserFunction{
		Name:  "getpagesize",
		Value: FuncARI(os.Getpagesize),
	}, // getpagesize() => int
	"getpid": &slim.UserFunction{
		Name:  "getpid",
		Value: FuncARI(os.Getpid),
	}, // getpid() => int
	"getppid": &slim.UserFunction{
		Name:  "getppid",
		Value: FuncARI(os.Getppid),
	}, // getppid() => int
	"getuid": &slim.UserFunction{
		Name:  "getuid",
		Value: FuncARI(os.Getuid),
	}, // getuid() => int
	"getwd": &slim.UserFunction{
		Name:  "getwd",
		Value: FuncARSE(os.Getwd),
	}, // getwd() => string/error
	"hostname": &slim.UserFunction{
		Name:  "hostname",
		Value: FuncARSE(os.Hostname),
	}, // hostname() => string/error
	"lchown": &slim.UserFunction{
		Name:  "lchown",
		Value: FuncASIIRE(os.Lchown),
	}, // lchown(name string, uid int, gid int) => error
	"link": &slim.UserFunction{
		Name:  "link",
		Value: FuncASSRE(os.Link),
	}, // link(oldname string, newname string) => error
	"lookup_env": &slim.UserFunction{
		Name:  "lookup_env",
		Value: osLookupEnv,
	}, // lookup_env(key string) => string/false
	"mkdir":     osFuncASFmRE("mkdir", os.Mkdir),        // mkdir(name string, perm int) => error
	"mkdir_all": osFuncASFmRE("mkdir_all", os.MkdirAll), // mkdir_all(name string, perm int) => error
	"readlink": &slim.UserFunction{
		Name:  "readlink",
		Value: FuncASRSE(os.Readlink),
	}, // readlink(name string) => string/error
	"remove": &slim.UserFunction{
		Name:  "remove",
		Value: FuncASRE(os.Remove),
	}, // remove(name string) => error
	"remove_all": &slim.UserFunction{
		Name:  "remove_all",
		Value: FuncASRE(os.RemoveAll),
	}, // remove_all(name string) => error
	"rename": &slim.UserFunction{
		Name:  "rename",
		Value: FuncASSRE(os.Rename),
	}, // rename(oldpath string, newpath string) => error
	"setenv": &slim.UserFunction{
		Name:  "setenv",
		Value: FuncASSRE(os.Setenv),
	}, // setenv(key string, value string) => error
	"symlink": &slim.UserFunction{
		Name:  "symlink",
		Value: FuncASSRE(os.Symlink),
	}, // symlink(oldname string newname string) => error
	"temp_dir": &slim.UserFunction{
		Name:  "temp_dir",
		Value: FuncARS(os.TempDir),
	}, // temp_dir() => string
	"truncate": &slim.UserFunction{
		Name:  "truncate",
		Value: FuncASI64RE(os.Truncate),
	}, // truncate(name string, size int) => error
	"unsetenv": &slim.UserFunction{
		Name:  "unsetenv",
		Value: FuncASRE(os.Unsetenv),
	}, // unsetenv(key string) => error
	"create": &slim.UserFunction{
		Name:  "create",
		Value: osCreate,
	}, // create(name string) => imap(file)/error
	"open": &slim.UserFunction{
		Name:  "open",
		Value: osOpen,
	}, // open(name string) => imap(file)/error
	"open_file": &slim.UserFunction{
		Name:  "open_file",
		Value: osOpenFile,
	}, // open_file(name string, flag int, perm int) => imap(file)/error
	"find_process": &slim.UserFunction{
		Name:  "find_process",
		Value: osFindProcess,
	}, // find_process(pid int) => imap(process)/error
	"start_process": &slim.UserFunction{
		Name:  "start_process",
		Value: osStartProcess,
	}, // start_process(name string, argv array(string), dir string, env array(string)) => imap(process)/error
	"exec_look_path": &slim.UserFunction{
		Name:  "exec_look_path",
		Value: FuncASRSE(exec.LookPath),
	}, // exec_look_path(file) => string/error
	"exec": &slim.UserFunction{
		Name:  "exec",
		Value: osExec,
	}, // exec(name, args...) => command
	"stat": &slim.UserFunction{
		Name:  "stat",
		Value: osStat,
	}, // stat(name) => imap(fileinfo)/error
	"read_file": &slim.UserFunction{
		Name:  "read_file",
		Value: osReadFile,
	}, // readfile(name) => array(byte)/error
}

func osReadFile(args ...slim.Object) (ret slim.Object, err error) {
	if len(args) != 1 {
		return nil, slim.ErrWrongNumArguments
	}
	fname, ok := slim.ToString(args[0])
	if !ok {
		return nil, slim.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "string(compatible)",
			Found:    args[0].TypeName(),
		}
	}
	bytes, err := ioutil.ReadFile(fname)
	if err != nil {
		return wrapError(err), nil
	}
	if len(bytes) > slim.MaxBytesLen {
		return nil, slim.ErrBytesLimit
	}
	return &slim.Bytes{Value: bytes}, nil
}

func osStat(args ...slim.Object) (ret slim.Object, err error) {
	if len(args) != 1 {
		return nil, slim.ErrWrongNumArguments
	}
	fname, ok := slim.ToString(args[0])
	if !ok {
		return nil, slim.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "string(compatible)",
			Found:    args[0].TypeName(),
		}
	}
	stat, err := os.Stat(fname)
	if err != nil {
		return wrapError(err), nil
	}
	fstat := &slim.ImmutableMap{
		Value: map[string]slim.Object{
			"name":  &slim.String{Value: stat.Name()},
			"mtime": &slim.Time{Value: stat.ModTime()},
			"size":  &slim.Int{Value: stat.Size()},
			"mode":  &slim.Int{Value: int64(stat.Mode())},
		},
	}
	if stat.IsDir() {
		fstat.Value["directory"] = slim.TrueValue
	} else {
		fstat.Value["directory"] = slim.FalseValue
	}
	return fstat, nil
}

func osCreate(args ...slim.Object) (slim.Object, error) {
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
	res, err := os.Create(s1)
	if err != nil {
		return wrapError(err), nil
	}
	return makeOSFile(res), nil
}

func osOpen(args ...slim.Object) (slim.Object, error) {
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
	res, err := os.Open(s1)
	if err != nil {
		return wrapError(err), nil
	}
	return makeOSFile(res), nil
}

func osOpenFile(args ...slim.Object) (slim.Object, error) {
	if len(args) != 3 {
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
	i2, ok := slim.ToInt(args[1])
	if !ok {
		return nil, slim.ErrInvalidArgumentType{
			Name:     "second",
			Expected: "int(compatible)",
			Found:    args[1].TypeName(),
		}
	}
	i3, ok := slim.ToInt(args[2])
	if !ok {
		return nil, slim.ErrInvalidArgumentType{
			Name:     "third",
			Expected: "int(compatible)",
			Found:    args[2].TypeName(),
		}
	}
	res, err := os.OpenFile(s1, i2, os.FileMode(i3))
	if err != nil {
		return wrapError(err), nil
	}
	return makeOSFile(res), nil
}

func osArgs(args ...slim.Object) (slim.Object, error) {
	if len(args) != 0 {
		return nil, slim.ErrWrongNumArguments
	}
	arr := &slim.Array{}
	for _, osArg := range os.Args {
		if len(osArg) > slim.MaxStringLen {
			return nil, slim.ErrStringLimit
		}
		arr.Value = append(arr.Value, &slim.String{Value: osArg})
	}
	return arr, nil
}

func osFuncASFmRE(
	name string,
	fn func(string, os.FileMode) error,
) *slim.UserFunction {
	return &slim.UserFunction{
		Name: name,
		Value: func(args ...slim.Object) (slim.Object, error) {
			if len(args) != 2 {
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
			i2, ok := slim.ToInt64(args[1])
			if !ok {
				return nil, slim.ErrInvalidArgumentType{
					Name:     "second",
					Expected: "int(compatible)",
					Found:    args[1].TypeName(),
				}
			}
			return wrapError(fn(s1, os.FileMode(i2))), nil
		},
	}
}

func osLookupEnv(args ...slim.Object) (slim.Object, error) {
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
	res, ok := os.LookupEnv(s1)
	if !ok {
		return slim.FalseValue, nil
	}
	if len(res) > slim.MaxStringLen {
		return nil, slim.ErrStringLimit
	}
	return &slim.String{Value: res}, nil
}

func osExpandEnv(args ...slim.Object) (slim.Object, error) {
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
	var vlen int
	var failed bool
	s := os.Expand(s1, func(k string) string {
		if failed {
			return ""
		}
		v := os.Getenv(k)

		// this does not count the other texts that are not being replaced
		// but the code checks the final length at the end
		vlen += len(v)
		if vlen > slim.MaxStringLen {
			failed = true
			return ""
		}
		return v
	})
	if failed || len(s) > slim.MaxStringLen {
		return nil, slim.ErrStringLimit
	}
	return &slim.String{Value: s}, nil
}

func osExec(args ...slim.Object) (slim.Object, error) {
	if len(args) == 0 {
		return nil, slim.ErrWrongNumArguments
	}
	name, ok := slim.ToString(args[0])
	if !ok {
		return nil, slim.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "string(compatible)",
			Found:    args[0].TypeName(),
		}
	}
	var execArgs []string
	for idx, arg := range args[1:] {
		execArg, ok := slim.ToString(arg)
		if !ok {
			return nil, slim.ErrInvalidArgumentType{
				Name:     fmt.Sprintf("args[%d]", idx),
				Expected: "string(compatible)",
				Found:    args[1+idx].TypeName(),
			}
		}
		execArgs = append(execArgs, execArg)
	}
	return makeOSExecCommand(exec.Command(name, execArgs...)), nil
}

func osFindProcess(args ...slim.Object) (slim.Object, error) {
	if len(args) != 1 {
		return nil, slim.ErrWrongNumArguments
	}
	i1, ok := slim.ToInt(args[0])
	if !ok {
		return nil, slim.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "int(compatible)",
			Found:    args[0].TypeName(),
		}
	}
	proc, err := os.FindProcess(i1)
	if err != nil {
		return wrapError(err), nil
	}
	return makeOSProcess(proc), nil
}

func osStartProcess(args ...slim.Object) (slim.Object, error) {
	if len(args) != 4 {
		return nil, slim.ErrWrongNumArguments
	}
	name, ok := slim.ToString(args[0])
	if !ok {
		return nil, slim.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "string(compatible)",
			Found:    args[0].TypeName(),
		}
	}
	var argv []string
	var err error
	switch arg1 := args[1].(type) {
	case *slim.Array:
		argv, err = stringArray(arg1.Value, "second")
		if err != nil {
			return nil, err
		}
	case *slim.ImmutableArray:
		argv, err = stringArray(arg1.Value, "second")
		if err != nil {
			return nil, err
		}
	default:
		return nil, slim.ErrInvalidArgumentType{
			Name:     "second",
			Expected: "array",
			Found:    arg1.TypeName(),
		}
	}

	dir, ok := slim.ToString(args[2])
	if !ok {
		return nil, slim.ErrInvalidArgumentType{
			Name:     "third",
			Expected: "string(compatible)",
			Found:    args[2].TypeName(),
		}
	}

	var env []string
	switch arg3 := args[3].(type) {
	case *slim.Array:
		env, err = stringArray(arg3.Value, "fourth")
		if err != nil {
			return nil, err
		}
	case *slim.ImmutableArray:
		env, err = stringArray(arg3.Value, "fourth")
		if err != nil {
			return nil, err
		}
	default:
		return nil, slim.ErrInvalidArgumentType{
			Name:     "fourth",
			Expected: "array",
			Found:    arg3.TypeName(),
		}
	}

	proc, err := os.StartProcess(name, argv, &os.ProcAttr{
		Dir: dir,
		Env: env,
	})
	if err != nil {
		return wrapError(err), nil
	}
	return makeOSProcess(proc), nil
}

func stringArray(arr []slim.Object, argName string) ([]string, error) {
	var sarr []string
	for idx, elem := range arr {
		str, ok := elem.(*slim.String)
		if !ok {
			return nil, slim.ErrInvalidArgumentType{
				Name:     fmt.Sprintf("%s[%d]", argName, idx),
				Expected: "string",
				Found:    elem.TypeName(),
			}
		}
		sarr = append(sarr, str.Value)
	}
	return sarr, nil
}
