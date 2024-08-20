package jest_test

import (
	"testing"

	"github.com/dunstack/go-jest"
)

func TestBuiltinMatcher(t *testing.T) {
	mt := new(MockT)

	t.Run("ToBe", func(t *testing.T) {
		jest.Test(mt, func(j *jest.J[jest.BuiltinMatcher]) {
			if j.Expect(1).ToBe(1); mt.fail {
				t.Fail()
			}

			mt.Reset()
			if j.Expect(1).Not().ToBe(1); mt.message != "\nwant: not 1\ngot: 1" {
				t.Fail()
			}

			mt.Reset()
			if j.Expect("a").ToBe("b"); mt.message != "\nwant: \"b\"\ngot: \"a\"" {
				t.Fail()
			}

			mt.Reset()
			if j.Expect("a").Not().ToBe("b"); mt.fail {
				t.Fail()
			}
		})
	})
}
