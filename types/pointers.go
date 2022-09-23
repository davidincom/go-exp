package types

import "fmt"

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}

func TestPointers() {
	i := 1
	fmt.Println("initial:", i)
	zeroval(i)
	fmt.Println("zeroval:", i)

	zeroptr(&i)
	fmt.Println("zeroptr:", i)
	fmt.Println("pointer:", &i)

	myNumber := 23
	var ptr = &myNumber

	fmt.Println("pointer value is", ptr, myNumber)

	var val = *ptr

	*ptr = *ptr + 2

	fmt.Println("pointer value is", ptr, myNumber)

	var nilPtr *int
	// skip warning
	nilPtr = &myNumber

	fmt.Println("pointer value is", ptr, "actual value is", val, nilPtr)
}
