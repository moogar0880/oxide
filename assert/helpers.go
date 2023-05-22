package assert

func IsNegative(i *int) bool {
	return *i < 0
}

func IsPositive(i *int) bool {
	return !IsNegative(i)
}

func IsEven(i *int) bool {
	return *i%2 == 0
}

func Is100(i *int) bool {
	return *i == 100
}
