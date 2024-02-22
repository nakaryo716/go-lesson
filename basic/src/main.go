package main

import (
	"fmt"
	"math"
)

func main() {
	num := math.Sqrt(145)
	fmt.Println("The Sqrt of 144 is:", num)

	fmt.Println("The number of pi is: ", math.Pi)

	fmt.Println(add(128, 512))

	firstText := "World!!!"
	secondText := "Hello"

	response1, response2 := swap(firstText, secondText)

	fmt.Println(response1, response2)
}

func add(number1 int, number2 int) int {
	response := number1 + number2
	return response
}

func swap(text1 string, text2 string) (string, string) {
	return text2, text1
}
