package main

import (
	"fmt"
	"github.com/phil0lucas/GoForCP/Listing/Listing"
	"github.com/phil0lucas/GoForCP/CPUtils"
	"github.com/jung-kurt/gofpdf"
	"flag"
// 	"time"
// 	"log"
// 	"os"
// 	"strconv"
// 	"path/filepath"
// 	"strings"
)

var infile = flag.String("i", "../DM/dm.csv", "Name of input file")
var outfile = flag.String("o", "listing02.pdf", "Name of output file")


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
		head5Centre	:	"Listing of Demographic Data by Treatment Arm",
		head6Centre	:	"All Randomised Subjects",
	}
	return h
}

func footnotes(screened string, failures string) *footers{
	f2 := "Of the original " + screened + " screened subjects, " + 
		failures + " were excluded at Screening and are not counted"
	f := &footers{
		foot1Left	:	"A long explanatory text",
		foot2Left	:	f2,
		foot3Left	:	"All measurements were taken at the screening visit",
		foot4Left	:	"Page %d of {nb}",
		foot4Right	:	"Run: " + CPUtils.TimeStamp(),
		foot4Centre	:	CPUtils.GetCurrentProgram(),
	}
	return f
}

func main() {
// 	Read the input file into a struct of values
	dm := Listing.ReadFile(infile)
	fmt.Printf("%T\n", dm)

// 	Determine the unique treatment group (Arm) values
	TGlist := CPUtils.UniqueTG(dm)
	fmt.Println(TGlist)
	
	pdf := gofpdf.New("L", "mm", "A4", "")
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
	pdf.AddPage()
// 	For each TG, get a subset of the data based on the Arm value
/*	
	for _, v := range TGlist {
		subDM := Listing.SubsetByArm(dm, v)
		fmt.Println(v)
		fmt.Println(subDM)
	}
*/

// 	Output
	err := pdf.OutputFileAndClose(*outputFile)
	fmt.Println(err)
}
