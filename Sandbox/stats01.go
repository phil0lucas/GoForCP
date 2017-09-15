package main

import (
	"fmt"
	"github.com/montanaflynn/stats"
)

func main() {
	d := stats.LoadRawData([]interface{}{1.1, "2", 3.0, 4, "5"})
	fmt.Println(stats.Mean(d))
}
	
