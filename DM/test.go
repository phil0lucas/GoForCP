package main

import (
	"fmt"
 	"math/rand"
	"time"
)

func main () {
	m := map[int]string{
		0: "AAA",
		1: "BBB",
		2: "CCC",
	}

	rand.Seed(time.Now().UTC().UnixNano())
	key := rand.Intn(len(m))
    fmt.Println(m[key])	
}