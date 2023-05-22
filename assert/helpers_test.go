package assert

import "testing"

func TestStubs(t *testing.T) {
	for i := -100; i < 100; i++ {
		neg := IsNegative(&i)
		pos := IsPositive(&i)
		even := IsEven(&i)
		hundred := Is100(&i)

		if i < 0 {
			Equal(t, true, neg)
			Equal(t, false, pos)
		} else {
			Equal(t, false, neg)
			Equal(t, true, pos)
		}

		if i%2 == 0 {
			Equal(t, true, even)
		} else {
			Equal(t, false, even)
		}

		if i == 100 {
			Equal(t, true, hundred)
		} else {
			Equal(t, false, hundred)
		}
	}
}
