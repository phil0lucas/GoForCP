package main

import "fmt"

type notifier interface {
	notify()
}

type duration int

func (d duration) notify() {
	fmt.Println("Sending Notification in", d)
}

func main() {
	var n notifier
	n = duration(42)
	n.notify()
}
