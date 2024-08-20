package jest

type MatcherFn[M any] func(e *Expect) M
