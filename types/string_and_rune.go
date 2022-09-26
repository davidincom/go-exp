package types

import (
	"fmt"
	"unicode/utf8"
)

// a Go string is a read-only slice of bytes. = []byte
// a rune is an integer that represents a Unicode code point

/*
double quotes:
	It is used to define a string.
	A string defined within double quotes will honor escaping characters.
	x := "tit\nfor\ttat"

back quotes:
	It is also used to define a string.
	A string encoded in back quotes is a raw literal string
	and doesn’t honor any kind of escaping.
	y := `tit\nfor\ttat`

single quotes:
	To declare either a byte or a rune we use single quotes.
	While declaring byte we have to specify the type.
	If we don’t specify the type, then the default type is meant as a rune.
	A single quote will allow only one character.
	r := '£'
	var b byte = 'a'
*/

func TestStringRune() {
	const s = "你好\n中国"
	// const s = "hello world"
	fmt.Println("len: ", len(s))

	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Println()

	fmt.Println("Rune Count: ", utf8.RuneCountInString(s))

	for idx, runeValue := range s {
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
	}

	fmt.Println("\nUsing DecodedRuneInString")
	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at %d, with width= %d bytes\n", runeValue, i, width)
		w = width
	}

	var zhong rune = '中'
	fmt.Println(zhong)

	var zhongc byte = '$'
	fmt.Println(zhongc)
}
