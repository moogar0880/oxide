package oxide

import (
	"testing"

	"github.com/moogar0880/collections/iter/assert"
)

type aTestStruct struct {
	X int
}

func TestOption_IsSome(t *testing.T) {
	testIO := []struct {
		name  string
		value interface{}
	}{
		{
			value: "foobar",
		},
		{
			value: 5,
		},
		{
			value: aTestStruct{X: 5},
		},
		{
			value: &aTestStruct{X: 10},
		},
	}

	for _, test := range testIO {
		t.Run(test.name, func(t *testing.T) {
			option := Some(test.value)

			assert.Equal(t, true, option.IsSome())
			assert.Equal(t, false, option.IsNone())
			assert.Equal(t, test.value, option.Value())

			actual, ok := option.Unpack()
			assert.Equal(t, test.value, actual)
			assert.Equal(t, true, ok)
		})
	}
}

func TestOption_IsNone(t *testing.T) {
	option := None[int]()

	assert.Equal(t, false, option.IsSome())
	assert.Equal(t, true, option.IsNone())
	assert.Equal(t, 0, option.Value())

	actual, ok := option.Unpack()
	assert.Equal(t, 0, actual)
	assert.Equal(t, false, ok)
}
