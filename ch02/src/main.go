package main

import "fmt"

func main() {
	var x int = 256
	var y int = 128
	fmt.Println(x + y)

	var flag bool = true
	fmt.Println(flag)

	text := "Hello"

	text2 := text

	text = "World!"

	fmt.Println(text)
	fmt.Println(text2)
}
