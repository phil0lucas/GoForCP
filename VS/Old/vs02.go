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
	"sort"
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
	vsblfl		bool
	vsdtc		time.Time
	vsdy		int
}

// Sorting depends on an interface that demands 3 methods
type lessFunc func(p1, p2 *vsrec) bool

// multiSorter implements the Sort interface, sorting the changes within.
type multiSorter struct {
	vsrecs  []vsrec
	less    []lessFunc
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



// Sort sorts the argument slice according to the less functions passed to OrderedBy.
func (ms *multiSorter) Sort(vsrecs []vsrec) {
	ms.vsrecs = vsrecs
	sort.Sort(ms)
}

// OrderedBy returns a Sorter that sorts using the less functions, in order.
// Call its Sort method to sort the data.
func OrderedBy(less ...lessFunc) *multiSorter {
	return &multiSorter{
		less: less,
	}
}

// Len is part of sort.Interface.
func (ms *multiSorter) Len() int {
	return len(ms.vsrecs)
}

// Swap is part of sort.Interface.
func (ms *multiSorter) Swap(i, j int) {
	ms.vsrecs[i], ms.vsrecs[j] = ms.vsrecs[j], ms.vsrecs[i]
}

// Less is part of sort.Interface. It is implemented by looping along the
// less functions until it finds a comparison that is either Less or
// !Less. Note that it can call the less functions twice per call. 
func (ms *multiSorter) Less(i, j int) bool {
	p, q := &ms.vsrecs[i], &ms.vsrecs[j]
	// Try all but the last comparison.
	var k int
	for k = 0; k < len(ms.less)-1; k++ {
		less := ms.less[k]
		switch {
		case less(p, q):
			// p < q, so we have a decision.
			return true
		case less(q, p):
			// p > q, so we have a decision.
			return false
		}
		// p == q; try the next comparison.
	}
	// All comparisons to here said "equal", so just return whatever
	// the final comparison reports.
	return ms.less[k](p, q)
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
		dmdtc, _ := time.Parse("2006-01-02", strings.Split(str, ",")[5])
		endv := strings.Split(str, ",")[6]
		endvn, _ := strconv.Atoi(endv)
		armcd, _ := strconv.Atoi(strings.Split(str, ",")[9])
		fmt.Printf("Subject %s\n", usubjid)
		
		// Add in the visits up to the generated end-visit
		// Subjects with just visit 0 are screening failures.
		// Subjects with a final visit number < 14 are withdrawers.
		
		// Test codes
		for j := 0; j < len(testcodes); j++ {
			baseline := genBaseline(testcodes[j])
			fmt.Printf("Test code %s value %v\n", testcodes[j], baseline)	
			vstestcd := testcodes[j]
			vstest := testnames[j]
			vsorresu, vsstresu := getUnits(vstestcd)
			//fmt.Printf("   Test code %s\n", vstestcd)
			var vsblfl bool
			
			// Visits
			for k := 0; k <= endvn; k++ {
				if k == 1 {
					vsblfl = true
				} else {
					vsblfl = false
				}
				vsorres := getOrigRes(baseline, vstestcd, k, armcd)
				vsdtc := dmdtc.AddDate(0, 0, (k*14))
				vsdy := k * 14
				
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
					vsblfl: vsblfl,
					vsdtc: vsdtc,
					vsdy: vsdy,
				})
				
				fmt.Println(*vs[i])
			}
			
		}
	}
	
	// Sort the struct of VS 'records'
	// Closures that order the Change structure.
	kusubjid := func(c1, c2 vsrec) bool {
		return c1.usubjid < c2.usubjid
	}
	kvisitnum := func(c1, c2 vsrec) bool {
		return c1.visitnum < c2.visitnum
	}
	ktestcd := func(c1, c2 vsrec) bool {
		return c1.vstestcd < c2.vstestcd
	}

	OrderedBy(kusubjid, kvisitnum, ktestcd).Sort(vs)
	fmt.Println(vs)
	
	
	
	
	
	
	
	
	
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