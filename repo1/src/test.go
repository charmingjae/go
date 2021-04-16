package main

func main() {
	ch := make(chan int, 2)

	ch <- 1
	ch <- 2

	close(ch)

	println(<-ch)
	println(<-ch)

	if a, success := <-ch; !success {
		println(a, success)
	}
}
