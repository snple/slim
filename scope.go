package slim

import (
	"fmt"
	"sync"

	"github.com/snple/slim/parser"
)

type Variables map[string]*Variable

func NewVariables() Variables {
	return make(map[string]*Variable)
}

func (vs Variables) Clear() {
	vs = make(map[string]*Variable)
}

func (vs Variables) Len() int {
	return len(vs)
}

func (vs Variables) IsEmpty() bool {
	return vs.Len() == 0
}

func (vs Variables) Contains(name string) bool {
	_, has := vs[name]
	return has
}

func (vs Variables) SetAny(name string, value any) error {
	obj, err := FromInterface(value)
	if err != nil {
		return err
	}

	vs.SetValue(name, &Variable{
		name:  name,
		value: obj,
	})

	return nil
}

func (vs Variables) SetValue(name string, value *Variable) {
	vs[name] = value
}

func (vs Variables) GetValue(name string) *Variable {
	if v, has := vs[name]; has {
		return v
	}

	return &Variable{
		name:  name,
		value: UndefinedValue,
	}
}

func (vs Variables) Remove(name string) bool {
	if _, ok := vs[name]; !ok {
		return false
	}
	delete(vs, name)
	return true
}

type Scope struct {
	lock            sync.RWMutex
	symbolTable     *SymbolTable
	globals         []Object
	globalIndexes   map[string]int
	maxAllocs       int64
	maxConstObjects int
}

func NewScope(variables Variables) *Scope {
	s := &Scope{
		symbolTable:     NewSymbolTable(),
		globals:         make([]Object, GlobalsSize),
		maxAllocs:       -1,
		maxConstObjects: -1,
	}

	for idx, fn := range builtinFuncs {
		s.symbolTable.DefineBuiltin(idx, fn.Name)
	}

	idx := 0
	for name, value := range variables {
		symbol := s.symbolTable.Define(name)
		if symbol.Index != idx {
			panic(fmt.Errorf("wrong symbol index: %d != %d",
				idx, symbol.Index))
		}

		s.globals[symbol.Index] = value.value

		idx += 1
	}

	return s
}

func (s *Scope) Complie(name string, src []byte, modules ModuleGetter) (*Bytecode, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	if name == "" {
		name = "main"
	}

	// parser
	fileSet := parser.NewFileSet()
	input := []byte(src)
	srcFile := fileSet.AddFile(fmt.Sprintf("(%s)", name), -1, len(input))

	p := parser.NewParser(srcFile, input, nil)
	file, err := p.ParseFile()
	if err != nil {
		return nil, err
	}

	// compile
	c := NewCompiler(srcFile, s.symbolTable, nil, modules, nil)
	if err := c.Compile(file); err != nil {
		return nil, err
	}

	// reduce globals size
	// s.globals = s.globals[:s.symbolTable.MaxSymbols()+1]

	s.globalIndexes = make(map[string]int, len(s.globals))
	for _, name := range s.symbolTable.Names() {
		symbol, _, _ := s.symbolTable.Resolve(name, false)
		if symbol.Scope == ScopeGlobal {
			s.globalIndexes[name] = symbol.Index
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

	return bytecode, nil
}

func (s *Scope) Run(bytecode *Bytecode) error {
	s.lock.Lock()
	defer s.lock.Unlock()

	vm := NewVM(bytecode, s.globals, s.maxAllocs)
	err := vm.Run()
	if err != nil {
		return err
	}

	return nil
}

func (s *Scope) ComplieAndRun(name string, src []byte, modules ModuleGetter) error {
	bytecode, err := s.Complie(name, src, modules)
	if err != nil {
		return err
	}

	return s.Run(bytecode)
}

func (s *Scope) IsDefined(name string) bool {
	s.lock.RLock()
	defer s.lock.RUnlock()

	idx, ok := s.globalIndexes[name]
	if !ok {
		return false
	}
	v := s.globals[idx]
	if v == nil {
		return false
	}
	return v != UndefinedValue
}

func (s *Scope) Get(name string) *Variable {
	s.lock.RLock()
	defer s.lock.RUnlock()

	value := UndefinedValue
	if idx, ok := s.globalIndexes[name]; ok {
		value = s.globals[idx]
		if value == nil {
			value = UndefinedValue
		}
	}
	return &Variable{
		name:  name,
		value: value,
	}
}

func (s *Scope) Set(name string, value interface{}) error {
	s.lock.Lock()
	defer s.lock.Unlock()

	obj, err := FromInterface(value)
	if err != nil {
		return err
	}
	idx, ok := s.globalIndexes[name]
	if !ok {
		return fmt.Errorf("'%s' is not defined", name)
	}
	s.globals[idx] = obj
	return nil
}

func (s *Scope) GetAll() Variables {
	s.lock.RLock()
	defer s.lock.RUnlock()

	vars := make(map[string]*Variable)

OUT:
	for name, idx := range s.globalIndexes {
		value := s.globals[idx]
		if value == nil {
			value = UndefinedValue
		}

		if value.CanCall() {
			continue
		}

		switch value.TypeName() {
		case "immutable-map":
			continue OUT
			// switch val := value.(type) {
			// case *ImmutableMap:
			// 	for _, v := range val.Value {
			// 		_ = v
			// 		if v.CanCall() {
			// 			continue OUT
			// 		}
			// 	}
			// }
		case "user-function", "compiled-function":
			continue OUT
		}

		vars[name] = &Variable{
			name:  name,
			value: value,
		}
	}

	return vars
}
