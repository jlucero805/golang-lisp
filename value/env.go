package value

type Env struct {
	Values map[Ident]Value
	Outer  *Env
}

func (e *Env) Set(ident Ident, setVal Value) {
	e.Values[ident] = setVal
}

func (e *Env) Get(ident Ident) Value {
	val, ok := e.Values[ident]
	if ok {
		return val
	}
	if e.Outer != nil {
		return e.Get(ident)
	}
	panic("Get failed")
}

func NewEnv() *Env {
	values := make(map[Ident]Value)
	values[Ident{Literal: "nil"}] = Number{}
	return &Env{
		Values: values,
	}
}

func WrapEnv(env *Env) *Env {
	newEnv := NewEnv()
	newEnv.Outer = env
	return newEnv
}
