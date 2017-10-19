// Program to generate SDTM data for a fictitious study.
//
// This is not meant to represent an SC domain, but just be a framework upon which DM and VS can be built,
// so this data is used as a template for DM and VS, ensuring they are both internally consistent.
//
// Structure of the study:
// - Choose 100 subjects allocated to 5 sites.
// - Recruitment between 01Jan2010 and 31Dec2010 (any date in that window with no allowance for weekends, public holidays etc.)
// - A random 5% of the population will be screening failures. (RECTYPE=0)
// - Of the remainder, 35% will withdraw at any time after their start (RECTYPE=1).
// - 60% will last the 15 visits of the study (RECTYPE=2)
// - The full course of the study will be fortnightly visits for a maximum of 28 weeks
// - For simplicity, withdrawal is assumed at a scheduled visit, no unscheduled visits will be considered.
// - Screening (demog data) will be visit 0; subsequent visits (VS data) will be 1, 2, 3 etc to a maximum of 14
// Metadata:
// - STUDYID Char 6 (constant) Study Identifier
// - USUBJID Char 18 STUDYID-SITEID-SUBJID Unique Subject Identifier
// - SUBJID  Char 6 Subject Identifier
// - SITEID  Char 4 Site Identifier
// - RFSTDTC Char ISO8601 First date of study med exposure
// - RFENDTC Char ISO8601 Last date of study med exposure
// - DMDTC   Char ISO8601 Date/Time of Collection
// - RECTYPE Num  0=SF, 1=WD, 2=Completer
// - ENDV    Num  Last visit attended in study. RECTYPE=0 records will have 0 for this.
// - ARMCD   Num     Treatment Arm code
// - ARM     Char 7  Treatment Arm

package main

import (
//	"bufio"
	"flag"
//	"log"
	"math/rand"
//	"os"
	"strconv"
	"strings"
	"time"
//     "fmt"
)

// These types can be missing, so they are modelled by a pointer.
// In the case of an MV the value of the pointer address is Nil.
// BUT: Named types (eg RefStartDate) and pointers to them (*RefStartDate)
// are the only types that can appear in a receiver declaration.
// Method declarations are not permitted on named types that are themselves
// pointers. So:
// type P *int
// func (P) f() {  /* compile error. Invalid receiver type */ }
type RefStartDate   time.Time
type RefEndDate     time.Time
type ArmCD          int
type Arm            string

type Subject struct {
	studyid string
	subjid  string
	siteid  string
	usubjid string
	rectype int
	dmdtc   time.Time
	endv    int
	rfstdtc RefStartDate
	rfendtc RefEndDate
	armcd	ArmCD
	arm		Arm
}

// NB This defines Methods acting on types, not functions
type DateFmt interface {
    dateFmt() string
}

// Constants can only be numbers, strings or boolean
const (
	studyid   = "XYZ123" // Study Identifier
	nSubj     = 100      // Number of subjects in the study
	lastVisit = 14
)

// The study can start at any date within 2010
var baseDate = time.Date(2010, time.January, 1, 0, 0, 0, 0, time.UTC)

// The SITEID is chosen from one of these 5 values
var siteids = []string{"1", "2", "3", "4", "5"}

// The program will be run with a flag to specify the output file
var outfile = flag.String("o", "sc.csv", "Name of output file")

// To allow a random choice of arm
var arm = map[int]string{0:"Placebo", 1:"Active"}

/*
func (s *RefStartDate) dateFmt() string {
    var d time.Time
    d = *s
    return d.Format("2006-01-02")
}

func (s *RefEndDate) dateFmt() string {
    var d time.Time
    d = *s
    return d.Format("2006-01-02")
}
*/

// This pads the string in the 1st arg to the length
// in the 3rd arg with the char in the 2nd arg
func leftPad2Len(s string, padStr string, overallLen int) string {
	var padCountInt int
	padCountInt = 1 + ((overallLen - len(padStr)) / len(padStr))
	var retStr = strings.Repeat(padStr, padCountInt) + s
	return retStr[(len(retStr) - overallLen):]
}

// Select a random member from a slice of strings
func choice(s []string) string {
	// Allocate seed for generating random numbers
	rand.Seed(time.Now().UTC().UnixNano())
	return s[rand.Intn(len(s))]
}

// Use this to flag subjects as
// - screening failures (~5%)
// - withdrawers (~35%)
// - completers (~60%)
func ptype() int {
	rand.Seed(time.Now().UTC().UnixNano())
	x := rand.Float64()
	switch {
	case x <= 0.05:
		return 0
	case x > 0.05 && x < 0.4:
		return 1
	default:
		return 2
	}
}

