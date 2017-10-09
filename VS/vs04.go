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
	"flag"
	"fmt"
// 	"log"
	"math/rand"
	"os"
// 	"sort"
	"strconv"
	"strings"
	"time"
)

// This will mirror the metadata above with more natural types.
// The elements modelled as pointers may have missing values i.e.
// a nil pointer
type vsrec struct {
	studyid  string
	domain   string
	usubjid  string
	subjid   string
	siteid   string
	vsseq    int
	visitnum int
	vstestcd string
	vstest   string
	vsorres  *float64
	vsorresu *string
	vsstresc *string
	vsstresn *float64
	vsstresu *string
	vsblfl   bool
	vsdtc    time.Time
	vsdy     int
}

//	The type vsrecs models a 'data set' as a slice of 
//	pointers to vsrec structs.
type vsrecs []*vsrec

// The program will be run with flags to specify the input & output files
var infile = flag.String("i", "../SC/sc.csv", "Name of input file")
var outfile = flag.String("o", "vs4.csv", "Name of output file")
var testcodes = []string{"SBP", "DBP", "HR"}
var testnames = []string{"Systolic Blood Pressure", "Diastolic Blood Pressure", "Heart Rate"}

const (
	domain = "VS"
)

func flagMiss () bool {
	rand.Seed(time.Now().UTC().UnixNano())
	if rand.Float64() >= 0.05 {
		return false
	} else {
		return true
	}
}

func randValue(max, min int) float64 {
	rand.Seed(time.Now().UTC().UnixNano())
	return float64(rand.Intn(max-min) + min)
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
			return baseline*0.975 + randValue(2, -3)
		} else if visitnum < 8 {
			return baseline*0.95 + randValue(1, -5)
		} else if visitnum < 11 {
			return baseline*0.925 + randValue(0, -7)
		} else {
			return baseline*0.9 + randValue(-3, -10)
		}
	}
}

func getUnits(testcode string) (string, string) {
	if testcode == "HR" {
		return "bpm", "bpm"
	} else {
		return "mmHg", "mmHg"
	}
}

func (t vsrecs) Len() int {
	return len(t)
}

func (t vsrecs) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func (t vsrecs) Less(i, j int) bool {
	if t[i].usubjid < t[j].usubjid {
		return true
	}
	if t[i].usubjid > t[j].usubjid {
		return false
	}
	// If USUBJIDs are equal
	if t[i].vstestcd < t[j].vstestcd {
		return true
	}
	if t[i].vstestcd > t[j].vstestcd {
		return false
	}
	return t[i].visitnum < t[j].visitnum
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
// 	var vs vsrecs

	// Pass the opened file to a scanner
	scanner := bufio.NewScanner(file)

	// For each subject
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
		dmdtc, _ := time.Parse("2006-01-02", strings.Split(str, ",")[5])
		endv := strings.Split(str, ",")[6]
		endvn, _ := strconv.Atoi(endv)
		armcd, _ := strconv.Atoi(strings.Split(str, ",")[9])
		fmt.Printf("Study=%s Subject=%s Subjid=%s Siteid=%s\n", studyid, usubjid, subjid, siteid)
		fmt.Printf("%v %v %v %v\n", dmdtc, endv, endvn, armcd)

		// Add in the visits up to the generated end-visit
		// Subjects with just visit 0 are screening failures.
		// Subjects with a final visit number < 14 are withdrawers.

		// Test codes
		for j := 0; j < len(testcodes); j++ {
			vstestcd := testcodes[j]
			vstest := testnames[j]
			fmt.Printf("Testcode=%s Test=%s\n", vstestcd, vstest)
			baseline := genBaseline(testcodes[j])
			fmt.Printf("Test code %s value %v\n", testcodes[j], baseline)

// 			vsorresu, vsstresu := getUnits(vstestcd)
// 			//fmt.Printf("   Test code %s\n", vstestcd)
// 			var vsblfl bool

			// Visits
// 			for k := 0; k <= endvn; k++ {
// 				if k == 1 {
// 					vsblfl = true
// 				} else {
// 					vsblfl = false
// 				}
// 				vsorres := getOrigRes(baseline, vstestcd, k, armcd)
// 				vsdtc := dmdtc.AddDate(0, 0, (k * 14))
// 				vsdy := k * 14

// 				vs = append(vs, vsrec{
// 					studyid:  studyid,
// 					domain:   domain,
// 					usubjid:  usubjid,
// 					subjid:   subjid,
// 					siteid:   siteid,
// 					visitnum: k,
// 					vstestcd: vstestcd,
// 					vstest:   vstest,
// 					vsorres:  vsorres,
// 					vsstresn: vsorres,
// 					vsstresc: strconv.FormatFloat(vsorres, 'f', 2, 64),
// 					vsorresu: vsorresu,
// 					vsstresu: vsstresu,
// 					vsblfl:   vsblfl,
// 					vsdtc:    vsdtc,
// 					vsdy:     vsdy,
// 				})
// 
// 				//fmt.Println(vs[i])
// 			}

		}	//	End j loop
	}	// End i loop
	
	/* Rubbish
	rand.Seed(time.Now().UTC().UnixNano())
	for ii, _ := range vs{
        randno := rand.Intn(100)
        if randno < 5 {
            vs[ii].vsorres = -0.0
            vs[ii].vsstresn = -0.0
            vs[ii].vsstresc = "-0.0"
        }
    }
    */
	
	
	

/*	
	// Sort the struct of VS 'records'
	sort.Sort(vsrecs(vs))
	// fmt.Println(vs)

	// Define VSSEQ as key as running int within each subject
	// Need to define a variable external to the loop otherwise
	// scope will make each value 1
	var count int
	for ii := 0; ii < len(vs); ii++ {
		if ii == 0 || ((vs[ii].usubjid != vs[ii-1].usubjid)) {
			count = 0
		}
		count++
		vs[ii].vsseq = count
	}
	//fmt.Println(vs)
*/

/*	
	// Write to external file.
	fo, err := os.Create(*outfile)
	if err != nil {
		log.Fatal(err)
	}
	defer fo.Close()

	// Create a buffered writer from the file
	w := bufio.NewWriter(fo)

	for ii, _ := range vs {
		bytesWritten, err := w.WriteString(
			vs[ii].studyid + "," +
				vs[ii].domain + "," +
				vs[ii].subjid + "," +
				vs[ii].siteid + "," +
				vs[ii].usubjid + "," +
				strconv.Itoa(vs[ii].vsseq) + "," +
				strconv.Itoa(vs[ii].visitnum) + "," +
				vs[ii].vstestcd + "," +
				vs[ii].vstest + "," +
				strconv.FormatFloat(vs[ii].vsorres, 'f', 1, 64) + "," +
				strconv.FormatFloat(vs[ii].vsstresn, 'f', 1, 64) + "," +
				vs[ii].vsstresc + "," +
				vs[ii].vsorresu + "," +
				vs[ii].vsstresu + "," +
				strconv.FormatBool(vs[ii].vsblfl) + "," +
				vs[ii].vsdtc.Format("2006-01-02") + "," +
				strconv.Itoa(vs[ii].vsdy) +
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
