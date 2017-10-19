// Test sorting of a slice with multiple keys

package main

import (
	"fmt"
	"sort"
)

type tmap struct{
	Code int
	Test string
}

type testCodes []tmap

func (t testCodes) Len() int {
	return len(t)
}

func (t testCodes) Swap(i, j int) {
	t[i], t[j] = t[j], t[i] 
}

func (t testCodes) Less(i, j int) bool {
	if t[i].Code < t[j].Code {
		return true
	}
	if t[i].Code > t[j].Code {
		return false
	}	
	// If Codes are equal
	return t[i].Test < t[j].Test
}

func main() {
	s := testCodes{
		{1,"SBP"}, 
		{2,"DBP"}, 
		{3,"HR"},
		{3,"SO2"},
	}
	sort.Sort(s)
	fmt.Println(s)
}