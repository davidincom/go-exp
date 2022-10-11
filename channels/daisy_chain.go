package channels

import "fmt"

func f(left, right chan int) {
	left <- 1 + <-right
}

func create_chain(n int) int {
	leftmost := make(chan int)
	right := leftmost
	left := leftmost

	for i := 0; i < n; i++ {
		right = make(chan int)
		go f(left, right)
		left = right
	}

	go func(c chan int) {
		c <- 1
	}(right)

	return <-leftmost
}

func TestDaisyChain() {
	const n = 100000
	result := create_chain(n)

	fmt.Println(result)
}
