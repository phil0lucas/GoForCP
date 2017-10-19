// Program to generate SDTM data for a fictitious study.
// Domain VS
// Metadata :
// - STUDYID 	Char 6  (constant) Study Identifier
// - DOMAIN  	Char 2  (constant) Domain abbreviation
// - USUBJID 	Char 18 STUDYID-SITEID-SUBJID Unique Subject Identifier (Key variable 1)
// - SUBJID  	Char 6  Subject Identifier
// - SITEID  	Char 4  Site Identifier
// - VSSEQ   	Num	 	Sequence number (Key variable 2)
// - VISITNUM	Num     Visit number (0=Screening, 1-14=Dosing visits and assessments)
// - VSTESTCD	Char 3  Test code
// - VSTEST		Char 30 Test description
// - VSORRES	Num 	Original recorded result
// - VSORRESU   Char    Units of original result
// - VSSTRESC   Char	Standardized result in char form
// - VSSTRESN   Num     Standardized result in numeric form
// - VSSTRESU   Char    Units of result in standardized form
// - VSBLFL     Char    Flags baseline visit
// - VSDTC      Date    Date of visit in ISO8601  
// - VSDY    	Num     Study Day of collection

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
	"flag"
//	"math/rand"
	"strconv"
//	"log"
)

// This will mirror the metadata above with more natural types
type vsrec struct {
	studyid 	string
	domain  	string
	usubjid 	string
	subjid  	string
	siteid  	string
	vsseq   	int
	visitnum	int
	vstestcd	string
	vstest		string
	vsorres		float64
	vsorresu	string
	vsstresc	string
	vsstresn	float64
	vsstresu	string
	vsblfl		string
	vsdtc		time.Time
	vsdy		int
}

// The program will be run with flags to specify the input & output files
var infile = flag.String("i", "../SC/sc.csv", "Name of input file")
var outfile = flag.String("o", "vs.csv", "Name of output file")
var testcodes = []string{"SBP", "DBP", "HR"}
var testnames = []string{"Systolic Blood Pressure", "Diastolic Blood Pressure", "Heart Rate"}

const (
	domain 	= "VS"
)

func genBaseline(tcode string, visit int) float64 {
	if visit == 0 {
		return 100.00
	}
}

func main() {
	flag.Parse()
	
	// open the file and pass it to a Scanner object
	file, err := os.Open(*infile)
	if err != nil {
		panic(fmt.Sprintf("error opening %s: %v", infile, err))
	}
	defer file.Close()

	// Output slice of pointers to structs
	var vs []*vsrec

	// Pass the opened file to a scanner
	scanner := bufio.NewScanner(file)

	for i := 0; scanner.Scan(); i++ {
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "error reading from file:", err)
			os.Exit(3)
		}
		str := scanner.Text()
		studyid := strings.Split(str, ",")[0]
		usubjid := strings.Split(str, ",")[3]
		subjid := strings.Split(str, ",")[1]
		siteid := strings.Split(str, ",")[2]
		endv := strings.Split(str, ",")[6]
		endvn, _ := strconv.Atoi(endv)
		//fmt.Println(endvn)
		
		// Add in the visits up to the generated end-visit
		// Subjects with just visit 0 are screening failures.
		// Subjects with a final visit number < 14 are withdrawers.
		for j := 0; j <= endvn; j++ {
			for k := 0; k< len(testcodes); k++ {
				baseline := genBaseline(testcodes[k], j)
				fmt.Println(baseline)	
				vs = append(vs, &vsrec{
					studyid: studyid,
					domain:  domain,
					usubjid: usubjid,
					subjid: subjid,
					siteid: siteid,
					visitnum: j,
					vstestcd: testcodes[k],
					vstest: testnames[k]})
				
				
				
				
				
					//fmt.Println(*vs[i])
			}
		}
	}
	
	/*
	fo, err := os.Create(*outfile)
	if err != nil {
		log.Fatal(err)
	}
	defer fo.Close()
	
	// Create a buffered writer from the file
	w := bufio.NewWriter(fo)
	
	for ii, _ := range vs{
		bytesWritten, err := w.WriteString(
			vs[ii].studyid + "," +
			vs[ii].domain + "," +
			vs[ii].subjid + "," +
			vs[ii].siteid + "," +
			vs[ii].usubjid + "," +
			strconv.Itoa(vs[ii].visitnum) + "," +
			vs[ii].vstestcd + "," +
			vs[ii].vstest +
			"\n")

		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Bytes written: %d\n", bytesWritten)
	}

	// Write to disk
	w.Flush()	
	*/
}