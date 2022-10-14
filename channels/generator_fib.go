package channels

import "fmt"

func fibGenerator(n int) chan int {
	c := make(chan int)

	go func() {
		a, b := 0, 1
		for i := 0; i < n; i++ {
			a, b = b, a+b
			c <- a
		}
		close(c)
	}()

	return c
}

func TestFibGenerator() {
	for x := range fibGenerator(50) {
		fmt.Println(x)
	}
}
