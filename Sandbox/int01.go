// Calculate median for slice of ints
package main

import (
	"fmt"
	"math"
)

type Age []int

var i = Age{0,3,8,11,12,56,34}
var i2 = Age{0, 3, 8, 11}

func (i Age) Median() float64 {
	if l := len(i); l > 0 {
		if l % 2 == 0 {
			fmt.Println("Even")
			lowerBound := (l/2) - 1
			upperBound := l/2
			fmt.Println("lower index value", i[lowerBound])
			fmt.Println("upper index value", i[upperBound])
			fmt.Printf("%T\n", i[lowerBound])
			mid := float64((int(i[lowerBound]) + int(i[upperBound])) / 2)
			fmt.Printf("%f", mid)
			//for pp := range midTwo{
			//	fmt.Println(pp)
			//}
			return 0.0
		} else {
			fmt.Println("Odd")
			mid := int(math.Floor(float64(l / 2)))
			return float64(i[mid])
		}
	// Empty slice
	} else {
		return math.NaN() // ???
	}
}


func main() {
	//fmt.Println(i.Median())
	fmt.Println(i2.Median())
}
