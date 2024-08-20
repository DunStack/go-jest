package jest

type Expect struct {
	T
	v   any
	not bool
}

func (e Expect) Value() any {
	return e.v
}

func (e Expect) Not() any {
	return e.not
}

func (e *Expect) Inverse() *Expect {
	e.not = !e.not
	return e
}
