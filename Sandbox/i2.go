// Fails - how you can't always take the address of a value in Go
package main

import "fmt"

type notifier interface {
	notify()
}

type duration int

func (d *duration) notify() {
	fmt.Println("Sending Notification in", *d)
}

func main() {
	duration(42).notify()
}
