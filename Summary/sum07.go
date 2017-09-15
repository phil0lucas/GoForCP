package main

import (
	"fmt"
	"github.com/phil0lucas/GoForCP/Summary/SummaryReport"
	"github.com/jung-kurt/gofpdf"
	"flag"
	"time"
	"log"
	"os"
	"strconv"
// 	"path/filepath"
// 	"strings"
)

var infile = flag.String("i", "../DM/dm.csv", "Name of input file")
var outfile = flag.String("o", "summary07.pdf", "Name of output file")

type headers struct {
	head1Left	string
	head1Right	string
	head2Left	string
	head2Right	string
	head3Left	string
	head4Centre	string
	head5Centre	string
	head6Centre	string	
}

type footers struct {
	foot1Left	string
	foot2Left	string
	foot3Left	string
	foot4Left	string
	foot4Centre	string
	foot4Right	string
}

func titles() *headers{
	h := &headers{
		head1Left	: 	"Acme Corp",
		head1Right	: 	"CONFIDENTIAL",
		head2Left	: 	"XYZ123 / Anti-Hypertensive",
		head2Right	:	"Draft",
		head3Left	:	"Protocol XYZ123",
		head4Centre	:	"Study XYZ123",
		head5Centre	:	"Summary of Demographic Data by Treatment Arm",
		head6Centre	:	"All Subjects",
	}
	return h
}

func footnotes() *footers{
	f := &footers{
		foot1Left	:	"A long explanatory text",
		foot2Left	:	"All subjects are included, including the screening failures",
		foot3Left	:	"All measurements were taken at the screening visit",
		foot4Left	:	"Page %d of {nb}",
		foot4Right	:	"Run: " + timeStamp(),
		foot4Centre	:	getCurrentProgram(),
	}
	return f
}
 
func timeStamp () string {
	t := time.Now()
	return t.Format("2006-01-02 15:04:05")	
}

func getCurrentProgram () string {
	ex, err := os.Executable()
    if err != nil { log.Fatal(err) }
	return ex + ".go"
}

// func columnHeaders() {
// 
// }
	
func WriteReport(outputFile *string, h *headers, f *footers, nTG map[string]int) error {					
	pdf := gofpdf.New("L", "mm", "A4", "")
	fmt.Printf("%T\n", pdf)
	pdf.SetHeaderFunc(func() {
		pdf.SetFont("Courier", "", 10)
		pdf.CellFormat(0, 10, (*h).head1Left, "0", 0, "L", false, 0, "")
		pdf.CellFormat(0, 10, (*h).head1Right, "0", 0, "R", false, 0, "")
		pdf.Ln(4)
		pdf.CellFormat(0, 10, (*h).head2Left, "0", 0, "L", false, 0, "")
		pdf.CellFormat(0, 10, (*h).head2Right, "0", 0, "R", false, 0, "")
		pdf.Ln(4)
		pdf.CellFormat(0, 10, (*h).head3Left, "0", 0, "L", false, 0, "")
		pdf.Ln(4)
		pdf.CellFormat(0, 10, (*h).head4Centre, "0", 0, "C", false, 0, "")
		pdf.Ln(4)
		pdf.CellFormat(0, 10, (*h).head5Centre, "0", 0, "C", false, 0, "")
		pdf.Ln(4)
		pdf.CellFormat(0, 10, (*h).head6Centre, "0", 0, "C", false, 0, "")
		pdf.Ln(10)		
	})
	
	pdf.SetFooterFunc(func() {
		pdf.SetY(-30)
		pdf.SetFont("Courier", "", 10)
		pdf.CellFormat(0, 10, (*f).foot1Left, "0", 0, "L", false, 0, "")
		pdf.Ln(4)
		pdf.CellFormat(0, 10, (*f).foot2Left, "0", 0, "L", false, 0, "")
		pdf.Ln(4)
		pdf.CellFormat(0, 10, (*f).foot3Left, "0", 0, "L", false, 0, "")
		pdf.Ln(4)		
		pdf.CellFormat(0, 10, fmt.Sprintf((*f).foot4Left, pdf.PageNo()), "", 0, "L", false, 0, "")
		pdf.SetX(40)
		pdf.CellFormat(0, 10, (*f).foot4Centre, "", 0, "L", false, 0, "")
		pdf.CellFormat(0, 10, (*f).foot4Right, "", 0, "R", false, 0, "")
	})
	pdf.AliasNbPages("")
	
// 	AddPage() executes the generated Header and Footer functions
	pdf.AddPage()
	
// 	Column headers
	colHeaderSlice := []string{"Characteristic", "Statistic", "Placebo", "Active", "Overall"}
	colWidthSlice := []float64{60, 60, 50, 50, 50}
	colJustSlice := []string{"L", "L", "L", "L", "L"}
	for i, str := range colHeaderSlice {
		pdf.CellFormat(colWidthSlice[i], 8, str, "TB", 0, colJustSlice[i], false, 0, "")
	}
	pdf.Ln(8)
	
//	Number of Subjects By TG
	textSlice := []string{"Number of Subjects", "N", 
		strconv.Itoa(nTG["Placebo"]), 
		strconv.Itoa(nTG["Active"]),
		strconv.Itoa(nTG["Overall"])}
	for i, str := range textSlice {
		pdf.CellFormat(colWidthSlice[i], 8, str, "", 0, colJustSlice[i], false, 0, "")
	}
	pdf.Ln(8)	
	
	
	
	
	
// 	Output
	err := pdf.OutputFileAndClose(*outputFile)
	fmt.Println(err)
	return err
} 

func countByTG (dm []*SummaryReport.DMrec) map[string]int {
	m := make(map[string]int)
	for _, v := range dm {
		m[v.Arm]++
	}
	
	total := 0
	for _, v := range m{
		total += v
	}
	m["Overall"] = total
	return m
}

func main() {
	// Read the file and dump into the slice of structs
	dm := SummaryReport.ReadFile(infile)
// 	fmt.Printf("%T\n", dm)
// 	for _, v := range dm {
// 		fmt.Println(*dm[0])
// 	}

// 	Compute number of subjects by treatment group
	nTG := countByTG(dm)
	fmt.Println(nTG)
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
	f := footnotes()
	err := WriteReport(outfile, h, f, nTG)
	if err != nil {
		fmt.Println(err)
	}
}
