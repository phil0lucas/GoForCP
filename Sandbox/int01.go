package main

import "fmt"

type number interface{
	IsOdd () bool
	Median () float64
}

type BP []float64
type Age []int

var f = BP{1.2, 3.4, 8.9}
var f2 = BP{1.2, 3.4, 8.9, 11.3}
var i = Age{1, 3, 8}
var i2 = Age{0, 3, 8, 11}

func (i Age) IsOdd() bool {
	if len(i) % 2 == 0 {
		return false
	} else {
		return true
	}
}

func (f BP) IsOdd() bool {
	if len(f) % 2 == 0 {
		return false
	} else {
		return true
	}
}

func (i Age) Median() float64 {
	if i.IsOdd() {
		pos := (len(i) / 2) - 1
		return float64(i[pos])
	} else {
		return 1.0
	}
}

func main() {
	fmt.Println(f)
	fmt.Println(f.IsOdd())
	fmt.Println(f2.IsOdd())
	fmt.Println(i.IsOdd())
	fmt.Println(i2.IsOdd())
	
	fmt.Println(i.Median())
}
