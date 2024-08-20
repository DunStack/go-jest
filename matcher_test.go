package jest_test

import (
	"fmt"
	"testing"

	"github.com/dunstack/go-jest"
	"github.com/fatih/color"
)

func TestBuiltinMatcher(t *testing.T) {
	mt := new(MockT)

	t.Run("ToBe", func(t *testing.T) {
		jest.Test(mt, func(j *jest.J[jest.BuiltinMatcher]) {
			mt.Reset()
			if j.Expect(1).ToBe(1); mt.fail {
				t.Fail()
			}

			mt.Reset()
			if j.Expect(1).Not().ToBe(1); mt.message != fmt.Sprintf("\nwant: %s\ngot: %s", color.GreenString("not 1"), color.RedString("1")) {
				t.Fail()
			}

			mt.Reset()
			if j.Expect("a").ToBe("b"); mt.message != fmt.Sprintf("\nwant: %s\ngot: %s", color.GreenString(`"b"`), color.RedString(`"a"`)) {
				t.Fail()
			}

			mt.Reset()
			if j.Expect("a").Not().ToBe("b"); mt.fail {
				t.Fail()
			}
		})
	})

	t.Run("ToEqual", func(t *testing.T) {
		jest.Test(mt, func(j *jest.J[jest.BuiltinMatcher]) {
			mt.Reset()
			if j.Expect([]int{1}).ToEqual([]int{1}); mt.fail {
				t.Fail()
			}

			mt.Reset()
			if j.Expect([]int{1}).Not().ToEqual([]int{1}); mt.message != fmt.Sprintf("\nwant: %s\ngot: %s", color.GreenString("not []int{1}"), color.RedString("[]int{1}")) {
				t.Fail()
			}

			mt.Reset()
			if j.Expect(map[string]string{"a": "b"}).ToEqual(map[string]string{"a": "c"}); mt.message != fmt.Sprintf("\nwant: %s\ngot: %s", color.GreenString(`map[string]string{"a":"c"}`), color.RedString(`map[string]string{"a":"b"}`)) {
				t.Fail()
			}

			mt.Reset()
			if j.Expect(map[string]string{"a": "b"}).Not().ToEqual(map[string]string{"a": "c"}); mt.fail {
				t.Fail()
			}
		})
	})

	t.Run("ToHaveLength", func(t *testing.T) {
		jest.Test(mt, func(j *jest.J[jest.BuiltinMatcher]) {
			mt.Reset()
			if j.Expect("abc").ToHaveLength(3); mt.fail {
				t.Fail()
			}

			mt.Reset()
			if j.Expect("abc").Not().ToHaveLength(3); mt.message != fmt.Sprintf("\nwant: %s\ngot: %s", color.GreenString("not 3"), color.RedString("3")) {
				t.Fail()
			}

			mt.Reset()
			if j.Expect([]int{1, 2}).ToHaveLength(3); mt.message != fmt.Sprintf("\nwant: %s\ngot: %s", color.GreenString("3"), color.RedString("2")) {
				t.Fail()
			}

			mt.Reset()
			if j.Expect([]int{1, 2}).Not().ToHaveLength(3); mt.fail {
				t.Fail()
			}

			mt.Reset()
			if j.Expect(1).ToHaveLength(0); mt.message != "cannot get length of 1" {
				t.Fail()
			}
		})
	})
}
