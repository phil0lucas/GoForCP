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
	"math/rand"
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
var outfile = flag.String("o", "vs2.csv", "Name of output file")
var testcodes = []string{"SBP", "DBP", "HR"}
var testnames = []string{"Systolic Blood Pressure", "Diastolic Blood Pressure", "Heart Rate"}

const (
	domain 	= "VS"
)

func randValue(max, min int) float64 {
	rand.Seed(time.Now().UTC().UnixNano())
	return float64(rand.Intn(max - min) + min)
}


func genBaseline(tcode string) float64 {
	switch tcode {
		// return rand.Intn(max - min) + min
		case "HR":
			return randValue(120, 70)
		case "SBP":
			return randValue(160, 120)
		case "DBP":
			return randValue(120, 90)
	}
	return 0.0
}

func getOrigRes(baseline float64, vstestcd string, visitnum int, armcd int) float64 {
	if armcd == 0 {
		if visitnum == 0 {
			return baseline
		} else {
			return baseline + randValue(5, -5)
		}
	} else {
		if visitnum == 0 {
			return baseline
		} else if visitnum < 5 {
			return baseline * 0.975 + randValue(2, -3)
		} else if visitnum < 8 {
			return baseline * 0.95 + randValue(1, -5)
		} else if visitnum < 11 {
			return baseline * 0.925 + randValue(0, -7)
		} else {
			return baseline * 0.9 + randValue(-3, -10)
		}
	}
}

func getUnits (testcode string) (string, string) {
	if testcode == "HR" {
		return "bpm", "bpm"
	} else {
		return "mmHg", "mmHg"
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
		armcd, _ := strconv.Atoi(strings.Split(str, ",")[9])
		fmt.Printf("Subject %s\n", usubjid)
		
		// Add in the visits up to the generated end-visit
		// Subjects with just visit 0 are screening failures.
		// Subjects with a final visit number < 14 are withdrawers.
		
		for j := 0; j < len(testcodes); j++ {
			baseline := genBaseline(testcodes[j])
			fmt.Printf("Test code %s value %v\n", testcodes[j], baseline)	
			vstestcd := testcodes[j]
			vstest := testnames[j]
			vsorresu, vsstresu := getUnits(vstestcd)
			//fmt.Printf("   Test code %s\n", vstestcd)
			
			for k := 0; k <= endvn; k++ {	
				vsorres := getOrigRes(baseline, vstestcd, k, armcd)
				vs = append(vs, &vsrec{
					studyid: studyid,
					domain:  domain,
					usubjid: usubjid,
					subjid: subjid,
					siteid: siteid,
					visitnum: k,
					vstestcd: vstestcd,
					vstest: vstest,
					vsorres: vsorres,
					vsstresn: vsorres,
					vsstresc: strconv.FormatFloat(vsorres, 'f', 2, 64),
					vsorresu: vsorresu,
					vsstresu: vsstresu,
					
				})
				
				//fmt.Printf("      Visit Number %v Std Char Result %v\n", k, vsorres)
				
				
				
				
				
				fmt.Println(*vs[i])
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