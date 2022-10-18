package main

import (
	"fmt"

	"github.com/davidiwu/goexp/channels"
	"github.com/davidiwu/goexp/gc"
	"github.com/davidiwu/goexp/goroutines"
	"github.com/davidiwu/goexp/types"
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
