package jest_test

import (
	"fmt"

	"github.com/dunstack/go-jest"
)

var _ jest.T = new(MockT)

type MockT struct {
	message string
	fail    bool
}

func (t *MockT) Helper() {}

func (t *MockT) Errorf(format string, args ...any) {
	t.fail = true
	t.message = fmt.Sprintf(format, args...)
}

func (t *MockT) Reset() {
	t.fail = false
	t.message = ""
}
