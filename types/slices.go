package types

import (
	"fmt"

	"golang.org/x/exp/slices"
)

func TestSlice() {

	s := make([]string, 3)
	fmt.Println("emp:", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	has := slices.Contains(s, "b")
	fmt.Println("slices contains b: ", has)

	fmt.Println("len:", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	l := s[2:5]
	fmt.Println("sl1:", l)

	l = s[:5]
	fmt.Println("sl2:", l)

	l = s[2:]
	fmt.Println("sl3:", l)

	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}

func RemoveDuplicate[T string | int](sliceList []T) []T {
	allkeys := make(map[T]bool)
	list := []T{}
	for _, item := range sliceList {
		if _, value := allkeys[item]; !value {
			allkeys[item] = true
			list = append(list, item)
		}
	}
	return list
}
