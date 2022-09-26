package channels

import (
	"fmt"
	"math/rand"
	"time"
)

func longTimeRequest() <-chan int32 {
	r := make(chan int32)

	go func() {
		time.Sleep(time.Second * 3)
		r <- rand.Int31n(100)
	}()

	return r
}

func sumSquares(a, b int32) int32 {
	return a*a + b*b
}

func PromiseTest() {
	rand.Seed(time.Now().UnixNano())

	a, b := longTimeRequest(), longTimeRequest()

	ra, rb := <-a, <-b

	result := sumSquares(ra, rb)

	fmt.Println(result)
}
