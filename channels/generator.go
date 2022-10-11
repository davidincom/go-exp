package channels

import (
	"fmt"
	"math/rand"
	"time"
)

// generator: function that returns a channel
func boring(msg string) <-chan string {
	c := make(chan string)

	go func() {
		for i := 0; ; i++ {
			gen_msg := fmt.Sprintf("%s %d", msg, i)
			c <- gen_msg
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()

	return c
}

// multiplexing: use a fan-in function to let whosoever is ready to talk
func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)

	go func() {
		for {
			c <- <-input1
		}
	}()

	go func() {
		for {
			c <- <-input2
		}
	}()

	return c
}

func fanInSelect(input1, input2 <-chan string) <-chan string {
	c := make(chan string)

	go func() {
		for {
			select {
			case s := <-input1:
				c <- s
			case s := <-input2:
				c <- s
			}
		}
	}()
	return c
}

func TestBoring() {
	joe := boring("Joe!")
	ann := boring("Ann")

	for i := 0; i < 5; i++ {
		fmt.Printf("You say: %q\n", <-joe)
		fmt.Printf("You say: %q\n", <-ann)
	}

	c := fanInSelect(joe, ann)

	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}

	s := fanIn(joe, ann)
	timeout := time.After(5 * time.Second)
	for {
		select {
		case s := <-s:
			fmt.Println("Ticking: ", s)
		case <-timeout:
			fmt.Println("You talk too much")
			return
		}
	}

	// fmt.Println("Leaving!")
}
