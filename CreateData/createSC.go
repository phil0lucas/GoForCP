package main

import (
	"flag"
	"github.com/phil0lucas/GoForCP/SC"
)
var outfile = flag.String("o", "sc3.csv", "Name of output file")

func main() {
	flag.Parse()
	SC.WriteSC(outfile)
}
