package main

import "fmt"
import "unicode"


type Number interface {
	int64 | float64
}

func main(){
	ints := map[string]int64 {
		"first": 34,
		"second": 12,
	}
	floats := map[string]float64 {
		"first": 34.2,
		"second": 1.2,
	}
	
	fmt.Printf("sums %v and %v %v", SumInts(ints), SumFLoats(floats), SumIntsAndFLoats(ints))
	fmt.Printf(ToUpper("Hola me amour"))
}




func SumInts(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}
	return s
}

func SumFLoats( m map[string]float64) float64 {
	var s float64
	for _, v := range m {
		s += v
	}
	return s
}



func SumIntsAndFLoats[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

func ToUpper(s string) (string) {

	var r = []rune(s)
	for index := range r {
		r[index] = unicode.ToUpper(r[index])

	}
	return string(r)
}