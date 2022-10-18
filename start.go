package main

import (
	"fmt"

	"github.com/davidiwu/go-exp/channels"
	"github.com/davidiwu/go-exp/gc"
	"github.com/davidiwu/go-exp/goroutines"
	"github.com/davidiwu/go-exp/types"
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
