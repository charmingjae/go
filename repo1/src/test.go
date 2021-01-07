// 패키지 지정
package main

import "fmt"

// 메인 함수
func main() {
	chk(2)
}

func chk(val int) {
	switch val {
	case 1:
		fmt.Println("1 이하")
		fallthrough
	case 2:
		fmt.Println("2 이하")
		fallthrough
	case 3:
		fmt.Println("3 이하")
		fallthrough
	default:
		fmt.Println("default")
	}
}
