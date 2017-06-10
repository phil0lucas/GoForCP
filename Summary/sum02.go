// 

package main

import (
//	"github.com/jung-kurt/gofpdf"
	"fmt"
	"bufio"
	"flag"
	"strings"
)

type dmrec struct {
	usubjid string
	age		int
	sex		string
	race	string
	armcd	int
	arm		string
}

// Slice of pointers to structs - 
//	one element per input record
type dmrecs []*dmrec

// The program will be run with flags to specify the input & output files
var infile = flag.String("i", "../DM/dm.csv", "Name of input file")
//var outfile = flag.String("o", "summary.pdf", "Name of output file")

func readFile(infile string) dm dmrecs {
	// open the file and pass it to a Scanner object
	file, err := os.Open(infile)
	if err != nil {
		panic(fmt.Sprintf("error opening %s: %v", infile, err))
	}
	defer file.Close()
	
	// Pass the opened file to a scanner
	scanner := bufio.NewScanner(file)

	for i := 0; scanner.Scan(); i++ {
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "error reading from file:", err)
			os.Exit(3)
		}
		str := scanner.Text()
		usubjid := strings.Split(str, ",")[3]
	}
	
}

func main() {

}


	/*
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, "Hello, world")
	err := pdf.OutputFileAndClose("hello.pdf")
	fmt.Println(err)
	*/