package jest

func Extend[M any](matcherFn MatcherFn[M]) *Jest[M] {
	return &Jest[M]{
		matcherFn: matcherFn,
	}
}

type Jest[M any] struct {
	matcherFn MatcherFn[M]
}

func (jest *Jest[M]) Test(t T, fn TestFn[M]) {
	j := &J[M]{
		jest: jest,
		t:    t,
	}
	fn(j)
}

var DefaultJest = Extend(NewBuiltinMatcher)

func Test(t T, fn TestFn[BuiltinMatcher]) {
	DefaultJest.Test(t, fn)
}
