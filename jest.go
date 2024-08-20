package jest

func Extend[M any](matcherFn MatcherFn[M]) *Jest[M] {
	return &Jest[M]{
		matcherFn: matcherFn,
	}
}

type Jest[M any] struct {
	matcherFn MatcherFn[M]
}

func (jest *Jest[M]) Test(t T, fn func(j *J[M])) {
	j := &J[M]{
		jest: jest,
		t:    t,
	}
	fn(j)
}
