package main

import "fmt"

func square(x *float64) *float64 {
    *x = *x * *x
    fmt.Printf("%T\n", x)
    return x
}
func main() {
  x := 1.5
  y := square(&x)
  fmt.Println(*y)
}
