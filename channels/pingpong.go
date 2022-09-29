package channels

import (
	"fmt"
	"os"
	"time"
)

type Ball uint64

func Play(playerName string, table chan Ball) {
	var lastValue Ball = 1
	for {
		fmt.Println("enter for loop.")
		ball := <-table
		fmt.Println(playerName, ball)
		ball += lastValue
		if ball < lastValue {
			os.Exit(0)
		}
		lastValue = ball
		table <- ball
		time.Sleep(time.Second)
	}
}

func TestPingPong() {
	table := make(chan Ball)
	go func() {
		fmt.Println("enter start.")
		table <- 1
	}()

	go Play("A:", table)
	Play("B:", table)
	fmt.Scanln()
}
