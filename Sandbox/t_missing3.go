package main

import (
    "fmt"
    "strconv"
    "errors"
)

type Age *int

func addr (s string) (Age, error) {
    i, err := strconv.Atoi(s)
    if err == nil {
        return &i, nil
    } else {
        return nil, errors.New("Missing Value")
    }
}

func main() {
    i, j := "12", ""
    
    a, err := addr(i)
    if err != nil {
        fmt.Println(err)
    }
    b, err := addr(j)
    if err != nil {
        fmt.Println(err)
    }    
    fmt.Println(a, b)
}
