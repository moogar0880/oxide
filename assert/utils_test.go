package assert

import (
	"fmt"
	"testing"
)

type astruct struct {
	Foo string
	Bar string
}

func TestEqual(t *testing.T) {
	testIO := []struct {
		left   interface{}
		right  interface{}
		expect bool
	}{
		{
			left:   nil,
			right:  nil,
			expect: true,
		},
		{
			left:   0,
			right:  0,
			expect: true,
		},
		{
			left:   5,
			right:  10,
			expect: false,
		},
		{
			left:   "foo",
			right:  "bar",
			expect: false,
		},
		{
			left:   "foobar",
			right:  "foobar",
			expect: true,
		},
		{
			left:   astruct{Foo: "foo", Bar: "bar"},
			right:  astruct{Foo: "foo", Bar: "oof"},
			expect: false,
		},
		{
			left:   astruct{Foo: "foo", Bar: "bar"},
			right:  astruct{Foo: "foo", Bar: "bar"},
			expect: true,
		},
		{
			left:   []byte("{}"),
			right:  []byte("{}"),
			expect: true,
		},
		{
			left:   []byte("{}"),
			right:  []byte("[{}, {}]"),
			expect: false,
		},
		{
			left:   []byte("{}"),
			right:  12,
			expect: false,
		},
		{
			left:   func() []byte { return nil },
			right:  func() []byte { return nil },
			expect: false,
		},
		{
			left:   []byte("{}"),
			right:  func() []byte { return nil },
			expect: false,
		},
		{
			left:   func() (data []byte) { return },
			right:  []byte("{}"),
			expect: false,
		},
		{
			left:   func() {},
			right:  func() {},
			expect: false,
		},
	}

	for _, test := range testIO {
		name := fmt.Sprintf("TestEqual(%v, %v)", test.left, test.right)
		t.Run(name, func(t *testing.T) {
			actual := Equal(new(mockTest), test.left, test.right)

			if actual != test.expect {
				t.Errorf("actual value did not match expected\nexpected: %v\n actual=%v", test.expect, actual)
			}
		})
	}
}
