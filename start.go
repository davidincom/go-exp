package main

import (
	"fmt"
	"goExp/channels"
	"goExp/goroutines"
	"goExp/types"
)

func main() {
	fmt.Println("Hello World!")
	types.TestAllTypes()
	goroutines.TestGoRoutines()
	channels.TestAllChannels()
}
