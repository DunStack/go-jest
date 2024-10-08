package jest

type TestFn[M any] func(j *J[M])

type J[M any] struct {
	jest *Jest[M]
	t    T
}

func (j *J[M]) Expect(v any) M {
	e := &Expect{
		T: j.t,
		v: v,
	}
	return j.jest.matcherFn(e)
}
