// Calculate simple descriptive stats for slice of ints or floats.
// Example of using an interface to specify multiple types,
package main

import (
	"fmt"
)

type gender string
// type race string

// var rSlice []race{"White", "Asian", "White", "Black", "Asian"}

func stringInSlice(a string, list []string) bool {
    for _, b := range list {
        if b == a {
            return true
        }
    }
    return false
}

func (g []gender) uniqueValues () []string {
	var uValues []string
	for _, v := range g {
		if !stringInSlice (v , uValues) {
			uValues = append(uValues, v)
		}
	}
	return uValues
}

func main() {
	gSlice := []gender{"M", "F", "F", "M", "M"}
	s := gSlice.uniqueValues()
	fmt.Println(s)
}
