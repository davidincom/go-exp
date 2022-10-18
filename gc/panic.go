package gc

import "fmt"

func mayPanic() {
	panic("A Problem!")
}

func TestPanicRecover() {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from Error: ", r)
		}
	}()

	mayPanic()

	fmt.Println("After mayPanic")
}
