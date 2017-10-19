package main

import (
	"flag"
	"github.com/phil0lucas/GoForCP/DM"
)

// The program will be run with flags to specify the input & output files
var infile = flag.String("i", "sc3.csv", "Name of input file")
var outfile = flag.String("o", "dm3.csv", "Name of output file")

func main() {
	flag.Parse()
	DM.WriteDM(infile, outfile)
}
