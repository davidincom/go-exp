// Good package names are short and clear.
// They are lower case, with no under_scores or mixedCaps. They are often simple nouns
package goroutines

import (
	"fmt"
	"time"
)

// A goroutine is a lightweight thread of execution.

func printString(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func TestRoutines() {
	printString("direct")

	go printString("goroutine")

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	time.Sleep(time.Second)
	fmt.Println("done")
}
