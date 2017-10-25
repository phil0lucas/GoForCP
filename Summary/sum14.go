package main

import (
	"fmt"
	"flag"
// 	"time"
// 	"log"
// 	"os"
	"strconv"
// 	"path/filepath"
// 	"strings"
	
	"github.com/phil0lucas/GoForCP/CPUtils"
	"github.com/phil0lucas/GoForCP/DM"	
	"github.com/jung-kurt/gofpdf"
	"github.com/montanaflynn/stats"	
)

var infile = flag.String("i", "../CreateData/dm3.csv", "Name of input file")
var outfile = flag.String("o", "summary14.pdf", "Name of output file")

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
				 max map[string]string,
				 sexPct map[Key]string,
				 racePct map[KeyR]string) error {
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
	
//	Gender	
	uV := uniqueValues (sexPct)
// 	fmt.Println(uV)
	var iter int
	var col1text string
	sexFmt := map[string]string{"F":"Female", "M": "Male"}
	for _, v := range uV{
// 		fmt.Println(v)
// 		index, ok := sexPct[Key{v,"Placebo"}]
// 		fmt.Println(index)
// 		fmt.Println(ok)
		if iter == 0 {
			col1text = "Gender"
		} else {
			col1text = ""
		}
		textSlice7 := []string{col1text, sexFmt[v], 
			sexPct[Key{v, "Placebo"}], 
			sexPct[Key{v, "Active"}],
			sexPct[Key{v, "Overall"}]}
		for i, str := range textSlice7 {
			pdf.CellFormat(colWidthSlice[i], 8, str, "", 0, colJustSlice[i], false, 0, "")
		}
		pdf.Ln(4)
		iter++
	}
	
//	Race

	pdf.Ln(4)
	uVr := uniqueValuesR (racePct)
	fmt.Println(uVr)	
	var iterR int
	var col1textR string
	for _, v := range uVr{
		if iterR == 0 {
			col1textR = "Race"
		} else {
			col1textR = ""
		}
		textSlice8 := []string{col1textR, v, 
			racePct[KeyR{v, "Placebo"}], 
			racePct[KeyR{v, "Active"}],
			racePct[KeyR{v, "Overall"}]}
		for i, str := range textSlice8 {
			pdf.CellFormat(colWidthSlice[i], 8, str, "", 0, colJustSlice[i], false, 0, "")
		}
		pdf.Ln(4)
		iterR++
	}
	
//	Underline	
	pdf.SetY(-36)
	colUnderSlice := []string{" ", " ", " ", " ", " "}
	for i, str := range colUnderSlice {
		pdf.CellFormat(colWidthSlice[i], 8, str, "B", 0, colJustSlice[i], false, 0, "")
	}	
	
	
// 	Output
	err := pdf.OutputFileAndClose(*outputFile)
	fmt.Println(err)
	return err
} 

func nMiss (dm []*DM.Dmrec) map[string]int {
	m := make(map[string]int)
// 	total := 0
	for _, v := range dm {
		if v.Age != nil {
			m[*v.Arm]++
			m["Overall"]++
		} else {
			m["Missing"]++
		}
	}
	return m
}

