package jest

import "reflect"

type MatcherFn[M any] func(e *Expect) M

func NewBuiltinMatcher(e *Expect) BuiltinMatcher {
	return BuiltinMatcher{e}
}

type BuiltinMatcher struct {
	*Expect
}

func (m BuiltinMatcher) Not() BuiltinMatcher {
	m.Inverse()
	return m
}

func (m BuiltinMatcher) ToBe(v any) {
	m.Helper()
	if g := m.Value(); m.Check(v != g) {
		m.Errorf("\nwant: %s\ngot: %s",
			m.WantSprint(v),
			m.GotSprint(g),
		)
	}
}

func (m BuiltinMatcher) ToEqual(v any) {
	m.Helper()
	if g := m.Value(); m.Check(!reflect.DeepEqual(v, g)) {
		m.Errorf("\nwant: %s\ngot: %s",
			m.WantSprint(v),
			m.GotSprint(g),
		)
	}
}
