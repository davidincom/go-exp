package channels

func Generate(ch chan<- int) {
	for i := 2; ; i++ {
		print("0 send value to channel: ", i, "\n")
		ch <- i
	}
}

func Filter(in <-chan int, out chan<- int, prime int, no int) {
	for {
		i := <-in
		print(no, " prime number to check: ", i, " seive ", prime, "\n")
		if i%prime != 0 {
			out <- i
			print(no, " channel received: ", i, " seive ", prime, "\n")
		}
	}
}

func TestConcPrimeSieve() {
	ch := make(chan int)
	go Generate(ch)

	for i := 0; i < 5; i++ {
		prime := <-ch
		print(prime, "\n")
		ch1 := make(chan int)
		go Filter(ch, ch1, prime, i)
		ch = ch1
		// time.Sleep(2 * time.Second)
	}
}
