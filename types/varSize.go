package types

import (
	"fmt"
	"math/bits"
	"reflect"
	"unsafe"
)

func TestSizeAndTypes() {
	sizeOfIntInBits := bits.UintSize
	fmt.Println(sizeOfIntInBits, "bits") // 64 bits

	var a int
	fmt.Println("int is", unsafe.Sizeof(a), "bytes") // 8 bytes
	fmt.Println("a's type is", reflect.TypeOf(a))

	b := 2                                        // default is int, int64 if on 64bit system
	f := 2.3                                      // default is float64 when no type is specified
	fmt.Println("b's type is", reflect.TypeOf(b)) // int
	fmt.Println("f's type is", reflect.TypeOf(f)) // float64

	tOff := reflect.TypeOf(f)
	size := tOff.Size()
	fmt.Println("size of", tOff, "is", size) // size of float64 is 8
}
