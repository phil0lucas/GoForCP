package main

import (
	"fmt"
	"flag"
	"strconv"
	"strings"
	
	"github.com/phil0lucas/GoForCP/DM"
	"github.com/phil0lucas/GoForCP/CPUtils"
	"github.com/jung-kurt/gofpdf"	
)

var infile = flag.String("i", "../CreateData/dm3.csv", "Name of input file")
var outfile = flag.String("o", "listing03.pdf", "Name of output file")

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
		failures + " were excluded at Screening and are not counted."
	f := &footers{
		foot1Left	:	"Created with Go 1.8 for linux/amd64.",
		foot2Left	:	f2,
		foot3Left	:	"All measurements were taken at the screening visit.",
		foot4Left	:	"Page %d of {nb}",
		foot4Right	:	"Run: " + CPUtils.TimeStamp(),
		foot4Centre	:	CPUtils.GetCurrentProgram(),
	}
	return f
}

//	Usubjid is displayed with leading studyid removed in
//	the style siteid-subjid
func SiteSubj (usubjid string) string {
	sl := strings.Split(usubjid,"-")
	return strings.Join(sl[1:],"-")
}	

func main() {
// 	Read the input file into a struct of values
	dm := DM.ReadDM(infile)
	
// 	Remove SF
	dm2 := DM.RemoveSF(dm)
	
//	Count by treatment group
	nTG := DM.CountByTG(dm)

// 	Determine the unique treatment group (Arm) values
	TGlist := DM.UniqueTG(dm2)
		
	pdf := gofpdf.New("L", "mm", "A4", "")
	h := titles()
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
	
	f_scr := strconv.Itoa(nTG["Screened"])
	f_sf := strconv.Itoa(nTG["SF"])
	f := footnotes(f_scr, f_sf)
	pdf.SetFooterFunc(func() {
		pdf.SetY(-30)
		pdf.SetFont("Courier", "", 10)
		pdf.CellFormat(0, 10, (*f).foot1Left, "T", 0, "L", false, 0, "")
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

	colHeaderSlice := []string{"SiteID-SubjectID", "Date of Birth", "Age (Years)", "Gender", "Ethnicity"}
	colWidthSlice := []float64{40, 40, 40, 40, 40}
	colJustSlice := []string{"L", "L", "L", "L", "L"}
// 	for i, str := range colHeaderSlice {
// 		pdf.CellFormat(colWidthSlice[i], 8, str, "TB", 0, colJustSlice[i], false, 0, "")
// 	}
	pdf.Ln(8)	
	for _, v := range TGlist {
		subDM := DM.SubsetByArm(dm2, v)
// 		fmt.Println(v)
// 		fmt.Println(subDM)
		pdf.CellFormat(0, 8, v, "", 0, "L", false, 0, "")
		pdf.Ln(8)
		for i, str := range colHeaderSlice {
			pdf.CellFormat(colWidthSlice[i], 8, str, "TB", 0, colJustSlice[i], false, 0, "")
		}
		pdf.Ln(8)
		for _, dd := range subDM {
			pdf.CellFormat(40, 8, SiteSubj(dd.Usubjid), "", 0, "L", false, 0, "")
			pdf.CellFormat(40, 8, CPUtils.DateP2Str(dd.Brthdtc), "", 0, "L", false, 0, "")
			pdf.CellFormat(40, 8, CPUtils.IntP2Str(dd.Age), "", 0, "L", false, 0, "")
			pdf.CellFormat(40, 8, CPUtils.StrP2Str(dd.Sex), "", 0, "L", false, 0, "")
			pdf.CellFormat(40, 8, CPUtils.StrP2Str(dd.Race), "", 0, "L", false, 0, "")
			pdf.Ln(4)
			if pdf.GetY() > float64(160) {
				pdf.AddPage()
				pdf.CellFormat(0, 8, v, "", 0, "L", false, 0, "")
				pdf.Ln(8)
				for i, str := range colHeaderSlice {
					pdf.CellFormat(colWidthSlice[i], 8, str, "TB", 0, colJustSlice[i], false, 0, "")
				}
				pdf.Ln(8)
			}
		}
		if v != TGlist[len(TGlist)-1] {
			pdf.AddPage()
		}
	}
	
	// 	Output
	err := pdf.OutputFileAndClose(*outfile)
	fmt.Println(err)
}