func prepareData(dm []*DM.Dmrec, tg []string) map[string][]float64{
	m := make(map[string][]float64)
	for _, s := range tg {
		var out []float64
		for _, v := range dm {
			if v.Age != nil {
				if s == *v.Arm {
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

type Key struct {
    sex string
    arm string
}


func countSexByTG (dm []*DM.Dmrec) map[Key]int{
	var r []Key
	for _, v := range dm {
// 		fmt.Println(v)
		var k Key
		if v.Sex != nil {
			k.sex = *v.Sex
			k.arm = *v.Arm
			r = append(r, k)
			k.arm = "Overall"
			r = append(r, k)
		}
	}
	
	// calculate sum:
	m := make(map[Key]int)
	for _, v := range r {
		m[v]++
	}

	return m
}

func pctSexByTG (m map[Key]int, tg map[string]int) map[Key]string{
	outmap := make(map[Key]string)
	for k, v := range m {
// 		fmt.Println(k.arm)
// 		fmt.Println(tg[k.arm])
// 		fmt.Println(v)
		var pct float64
		pct = (float64(v) / float64(tg[k.arm])) * 100
// 		fmt.Println(pct)
		c_pct := strconv.FormatFloat(pct, 'f', 2, 64)
// 		fmt.Println(c_pct)
		c_stat := strconv.FormatInt(int64(v), 10) + " (" +
			c_pct + "%)"
// 		fmt.Println(c_stat)
		outmap[k] = c_stat
	}
	return outmap
}

//	Determine the unique values of the non-TG key
func uniqueValues (m map[Key]string) []string {
	var uValues []string
	for k, _ := range m {
		if !CPUtils.StringInSlice (k.sex , uValues) {
			uValues = append(uValues, k.sex)
		}
	}
	return uValues
}

type KeyR struct {
    race string
    arm string
}

func countRaceByTG (dm []*DM.Dmrec) map[KeyR]int{
	var r []KeyR
	for _, v := range dm {
// 		fmt.Println(v)
		var k KeyR
		if v.Race != nil {
			k.race = *v.Race
			k.arm = *v.Arm
			r = append(r, k)
			k.arm = "Overall"
			r = append(r, k)
		}
	}
	// calculate sum:
	m := make(map[KeyR]int)
	for _, v := range r {
		m[v]++
	}
	return m
}

//	Determine the unique values of the non-TG key
func uniqueValuesR (m map[KeyR]string) []string {
	var uValues []string
	for k, _ := range m {
		if !CPUtils.StringInSlice (k.race , uValues) {
			uValues = append(uValues, k.race)
		}
	}
	return uValues
}

func pctRaceByTG (m map[KeyR]int, tg map[string]int) map[KeyR]string{
	outmap := make(map[KeyR]string)
	for k, v := range m {
		var pct float64
		pct = (float64(v) / float64(tg[k.arm])) * 100
// 		fmt.Println(pct)
		c_pct := strconv.FormatFloat(pct, 'f', 2, 64)
// 		fmt.Println(c_pct)
		c_stat := strconv.FormatInt(int64(v), 10) + " (" +
			c_pct + "%)"
// 		fmt.Println(c_stat)
		outmap[k] = c_stat
	}
	return outmap
}

func main() {
	// Read the file and dump into the slice of structs
	dm := DM.ReadDM(infile)
// 	fmt.Println(dm)

// 	Compute number of subjects by treatment group
	nTG := DM.CountByTG(dm)
	fmt.Println(nTG)
	
// Select treatment groups to display i.e. Placebo, Active, Overall	
	TGs := selectTGs(nTG)
	fmt.Println(TGs)
	
// Create version of dm without the SFs
	dm2 := DM.RemoveSF(dm)
	fmt.Printf("%T %v\n", dm2, len(dm2))
	
// 	Compute number of non-missing Age values by TG
	nAge := nMiss(dm2)
	fmt.Println(nAge)
	
//	Prepare the data for passing to the stats functions
//	Select only the TGs to display
//	Remove the missing values.
	rMiss := prepareData(dm2, TGs)
	fmt.Println(rMiss)
	
// 	Compute stats of age by TG
	mean := mStat(rMiss, "Mean", 2)
	fmt.Println(mean)
	sd := mStat(rMiss, "SD", 2)
	fmt.Println(sd)

//	Concatenate mean and SD values into a display string
	meansd := make(map[string]string)
	for k, _ := range mean {
		meansd[k] = mean[k] + " (" + sd[k] + ")"
	}
	fmt.Println(meansd)
	
	median := mStat(rMiss, "Median", 0)
	fmt.Println(median)
	min := mStat(rMiss, "Min", 0)
	fmt.Println(min)
	max := mStat(rMiss, "Max", 0)
	fmt.Println(max)


//  N and % of subjects by gender and TG
	keyValues := countSexByTG(dm2)
	fmt.Println(keyValues)

	pctMap := pctSexByTG(keyValues, nTG)
	fmt.Println(pctMap)
	
	raceValues := countRaceByTG(dm2)
	fmt.Println(raceValues)	
	
//  N and % of subjects by race and TG
	pctRace := pctRaceByTG(raceValues, nTG)
	fmt.Println(pctRace)	
	
// 	Report 
	h := titles()
	f_scr := strconv.Itoa(nTG["Screened"])
	f_sf := strconv.Itoa(nTG["SF"])
	f := footnotes(f_scr, f_sf)
	err := WriteReport(outfile, h, f, nTG, nAge, meansd, median, min, max, 							pctMap, pctRace)
	if err != nil {
		fmt.Println(err)
	}
}
