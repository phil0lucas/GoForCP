// Test sorting of a slice with single key

package main

import (
	"fmt"
	"sort"
)

type tmap struct{
	Code int
	Test string
}


func main() {
	s := []tmap{{1,"SBP"}, {2,"DBP"}, {3,"HR"}}
	
	sort.Slice(s, func(i, j int) bool { return s[i].Code < s[j].Code })
	fmt.Println("By Code:", s)

	sort.Slice(s, func(i, j int) bool { return s[i].Test < s[j].Test })
	fmt.Println("By Test:", s)
}