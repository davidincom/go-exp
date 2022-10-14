package main

import (
	"fmt"
	"goExp/channels"
	"goExp/gc"
	"goExp/goroutines"
	"goExp/types"
)

type Duck struct {
}

func NewDock() *Duck {
	return &Duck{}
}

func main() {
	fmt.Println("Hello World!")
	types.TestAllTypes()
	goroutines.TestGoRoutines()
	channels.TestAllChannels()
	// containers.TestMinHeap()
	// sort.TestMinHeapSort()

	// NewDock()
	gc.TestAllGC()
}
