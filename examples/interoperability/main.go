/*An example to demonstrate an alternative way to run slim functions from go.
 */
package main

import (
	"container/list"
	"context"
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/snple/slim"
)

// CallArgs holds function name to be executed and its required parameters with
// a channel to listen result of function.
type CallArgs struct {
	Func   string
	Params []slim.Object
	Result chan<- slim.Object
}

// NewGoProxy creates GoProxy object.
func NewGoProxy(ctx context.Context) *GoProxy {
	mod := new(GoProxy)
	mod.ctx = ctx
	mod.callbacks = make(map[string]slim.Object)
	mod.callChan = make(chan *CallArgs, 1)
	mod.moduleMap = map[string]slim.Object{
		"next":     &slim.UserFunction{Value: mod.next},
		"register": &slim.UserFunction{Value: mod.register},
		"args":     &slim.UserFunction{Value: mod.args},
	}
	mod.tasks = list.New()
	return mod
}

// GoProxy is a builtin slim module to register slim functions and run them.
type GoProxy struct {
	slim.ObjectImpl
	ctx       context.Context
	moduleMap map[string]slim.Object
	callbacks map[string]slim.Object
	callChan  chan *CallArgs
	tasks     *list.List
	mtx       sync.Mutex
}

// TypeName returns type name.
func (mod *GoProxy) TypeName() string {
	return "GoProxy"
}

func (mod *GoProxy) String() string {
	m := slim.ImmutableMap{Value: mod.moduleMap}
	return m.String()
}

// ModuleMap returns a map to add a builtin slim module.
func (mod *GoProxy) ModuleMap() map[string]slim.Object {
	return mod.moduleMap
}

// CallChan returns call channel which expects arguments to run a slim
// function.
func (mod *GoProxy) CallChan() chan<- *CallArgs {
	return mod.callChan
}

func (mod *GoProxy) next(args ...slim.Object) (slim.Object, error) {
	mod.mtx.Lock()
	defer mod.mtx.Unlock()
	select {
	case <-mod.ctx.Done():
		return slim.FalseValue, nil
	case args := <-mod.callChan:
		if args != nil {
			mod.tasks.PushBack(args)
		}
		return slim.TrueValue, nil
	}
}

func (mod *GoProxy) register(args ...slim.Object) (slim.Object, error) {
	if len(args) == 0 {
		return nil, slim.ErrWrongNumArguments
	}
	mod.mtx.Lock()
	defer mod.mtx.Unlock()

	switch v := args[0].(type) {
	case *slim.Map:
		mod.callbacks = v.Value
	case *slim.ImmutableMap:
		mod.callbacks = v.Value
	default:
		return nil, slim.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "map",
			Found:    args[0].TypeName(),
		}
	}
	return slim.UndefinedValue, nil
}

func (mod *GoProxy) args(args ...slim.Object) (slim.Object, error) {
	mod.mtx.Lock()
	defer mod.mtx.Unlock()

	if mod.tasks.Len() == 0 {
		return slim.UndefinedValue, nil
	}
	el := mod.tasks.Front()
	callArgs, ok := el.Value.(*CallArgs)
	if !ok || callArgs == nil {
		return nil, errors.New("invalid call arguments")
	}
	mod.tasks.Remove(el)
	f, ok := mod.callbacks[callArgs.Func]
	if !ok {
		return slim.UndefinedValue, nil
	}
	compiledFunc, ok := f.(*slim.CompiledFunction)
	if !ok {
		return slim.UndefinedValue, nil
	}
	params := callArgs.Params
	if params == nil {
		params = make([]slim.Object, 0)
	}
	// callable.VarArgs implementation is omitted.
	return &slim.ImmutableMap{
		Value: map[string]slim.Object{
			"result": &slim.UserFunction{
				Value: func(args ...slim.Object) (slim.Object, error) {
					if len(args) > 0 {
						callArgs.Result <- args[0]
						return slim.UndefinedValue, nil
					}
					callArgs.Result <- &slim.Error{
						Value: &slim.String{
							Value: slim.ErrWrongNumArguments.Error()},
					}
					return slim.UndefinedValue, nil
				}},
			"num_params": &slim.Int{Value: int64(compiledFunc.NumParameters)},
			"callable":   compiledFunc,
			"params":     &slim.Array{Value: params},
		},
	}, nil
}

