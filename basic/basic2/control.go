package basic2

import (
	"fmt"
)

func ForLoop() {
	for i := 10; i >= 0; i-- {
		fmt.Printf("%v second remain\n", i)
	}

	fmt.Println("Out of loop")
}

func IfControl(number int) {
	if number < 100 {
		fmt.Println("Under 100")
	} else if number == 110 {
		fmt.Println("110")
	} else {
		fmt.Println("over 100")
	}

}
