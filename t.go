package jest

type T interface {
	Helper()
	Errorf(format string, args ...any)
}
