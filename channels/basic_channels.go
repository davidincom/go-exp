package channels

import (
	"fmt"
)

/*
Channels are the pipes that connect concurrent goroutines.
You can send values into channels from one goroutine
and receive those values into another goroutine.

These channels are specially designed to prevent race conditions
when accessing shared memory using Goroutines.

unbuffered Channels:

	meaning that they will only accept sends (chan <-)
	if there is a corresponding receive (<- chan) ready to receive the sent value.
*/
func TestUnbufferedChannels() {
	messages := make(chan string)

	go func() {
		messages <- "ping"
	}()

	msg := <-messages

	fmt.Println(msg)
}

// Buffered channels accept a limited number of values without a corresponding receiver for those values.
// If zero, or the size is omitted, the channel is unbuffered.
// When the channel is buffered, the send in the goroutine is nonblocking.
// This is a common pattern to prevent goroutine leaks in case the channel is never read.
func TestBufferedChannels() {
	messages := make(chan string, 2)

	messages <- "buffered"
	messages <- "channel"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}

func TestChannels() {
	TestUnbufferedChannels()
	TestBufferedChannels()
}
