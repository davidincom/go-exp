package goroutines

import (
	"fmt"
	"sync"
	"time"
)

func worker2(id int) {
	fmt.Printf("worker %d starting \n", id)
	time.Sleep(time.Second)
	fmt.Printf("worker %d done \n", id)
}

/*
A WaitGroup waits for a collection of goroutines to finish.
The main goroutine calls Add to set the number of goroutines to wait for.
Then each of the goroutines runs and calls Done when finished.
At the same time, Wait can be used to block until all goroutines have finished.
*/
func TestWaitGroup() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)

		i := i

		go func() {
			defer wg.Done()
			worker2(i)
		}()
	}

	wg.Wait()
}