// For each subject randomly select their last visit
// depending upon whether they are withdrawers, completers or screening failures.
func endv(r int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	switch r {
	case 0:
		return 0
	case 1:
		// Because the discontinuing patient has been dosed,
		// cannot finish at visit 0
		// So choose randomly from 0 to 13 and then add 1
		return (rand.Intn(lastVisit - 1)) + 1
	default:
		return lastVisit
	}
}

// For screen failures, the ref start and end dates should really be missing.
// Hence pointer type used.
func startDate(r int, d time.Time) *RefStartDate {
    var d2 RefStartDate
	switch r {
	case 0:
		return nil
	default:
        d2 := d.AddDate(0, 0, 14)
		return &d2
	}
}

// Construct an end date dependent upon the last visit
func endDate(r int, e int, d time.Time) *RefEndDate {
	switch r {
	case 0:
		return nil
	case 1:
        d2 := d.AddDate(0, 0, (e * 14))
		return &d2
	default:
        d2 := d.AddDate(0, 0, (lastVisit * 14))
		return &d2
	}
}

func getArm(r int) (*ArmCD, *Arm) {
	rand.Seed(time.Now().UTC().UnixNano())
    if r != 0 {
        armcd := rand.Intn(len(arm))
        arm := arm[armcd]
        return &armcd, &arm
    } else {
        return nil, nil
    }
}

func main() {
	// Trap output file name
	flag.Parse()

	// Create slice of pointers to Subject types
	sSubj := make([]*Subject, nSubj)

	for ii := 0; ii < nSubj; ii++ {

		subjid := leftPad2Len(strconv.Itoa(ii+1), "0", 6)
		siteid := leftPad2Len(choice(siteids), "0", 4)
		usubjsl := []string{studyid, siteid, subjid}
		usubjid := strings.Join(usubjsl, "-")
		rectype := ptype()
		dmdtc := baseDate.AddDate(0, 0, rand.Intn(364))
        endv := endv(rectype)
		rfstdtc := startDate(rectype, dmdtc)
		
		rfendtc := endDate(rectype, endv, dmdtc)
		armcd, arm := getArm(rectype)

		sSubj[ii] = &Subject{
			studyid,
			subjid,
			siteid,
			usubjid,
			rectype,
			dmdtc,
            endv,
			rfstdtc,
			rfendtc,
			armcd,
			arm,
		}
		
		/*
		fmt.Println(*sSubj[ii])
        fmt.Println(((*sSubj[ii]).dmdtc).Format("2006-01-02"))
        // Automatic pointer derefencing
        if sSubj[ii].rfstdtc != nil {
            fmt.Println(sSubj[ii].rfstdtc)
        } else {
            fmt.Println("Missing StartDate for SF!!")
        }
        
        if sSubj[ii].rfendtc != nil {
            fmt.Println(sSubj[ii].rfendtc)
        } else {
            fmt.Println("Missing EndDate for SF!!")
        } 
        
        if sSubj[ii].arm != nil {
            fmt.Println(sSubj[ii].arm)
        } else {
            fmt.Println("Missing arm text for SF!!")
        }          
        
        if sSubj[ii].armcd != nil {
            fmt.Println(sSubj[ii].armcd)
        } else {
            fmt.Println("Missing arm code for SF!!")
        }   
        */
	}
	
	
	// Before writing to the output file (as strings)
	// the structs need to be sorted to ensure 
	// USUBJID order
	
	

	// Output to external file via strings
	// Could substitute rectype=0 ref dates with missing strings
	// but what would happen when the data is read into a
	// struct for further processing??
	
	/*
	fo, err := os.Create(*outfile)
	if err != nil {
		log.Fatal(err)
	}
	defer fo.Close()

	// Create a buffered writer from the file
	w := bufio.NewWriter(fo)

	for ii := 0; ii < nSubj; ii++ {
        var refstart, refend string
        
        if sSubj[ii].rfstdtc != nil {
            refstart = (sSubj[ii].rfstdtc).dataFmt
        } else {
            refstart = ""
        }
        
		bytesWritten, err := w.WriteString(
			sSubj[ii].studyid + "," +
				sSubj[ii].subjid + "," +
				sSubj[ii].siteid + "," +
				sSubj[ii].usubjid + "," +
				strconv.Itoa(sSubj[ii].rectype) + "," +
				sSubj[ii].dmdtc.Format("2006-01-02") + "," +
				strconv.Itoa(sSubj[ii].endv) + "," +
				refstart + "," +
				//sSubj[ii].rfendtc.Format("2006-01-02") + "," +
				//strconv.Itoa(sSubj[ii].armcd) + "," +
				//sSubj[ii].arm +
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
