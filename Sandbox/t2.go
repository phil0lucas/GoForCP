package main

import "fmt"


func main() {

    // Declare variable of type int with a value of 10.
    count := 10

    // Display the "value of" and "address of" count.
    fmt.Println("count:\tValue Of[", count, "]\tAddr Of[", &count, "]")

    // Pass the "value of" the count.
    increment(&count)

    fmt.Println("count:\tValue Of[", count, "]\tAddr Of[", &count, "]")
}

func increment(inc *int) {
    // Increment the "value of" inc.
    *inc++
    fmt.Println("inc:\tValue Of[", inc, "]\tAddr Of[", &inc, "]")
}
