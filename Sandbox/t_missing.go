package main

import (
    "fmt"
    "strconv"
)

type Age *int

func addr (i int) Age {
    return &i
}

func main() {
    i, j := 12, 13
    
    a := addr(i)
    b := addr(j)
    
    fmt.Println("Ages = " + strconv.Itoa(*a) + ", " + strconv.Itoa(*b))
}
