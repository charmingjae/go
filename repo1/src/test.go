package main

import "fmt"

func main() {

	var num int = 2

	for i := 1; i <= 9; i++ {
		fmt.Printf("%d X %d = %d\n", num, i, num*i)
	}
}