// ProxySource is a slim script to handle bidirectional arguments flow between
// go and pure slim functions. Note: you should add more if conditions for
// different number of parameters.
// TODO: handle variadic functions.
var ProxySource = `
 export func(args) {
	 if is_undefined(args) {
		 return
	 }
	 callable := args.callable
	 if is_undefined(callable) {
		 return
	 }
	 result := args.result
	 num_params := args.num_params
	 v := undefined
	 // add more else if conditions for different number of parameters.
	 if num_params == 0 {
		 v = callable()
	 } else if num_params == 1 {
		 v = callable(args.params[0])
	 } else if num_params == 2 {
		 v = callable(args.params[0], args.params[1])
	 } else if num_params == 3 {
		 v = callable(args.params[0], args.params[1], args.params[2])
	 }
	 result(v)
 }
 `

func main() {
	src := `
	 // goproxy and proxy must be imported.
	 goproxy := import("goproxy")
	 proxy := import("proxy")

	 global := 0

	 callbacks := {
		 sum: func(a, b) {
			 return a + b
		 },
		 multiply: func(a, b) {
			 return a * b
		 },
		 increment: func() {
			 global++
			 return global
		 }
	 }

	 // Register callbacks to call them in goproxy loop.
	 goproxy.register(callbacks)

	 // goproxy loop waits for new call requests and run them with the help of
	 // "proxy" source module. Cancelling the context breaks the loop.
	 for goproxy.next() {
		 proxy(goproxy.args())
	 }
`
	// 5 seconds context timeout is enough for an example.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	script := slim.NewScript([]byte(src))
	moduleMap := slim.NewModuleMap()
	goproxy := NewGoProxy(ctx)
	// register modules
	moduleMap.AddBuiltinModule("goproxy", goproxy.ModuleMap())
	moduleMap.AddSourceModule("proxy", []byte(ProxySource))
	script.SetImports(moduleMap)

	compiled, err := script.Compile()
	if err != nil {
		panic(err)
	}

	// call "sum", "multiply", "increment" functions from slim in a new goroutine
	go func() {
		callChan := goproxy.CallChan()
		result := make(chan slim.Object, 1)
		// TODO: check slim error from result channel.
	loop:
		for {
			select {
			case <-ctx.Done():
				break loop
			default:
			}
			fmt.Println("Calling slim sum function")
			i1, i2 := rand.Int63n(100), rand.Int63n(100)
			callChan <- &CallArgs{Func: "sum",
				Params: []slim.Object{&slim.Int{Value: i1},
					&slim.Int{Value: i2}},
				Result: result,
			}
			v := <-result
			fmt.Printf("%d + %d = %v\n", i1, i2, v)

			fmt.Println("Calling slim multiply function")
			i1, i2 = rand.Int63n(20), rand.Int63n(20)
			callChan <- &CallArgs{Func: "multiply",
				Params: []slim.Object{&slim.Int{Value: i1},
					&slim.Int{Value: i2}},
				Result: result,
			}
			v = <-result
			fmt.Printf("%d * %d = %v\n", i1, i2, v)

			fmt.Println("Calling slim increment function")
			callChan <- &CallArgs{Func: "increment", Result: result}
			v = <-result
			fmt.Printf("increment = %v\n", v)
			time.Sleep(1 * time.Second)
		}
	}()

	if err := compiled.RunContext(ctx); err != nil {
		fmt.Println(err)
	}
}
