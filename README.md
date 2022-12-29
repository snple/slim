# The Slim Language

Slim is a small, dynamic embedded script language for Go. It comes from [tengo](https://github.com/d5/tengo).

```golang
/* The Slim Language */
fmt := import("fmt")

each := func(seq, fn) {
    for x in seq { fn(x) }
}

sum := func(init, seq) {
    each(seq, func(x) { init += x })
    return init
}

fmt.println(sum(0, [1, 2, 3]))   // "6"
fmt.println(sum("", [1, 2, 3]))  // "123"
```

## Features

- Simple and highly readable
  [Syntax](https://github.com/snple/slim/blob/master/docs/tutorial.md)
  - Dynamic typing with type coercion
  - Higher-order functions and closures
  - Immutable values
- [Securely Embeddable](https://github.com/snple/slim/blob/master/docs/interoperability.md)
  and [Extensible](https://github.com/snple/slim/blob/master/docs/objects.md)
- Compiler/runtime written in native Go _(no external deps or cgo)_


## Quick Start

```
go get github.com/snple/slim
```

A simple Go example code that compiles/runs slim script code with some input/output values:

```golang
package main

import (
	"context"
	"fmt"

	"github.com/snple/slim"
)

func main() {
	// create a new Script instance
	script := slim.NewScript([]byte(
`each := func(seq, fn) {
    for x in seq { fn(x) }
}

sum := 0
mul := 1
each([a, b, c, d], func(x) {
    sum += x
    mul *= x
})`))

	// set values
	_ = script.Add("a", 1)
	_ = script.Add("b", 9)
	_ = script.Add("c", 8)
	_ = script.Add("d", 4)

	// run the script
	compiled, err := script.RunContext(context.Background())
	if err != nil {
		panic(err)
	}

	// retrieve values
	sum := compiled.Get("sum")
	mul := compiled.Get("mul")
	fmt.Println(sum, mul) // "22 288"
}
```

Or, if you need to evaluate a simple expression, you can use [Eval](https://pkg.go.dev/github.com/snple/slim#Eval) function instead:


```golang
res, err := slim.Eval(ctx,
	`input ? "success" : "fail"`,
	map[string]interface{}{"input": 1})
if err != nil {
	panic(err)
}
fmt.Println(res) // "success"
```

## References

- [Language Syntax](https://github.com/snple/slim/blob/master/docs/tutorial.md)
- [Object Types](https://github.com/snple/slim/blob/master/docs/objects.md)
- [Runtime Types](https://github.com/snple/slim/blob/master/docs/runtime-types.md)
  and [Operators](https://github.com/snple/slim/blob/master/docs/operators.md)
- [Builtin Functions](https://github.com/snple/slim/blob/master/docs/builtins.md)
- [Interoperability](https://github.com/snple/slim/blob/master/docs/interoperability.md)
- [slim CLI](https://github.com/snple/slim/blob/master/docs/slim-cli.md)
- [Standard Library](https://github.com/snple/slim/blob/master/docs/stdlib.md)
