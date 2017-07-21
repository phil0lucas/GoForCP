package main

import "fmt"


func method() []int {
  var  slice []int 
  for i := 0; i < 10; i++  {
    m1 := map[string]int{}
    key := fmt.Sprintf("variable%d", i)
    m1[key] = i
    slice = append(slice, m1[key])
  }
  return slice
}

func main() {
    fmt.Println(method())
}