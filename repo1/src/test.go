package main

func main() {
	add := func(x int, y int) int {
		return x + y
	}

	r1 := calc(add, 10, 20)
	println(r1)

	r2 := calc(func(x int, y int) int { return x - y }, 10, 20)
	println(r2)
}

func calc(f func(int, int) int, x int, y int) int {
	result := f(x, y)
	return result
}
