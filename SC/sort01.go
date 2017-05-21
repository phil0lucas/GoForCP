// Test sorting of a slice with single key

// Could do this with sort.Ints(s)

package main

import (
	"fmt"
	"sort"
)

type simpleInts []int

func (a simpleInts) Len() int { 
	return len(a) 
}

func (a simpleInts) Swap(i, j int) {
	a[i], a[j] = a[j], a[i] 
}

func (a simpleInts) Less(i, j int) bool {
	return a[i] < a[j]
}

func main() {
	s := []int{104, 23, 12, 354, 124, 106}
	sort.Sort(simpleInts(s))
	fmt.Println(s)
}