package basic_1

import (
	"fmt"
	"math"
)

func Add(number1 int, number2 int) int {
	response := number1 + number2
	return response
}

func Swap(text1 string, text2 string) (string, string) {
	return text2, text1
}

func Conversions() {
	pi := math.Pi

	intPi := int(pi)

	fmt.Println("PI value is:", pi, " int_PI value is:", intPi)
}
