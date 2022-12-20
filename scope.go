package tengo

import "sync"

type Scope struct {
	variables map[string]*Variable
	lock      sync.RWMutex
}

func NewScope() *Scope {
	return &Scope{
		variables: make(map[string]*Variable),
	}
}

func (s *Scope) Clear() {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.variables = make(map[string]*Variable)
}

func (s *Scope) Len() int {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return len(s.variables)
}

func (s *Scope) IsEmpty() bool {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return s.Len() == 0
}

func (s *Scope) Contains(name string) bool {
	s.lock.RLock()
	defer s.lock.RUnlock()

	_, has := s.variables[name]
	return has
}

func (s *Scope) SetAny(name string, value any) error {
	obj, err := FromInterface(value)
	if err != nil {
		return err
	}

	s.SetValue(name, &Variable{
		name:  name,
		value: obj,
	})

	return nil
}

func (s *Scope) SetValue(name string, value *Variable) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.variables[name] = value
}

func (s *Scope) GetValue(name string) *Variable {
	s.lock.RLock()
	defer s.lock.RUnlock()

	if v, has := s.variables[name]; has {
		return v
	}

	return &Variable{
		name:  name,
		value: UndefinedValue,
	}
}

func (s *Scope) Remove(name string) bool {
	s.lock.Lock()
	defer s.lock.Unlock()

	if _, ok := s.variables[name]; !ok {
		return false
	}
	delete(s.variables, name)
	return true
}

func (s *Scope) GetAll() []*Variable {
	s.lock.RLock()
	defer s.lock.RUnlock()

	vars := make([]*Variable, 0, len(s.variables))

	for _, v := range s.variables {
		vars = append(vars, v)
	}

	return vars
}

func (s *Scope) GetAllNames() []string {
	s.lock.RLock()
	defer s.lock.RUnlock()

	names := make([]string, 0, len(s.variables))

	for name := range s.variables {
		names = append(names, name)
	}

	return names
}
