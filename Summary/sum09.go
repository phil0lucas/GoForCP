package main

import (
	"fmt"
	"github.com/phil0lucas/GoForCP/Summary/SummaryReport"
	"github.com/jung-kurt/gofpdf"
	"github.com/montanaflynn/stats"
	"flag"
	"time"
	"log"
	"os"
	"strconv"
// 	"path/filepath"
// 	"strings"
)

var infile = flag.String("i", "../DM/dm.csv", "Name of input file")
var outfile = flag.String("o", "summary09.pdf", "Name of output file")

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

// 	Provides a map of keys for each TG and Overall, the toal number screened and the SFs
func countByTG (dm []*SummaryReport.DMrec) map[string]int {
	m := make(map[string]int)
	m["Screened"] = len(dm)
	for _, v := range dm {
		if v.Arm != "" {
			m[v.Arm]++
		} else {
			m["SF"]++
		}
	}
	
	total := 0
	for k, v := range m{
		if k != "SF" && k != "Screened" {
			total += v
		}
	}
	m["Overall"] = total
	return m
}

func selectTGs(m map[string]int) []string {
	var s []string
	for k, _ := range m{
		if k != "SF" && k != "Screened" {
			s = append(s, k)
		}
	}
	return s
}
	
func WriteReport(outputFile *string, h *headers, f *footers, 
				 nTG map[string]int, nAge map[string]int, 
				 meansd map[string]string,
				 median map[string]string,
				 min map[string]string,
				 max map[string]string) error {
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
	
//	Number of non-missing Ages By TG
	textSlice2 := []string{"Age (years)", "Number of Non-Missing", 
		strconv.Itoa(nAge["Placebo"]), 
		strconv.Itoa(nAge["Active"]),
		strconv.Itoa(nAge["Overall"])}
	for i, str := range textSlice2 {
		pdf.CellFormat(colWidthSlice[i], 8, str, "", 0, colJustSlice[i], false, 0, "")
	}
	pdf.Ln(4)
	
// 	Mean and Standard Deviation by TG
	textSlice3 := []string{" ", "Mean (SD)", 
		meansd["Placebo"], 
		meansd["Active"],
		meansd["Overall"]}
	for i, str := range textSlice3 {
		pdf.CellFormat(colWidthSlice[i], 8, str, "", 0, colJustSlice[i], false, 0, "")
	}
	pdf.Ln(4)	
	
//  Median
	textSlice4 := []string{" ", "Median", 
		median["Placebo"], 
		median["Active"],
		median["Overall"]}
	for i, str := range textSlice4 {
		pdf.CellFormat(colWidthSlice[i], 8, str, "", 0, colJustSlice[i], false, 0, "")
	}
	pdf.Ln(4)		
	
//  Minimum
	textSlice5 := []string{" ", "Minimum", 
		min["Placebo"], 
		min["Active"],
		min["Overall"]}
	for i, str := range textSlice5 {
		pdf.CellFormat(colWidthSlice[i], 8, str, "", 0, colJustSlice[i], false, 0, "")
	}
	pdf.Ln(4)			
	
//  Maximum
	textSlice6 := []string{" ", "Maximum", 
		max["Placebo"], 
		max["Active"],
		max["Overall"]}
	for i, str := range textSlice6 {
		pdf.CellFormat(colWidthSlice[i], 8, str, "", 0, colJustSlice[i], false, 0, "")
	}
	pdf.Ln(8)	
	
	
	
// 	Output
	err := pdf.OutputFileAndClose(*outputFile)
	fmt.Println(err)
	return err
} 

func removeSF(dm []*SummaryReport.DMrec) []*SummaryReport.DMrec {
	var dm2 []*SummaryReport.DMrec
	for _, v := range dm {
		// Exclude SFs
		if v.Arm != "" {
// 			fmt.Println(v)
			dm2 = append(dm2, v)
		}
	}
	return dm2
}


func nMiss (dm []*SummaryReport.DMrec) map[string]int {
	m := make(map[string]int)
// 	total := 0
	for _, v := range dm {
		if v.Age != nil {
			m[v.Arm]++
			m["Overall"]++
		} else {
			m["Missing"]++
		}
	}
	return m
}

func prepareData(dm []*SummaryReport.DMrec, tg []string) map[string][]float64{
	m := make(map[string][]float64)
	for _, s := range tg {
		var out []float64
		for _, v := range dm {
			if v.Age != nil {
				if s == v.Arm {
					out = append(out, float64(*v.Age))
				} else if s == "Overall" {
					out = append(out, float64(*v.Age))
				}
			}
		}
		m[s] = out
	}
	return m
}


func mStat (indata map[string][]float64, stat string, dec int) map[string]string {
	m := make(map[string]string)
	for i, v := range indata {
		var result float64
		if stat == "Mean" {
			result, _ = stats.Mean(v)
		} else if stat == "SD" {
			result, _ = stats.StandardDeviationPopulation(v)
		} else if stat == "Median" {
			result, _ = stats.Median(v)
		} else if stat == "Min" {
			result, _ = stats.Min(v)
		} else if stat == "Max" {
			result, _ = stats.Max(v)
		}		
		c_stat := strconv.FormatFloat(result, 'f', dec, 64)
		m[i] = c_stat
// 		fmt.Println(err)
	}
	return m
}

func main() {
	// Read the file and dump into the slice of structs
	dm := SummaryReport.ReadFile(infile)

// 	Compute number of subjects by treatment group
	nTG := countByTG(dm)
// 	fmt.Println(nTG)

	TGs := selectTGs(nTG)
// 	fmt.Println(TGs)
	
// Create version of dm without the SFs
	dm2 := removeSF(dm)

// 	Compute number of non-missing Age values by TG
	nAge := nMiss(dm2)
		
//	Prepare the data for passing to the stats functions
//	Select only the TGs to display
//	Remove the missing values.
	rMiss := prepareData(dm2, TGs)	
	
	
// 	Compute stats of age by TG
	mean := mStat(rMiss, "Mean", 2)
	sd := mStat(rMiss, "SD", 2)
	
//	Concatenate mean and SD values into a display string
	meansd := make(map[string]string)
	for k, _ := range mean {
		meansd[k] = mean[k] + " (" + sd[k] + ")"
	}
	median := mStat(rMiss, "Median", 0)
	min := mStat(rMiss, "Min", 0)
	max := mStat(rMiss, "Max", 0)
		
//  N and % of subjects by gender and TG
// 	hits := make(map[string]map[string]int)
	
	
	
//  N and % of subjects by race and TG
	
	
	
	
	
	
	
// 	New Report 

	h := titles()
	f_scr := strconv.Itoa(nTG["Screened"])
	f_sf := strconv.Itoa(nTG["SF"])
	f := footnotes(f_scr, f_sf)
	err := WriteReport(outfile, h, f, nTG, nAge, meansd, median, min, max)
	if err != nil {
		fmt.Println(err)
	}
}