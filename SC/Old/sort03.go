// Test sorting of a slice with single key

package main

import (
	"fmt"
	"sort"
)

func sortmap(m map[int]string) []string {
	var keys []int
	var values []string
	for k := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, k := range keys{
		values = append(values, m[k])
	}
	return values
}

// The numeric key provides the order of the resulting slice of strings.
func main() {
	s := map[int]string{2:"SBP", 3:"DBP", 1:"HR"}
	v := sortmap(s)
	fmt.Println(v)
}