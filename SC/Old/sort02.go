// Test sorting of a slice with single key

package main

import (
	"fmt"
	"sort"
)

type custom []string

func (a custom) Len() int { 
	return len(a) 
}

func (a custom) Swap(i, j int) {
	a[i], a[j] = a[j], a[i] 
}

func (a custom) Less(i, j int) bool {
	return a[i] < a[j]
}

func main() {
	s := []string{"SBP", "DBP", "HR"}
	sort.Sort(custom(s))
	fmt.Println(s)
}