package generics

import "fmt"

type Number interface {
	int64 | float64
}

func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

func SumNumber[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

func TestGenericSum() {
	floats := map[string]float64{
		"first":  35.98,
		"second": 26.99,
	}

	sum := SumIntsOrFloats(floats)
	// sum := SumIntsOrFloats[string, float64](floats)

	fmt.Println("Generic sum for floats: ", sum)

	sum = SumNumber(floats)
	// sum := SumIntsOrFloats[string, float64](floats)

	fmt.Println("Generic sum for numbers: ", sum)
}
