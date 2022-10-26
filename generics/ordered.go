package generics

import (
	"fmt"
	"strings"

	"golang.org/x/exp/constraints"
)

func maxGeneric[T constraints.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func anyGeneric[T any](sth T) {
	fmt.Println(sth)
}

func reverseArray[T any](values []T) []T {
	if len(values) == 0 {
		return values
	}
	length := len(values)
	for i, j := 0, length-1; i < j; i, j = i+1, j-1 {
		values[i], values[j] = values[j], values[i]
	}
	return values
}

func exists[T constraints.Ordered](values []T, value T) bool {
	for _, v := range values {
		if v == value {
			return true
		}
	}
	return false
}

func TestMaxGeneric() {
	a, b := 42, 49
	fmt.Println(maxGeneric(a, b))

	c, d := "abc", "def"
	fmt.Println(maxGeneric(c, d))

	anyGeneric(c)

	alongstr := "this is a long string"
	strs := strings.Split(alongstr, " ")
	has := exists(strs, "long")
	fmt.Println("long exists in strs: ", has)

	fmt.Printf("address of strs %p, value is %v \n", strs, strs)
	reversed := reverseArray(strs)
	fmt.Printf("address of reversed %p, value is %v \n", reversed, reversed)
	fmt.Println("reversed array: ", strings.Join(reversed, " "))
}
