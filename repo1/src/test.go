package main

import (
	"fmt"
)

func main() {
	var inputX int
	var inputY int

	fmt.Scan(&inputX, &inputY)

	multiple(inputX, inputY)

}

func multiple(a int, b int) {
	for i := 1; i <= b; i++ {
		fmt.Printf("%d X %d = %d\n", a, i, a*i)
	}
}
