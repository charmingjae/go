// 패키지 지정
package main

// 메인 함수
func main() {
	println("Hello World!")

	var i, j, k int = 1, 2, 3
	println(i, j, k)

	const (
		Visa   = "Visa"
		Master = "MasterCard"
		Amex   = "American Express"
	)

	println(Visa, Master, Amex)

	const (
		a = iota
		b
		c
	)

	println(a, b, c)
}
