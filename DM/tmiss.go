package main

import (
	"fmt"
	"math/rand"
	"time"
)

func flagMiss () bool {
	rand.Seed(time.Now().UTC().UnixNano())
	if rand.Float64() >= 0.05 {
		return false
	} else {
		return true
	}
}

func main() {
	for ii := 0; ii < 100; ii++{
		fmt.Println (flagMiss())
	}
}
