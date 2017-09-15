//

package main

import (
	"fmt"
	"github.com/phil0lucas/GoForCP/Summary/SummaryReport"
	"github.com/jung-kurt/gofpdf"
	"flag"
// 	"os"
)

var infile = flag.String("i", "../DM/dm.csv", "Name of input file")
var outfile = flag.String("o", "summary06.pdf", "Name of output file")

type headers struct {
	line1a	string
}

func titles() *headers{
	h := &headers{line1a : "Acme Corp"}
	return h
}

// 		"Acme Corp", "XYZ123 / Anti-Hypertensive",
// 		"Protocol XYZ123", "Study XYZ123",
// 		"Summary of Demographic Data by Treatment Arm", "All Subjects")

func main() {
// 	fmt.Println(os.Args[0])
	// Read the file and dump into the slice of structs
	dm := SummaryReport.ReadFile(infile)
// 	for _, v := range dm {
		fmt.Println(*dm[0])
// 	}
	
// 	Compute number of subjects by treatment group
		
// 	Turn values into strings
		
// 	Compute number of non-mising Age values by TG
		
// 	Turn values into strings
		
// 	Compute mean of age by TG and SD by TG
		
// 	Turn Mean and SD values into strings
		
// 	Compute median values of Age by TG
		
// 	Turn median values into strings
		
// 	Compute min values of Age by TG
		
// 	Turn min values into strings
		
// 	Compute max values of Age by TG
		
// 	Turn max values into strings
	
// 	New Report 
	
	h := titles()
	err := SummaryReport.WriteReport(outfile, h)

	fmt.Println(err)
// 	
// 	
// 	
// 	pdf.AddPage()
// // 	basicTable()
// 
// 	err := pdf.OutputFileAndClose(*outputFile)
// 	fmt.Println(err)
}
