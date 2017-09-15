//

package main

import (
	"fmt"
	"github.com/phil0lucas/GoForCP/Summary/SummaryReport"
	//	"bufio"
	"flag"
	//	"strings"
	//	"os"
	//	"strconv"
)

var infile = flag.String("i", "../DM/dm.csv", "Name of input file")
var outfile = flag.String("o", "summary05.pdf", "Name of output file")

// Slice of pointers to structs -
//	one element per input record
//type dmrecs []*SummaryReport.DMrec

func main() {
	//var dm dmrecs
	// Read the file and dump into the slice of structs
	dm := SummaryReport.ReadFile(infile)
	for _, v := range dm {
		fmt.Println(*v)
	}
	err := SummaryReport.WriteReport(outfile, "Acme Corp", "XYZ123 / Anti-Hypertensive",
		"Protocol XYZ123", "Study XYZ123",
		"Summary of Demographic Data by Treatment Arm", "All Subjects")
	fmt.Println(err)
}
