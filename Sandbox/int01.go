// Calculate simple descriptive stats for slice of ints or floats.
// Example of using an interface to specify multiple types,
package main

import (
	"fmt"
	"math"
)

type numeric interface {
	Median() float64
	Mean() float64
}

type Age []int
type Temperature []float64

func (i Age) Median() float64 {
	if l := len(i); l > 0 {
		if l % 2 == 0 {
			fmt.Println("Even")
			lowerBound := (l/2) - 1
			upperBound := l/2
			return float64(i[lowerBound] + i[upperBound]) / 2
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

func (i Temperature) Median() float64 {
	if l := len(i); l > 0 {
		if l % 2 == 0 {
			fmt.Println("Even")
			lowerBound := (l/2) - 1
			upperBound := l/2
			return (i[lowerBound] + i[upperBound]) / 2
		} else {
			fmt.Println("Odd")
			mid := int(math.Floor(float64(l / 2)))
			return i[mid]
		}
	// Empty slice
	} else {
		return math.NaN() // ???
	}
}

func (a Age) Mean() float64 {
	fmt.Println("Mean of slice of Ints")
	var sum int
	for _, v := range a {
		sum += v
	}
	return float64(sum / len(a))
}

func (t Temperature) Mean() float64 {
	fmt.Println("Mean of slice of Floats")
	var sum float64
	for _, v := range t {
		sum += v
	}
	return sum / float64(len(t))
}


func printStats(n numeric) {
	fmt.Println(n.Median())
	fmt.Println(n.Mean())	
}


func main() {
	var i = Age{0,3,8,11,12,56,34}
	var i2 = Age{0, 3, 8, 11}
	var f = Temperature{35.3, 36.2, 36.7, 38.1, 38.9}
	var f2 = Temperature{36.5, 36.6, 36.8, 37.0}
	
	printStats(i)
	printStats(i2)
	printStats(f)
	printStats(f2)	
}
