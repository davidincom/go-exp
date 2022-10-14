package channels

import "fmt"

func prime(n int) bool {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func primeGenerator(n int) chan int {
	c := make(chan int)

	go func() {
		for i := 1; n > 0; i++ {
			if prime(i) {
				c <- i
				n--
			}
		}
		close(c)
	}()
	return c
}

func TestPrimeGenerator() {
	for p := range primeGenerator(100) {
		fmt.Println(p)
	}
}
