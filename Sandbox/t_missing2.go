package main

import (
    "fmt"
    "strconv"
)

type Line struct {
    One   *int
    Two   *int
    Three *int
}

func NewLine(s []string) (l *Line) {
    fmt.Println(len(s))
    if len(s) < 2 {
        return
    }
    l = &Line{}
    if i, err := strconv.Atoi(s[0]); err == nil {
        l.One = &i
    }
    if i, err := strconv.Atoi(s[1]); err == nil {
        l.Two = &i
    }
    if len(s) == 3 {
        if i, err := strconv.Atoi(s[2]); err == nil {
            l.Three = &i
        }
    }
    return
}

func main() {
    a := []string {"1","","3"}
    fmt.Println(NewLine(a))
}
