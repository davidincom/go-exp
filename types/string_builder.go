package types

import (
	"fmt"
	"strings"
)

// A Builder is used to efficiently build a string using Write methods.
// It minimizes memory copying. The zero value is ready to use.
func TestStringBuilder() {
	var str strings.Builder

	fmt.Println("inital value of string builder: ", str.String())

	for i := 0; i < 100; i++ {
		str.WriteString("a")
	}

	fmt.Println(str.String())
}
