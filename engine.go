package slim

import (
	"fmt"
	"strings"

	"github.com/snple/slim/parser"
)

type Engine struct {
	modules         ModuleGetter
	maxAllocs       int64
	maxConstObjects int
}

func NewEngine() *Engine {
	return &Engine{
		maxAllocs:       -1,
		maxConstObjects: -1,
	}
}

func (s *Engine) SetImports(modules ModuleGetter) {
	s.modules = modules
}

func (s *Engine) SetMaxAllocs(n int64) {
	s.maxAllocs = n
}

func (s *Engine) SetMaxConstObjects(n int) {
	s.maxConstObjects = n
}

func (s *Engine) Run(script string) error {
	return s.RunWithScope(NewScope(), script)
}

func (e *Engine) RunWithScope(scope *Scope, script string) error {
	complied, err := e.Compile(scope, script)
	if err != nil {
		return err
	}

	return complied.Run()
}

func (s *Engine) Compile(scope *Scope, script string) (*Compiled2, error) {
	// scope
	names := scope.GetAllNames()

	// symbol table
	symbolTable := NewSymbolTable()
	for idx, fn := range builtinFuncs {
		symbolTable.DefineBuiltin(idx, fn.Name)
	}

	globals := make([]Object, GlobalsSize)

	for idx, name := range names {
		symbol := symbolTable.Define(name)
		if symbol.Index != idx {
			panic(fmt.Errorf("wrong symbol index: %d != %d",
				idx, symbol.Index))
		}
		globals[symbol.Index] = scope.GetValue(name).value
	}

	// parser
	fileSet := parser.NewFileSet()
	input := []byte(script)
	srcFile := fileSet.AddFile("(main)", -1, len(input))

	p := parser.NewParser(srcFile, input, nil)
	file, err := p.ParseFile()
	if err != nil {
		return nil, err
	}

	// compile
	c := NewCompiler(srcFile, symbolTable, nil, s.modules, nil)
	if err := c.Compile(file); err != nil {
		return nil, err
	}

	// reduce globals size
	globals = globals[:symbolTable.MaxSymbols()+1]

	// global symbol names to indexes
	globalIndexes := make(map[string]int, len(globals))
	for _, name := range symbolTable.Names() {
		symbol, _, _ := symbolTable.Resolve(name, false)
		if symbol.Scope == ScopeGlobal {
			globalIndexes[name] = symbol.Index
		}
	}

	// remove duplicates from constants
	bytecode := c.Bytecode()
	bytecode.RemoveDuplicates()

	// check the constant objects limit
	if s.maxConstObjects >= 0 {
		cnt := bytecode.CountObjects()
		if cnt > s.maxConstObjects {
			return nil, fmt.Errorf("exceeding constant objects limit: %d", cnt)
		}
	}

	return &Compiled2{
		bytecode:      bytecode,
		globalIndexes: globalIndexes,
		globals:       globals,
		maxAllocs:     s.maxAllocs,
		scope:         scope,
	}, nil
}

type Compiled2 struct {
	bytecode      *Bytecode
	globalIndexes map[string]int // global symbol name to index
	globals       []Object
	maxAllocs     int64
	scope         *Scope
}

func (c *Compiled2) Run() error {
	vm := NewVM(c.bytecode, c.globals, c.maxAllocs)
	if err := vm.Run(); err != nil {
		return err
	}

OUT:
	for name, idx := range c.globalIndexes {
		if strings.HasPrefix(name, "_") {
			continue
		}

		value := c.globals[idx]
		if value == nil {
			value = UndefinedValue
		}

		switch value.TypeName() {
		case "immutable-map":
			switch val := value.(type) {
			case *ImmutableMap:
				for _, v := range val.Value {
					if v.String() == "<user-function>" || v.String() == "<compiled-function>" {
						continue OUT
					}
				}
			}
		case "user-function", "compiled-function":
			continue OUT
		}

		c.scope.SetValue(name, &Variable{
			name:  name,
			value: value,
		})
	}

	return nil
}

func (s *Compiled2) Scope() *Scope {
	return s.scope
}
