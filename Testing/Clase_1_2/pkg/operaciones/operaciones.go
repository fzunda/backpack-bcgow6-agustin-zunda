package operaciones

import "errors"

func Restar(num1, num2 int) int {
	return num1 - num2
}

func Dividir(num, den int) (result int, err error) {
	if den == 0 {
		err = errors.New("The denominator must be greater than zero")
	}
	result = num / den
	return
}
