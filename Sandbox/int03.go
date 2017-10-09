// Calculate simple descriptive stats for slice of ints or floats.
// Example of using an interface to specify multiple types,
package main

import (
	"fmt"
	"time"
)

type date interface {
	PrintDate()
}

type pdate struct {
	pd *time.Time
}

type ndate struct {
	nd time.Time
}

func (p pdate) PrintDate() {
	if p != nil{
		fmt.Println(p.Format("2006-01-02"))
	} else {
		fmt.Println("Missing Date")
	}
}

func (n ndate) PrintDate() {
	fmt.Println(n.Format("2006-01-02"))
}

func main() {
	gSlice := []gender{"M", "F", "F", "M", "M"}
	s := gSlice.uniqueValues()
	fmt.Println(s)
}
