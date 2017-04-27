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


package main

import (
	"bufio"
	"os"
	"log"
	"strconv"
	"strings"
	"math/rand"
	"time"
)

type Subject struct {
    studyid string
    subjid string
    siteid string
    usubjid string
    rectype int
    dmdtc time.Time
    endv int
    rfstdtc time.Time
    rfendtc time.Time
}

// Constants can only be numbers, strings or boolean
const (
	studyid = "XYZ123" 	// Study Identifier
	nSubj = 100			// Number of subjects in the study
	lastVisit = 14
)

var baseDate = time.Date(2010, time.January, 1, 0, 0, 0, 0, time.UTC)
var siteids = []string{"1", "2", "3", "4", "5"}

// This pads the string in the 1st arg to the length
// in the 3rd arg with the char in the 2nd arg
func leftPad2Len(s string, padStr string, overallLen int) string {
	var padCountInt int
	padCountInt = 1 + ((overallLen - len(padStr)) / len(padStr))
	var retStr = strings.Repeat(padStr, padCountInt) + s
	return retStr[(len(retStr) - overallLen):]
}

func choice (s []string) string {
	// Allocate seed for generating random numbers
	rand.Seed(time.Now().UTC().UnixNano())
	return s[rand.Intn(len(s))]
}

func ptype () int {
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

func endv(r int) int{
	rand.Seed(time.Now().UTC().UnixNano())
	switch r{
	case 0:
		return 0
	case 1:
		// Because the discontinuing patient has been dosed, 
		// cannot finish at visit 1
		// So choose randomly from 0 to 13 and then add 1
		return (rand.Intn(lastVisit-1)) + 1
	default:
		return lastVisit
	}	
}

// For screen failures, the ref start and end dates should really be missing
// Let's substitute the screening date
func startDate(r int, d time.Time) time.Time{
	switch r{
		case 0:
			return d
		default:
			return d.AddDate(0,0,14)
	}	
}

func endDate(r int, e int, d time.Time) time.Time {
	switch r{
		case 0:
			return d
		case 1:
			return d.AddDate(0,0,(e*14))
		default:
			return d.AddDate(0,0,(lastVisit * 14))
	}
}

func main(){
	
	// Create slice of pointers to Subject types
	sSubj := make([]*Subject, nSubj)
	
	for ii:=0; ii<nSubj; ii++{
		
		subjid := leftPad2Len(strconv.Itoa(ii), "0", 6)
		siteid := leftPad2Len(choice(siteids), "0", 4)
		usubjsl := []string{studyid, siteid, subjid}
        usubjid := strings.Join(usubjsl, "-")
		rectype := ptype()
		dmdtc := baseDate.AddDate(0,0,rand.Intn(364))
		rfstdtc := startDate(rectype, dmdtc)
		endv := endv(rectype)
		rfendtc := endDate(rectype, endv, dmdtc)
					
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
		}
	}
	
	
	
	
	
	
	// Output to external file via strings
	// Could substitute rectype=0 ref dates with missing strings
	// but what would happen when the data is read into a 
	// struct for further processing??
    fo, err := os.Create("output14.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer fo.Close()
	
    // Create a buffered writer from the file
	w := bufio.NewWriter(fo)
	
	for ii:=0; ii<nSubj; ii++{
		
		bytesWritten, err := w.WriteString(
			sSubj[ii].studyid + "," +
			sSubj[ii].subjid + "," +
			sSubj[ii].siteid + "," +
			sSubj[ii].usubjid + "," +
			strconv.Itoa(sSubj[ii].rectype) + "," +
			sSubj[ii].dmdtc.Format("2006-01-02") + "," +
			strconv.Itoa(sSubj[ii].endv) + "," +
			sSubj[ii].rfstdtc.Format("2006-01-02") + "," +
			sSubj[ii].rfendtc.Format("2006-01-02") +
			"\n")
		
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Bytes written: %d\n", bytesWritten)
	}
	
	// Write to disk
	w.Flush()
} 
