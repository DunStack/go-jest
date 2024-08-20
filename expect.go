package jest

import (
	"github.com/fatih/color"
)

type Expect struct {
	T
	v   any
	not bool
}

func (e Expect) Check(ok bool) bool {
	if e.not {
		return !ok
	}
	return ok
}

func (e Expect) Value() any {
	return e.v
}

func (e Expect) Not() bool {
	return e.not
}

func (e *Expect) Inverse() *Expect {
	e.not = !e.not
	return e
}

func (e Expect) WantSprint(v any) string {
	return e.WantSprintf("%#v", v)
}

func (e Expect) WantSprintf(format string, args ...any) string {
	if e.Not() {
		return color.GreenString("not "+format, args...)
	}
	return color.GreenString(format, args...)
}

func (e Expect) GotSprint(v any) string {
	return e.GotSprintf("%#v", v)
}

func (e Expect) GotSprintf(format string, args ...any) string {
	return color.RedString(format, args...)
}
