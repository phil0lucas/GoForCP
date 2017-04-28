// Program to generate SDTM data for a fictitious study.
// Domain DM
// Metadata :
// - STUDYID Char 6  (constant) Study Identifier
// - DOMAIN  Char 2  (constant) Domain abbreviation
// - USUBJID Char 18 STUDYID-SITEID-SUBJID Unique Subject Identifier
// - SUBJID  Char 6  Subject Identifier
// - SITEID  Char 4  Site Identifier
// - RFSTDTC Char 10 ISO8601 First date of study med exposure
// - RFENDTC Char 10 ISO8601 Last date of study med exposure
// - DMDTC   Char 10 ISO8601 Date/Time of Collection
// - INVID   Char 3  Investigator code
// - INVNAME Char 8  Investigator Name
// - COUNTRY Char 3  ISO3166 Country code
// - BRTHDTC Char 10 ISO8601 Subjects date of birth
// - AGE	 Char 2  Subject's age (min 20, Max 80)
// - AGEU    Char 5  (constant) Age units
// - SEX     Char 1  Subject's gender ((M/F)
// - RACE    Char 5  Subject's race (White, Black, Asian)
// - ARMCD   Char 1  Treatment Arm code
// - ARM     Char 7  Treatment Arm
// - DMDY    Char 3  Study Day of collection

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
	"flag"
	"math/rand"
)

// This will mirror the metadata above with more natural types
type dmrec struct {
	studyid string
	domain  string
	usubjid string
	subjid  string
	siteid  string
	rfstdtc time.Time
	rfendtc time.Time
	dmdtc   time.Time
	invid	string
	invname string
	country string
	ageu	string
	age		int
}

// The program will be run with flags to specify the input & output files
var infile = flag.String("i", "../SC/sc.csv", "Name of input file")
var outfile = flag.String("o", "dm.csv", "Name of output file")
var invid = map[int]string{0:"AAA", 1:"BBB", 2:"CCC", 3:"DDD", 4:"EEE"}
var invnm = map[int]string{0:"Smith", 1:"Jones", 2:"Robinson", 3:"Brown", 4:"Green"}
var country = map[int]string{1:"GBR", 2:"USA", 3:"FRA", 4:"GER", 5:"SWE"}

const (
	domain 	= "DM"
	ageu	=	"Years"
)

func getInv() (string, string) {
	rand.Seed(time.Now().UTC().UnixNano())
	key := rand.Intn(len(invid))
	return invid[key], invnm[key]
}

func getCty() string{
	rand.Seed(time.Now().UTC().UnixNano())
	key := rand.Intn(len(country))
	return country[key]
}

// min age = 20, max age = 80
func getAge() int {
	rand.Seed(time.Now().UTC().UnixNano())
	return rand.Intn(59) + 20
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
	var dm []*dmrec

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
		rfstdtc, _ := time.Parse("2006-01-02", strings.Split(str, ",")[7])
		rfendtc, _ := time.Parse("2006-01-02", strings.Split(str, ",")[8])
		dmdtc, _ := time.Parse("2006-01-02", strings.Split(str, ",")[5])
		invid, invname := getInv()
		country := getCty()
		age := getAge()
		dm = append(dm, &dmrec{
			studyid: studyid,
			domain:  domain,
			usubjid: usubjid,
			subjid: subjid,
			siteid: siteid,
			rfstdtc: rfstdtc,
			rfendtc: rfendtc,
			dmdtc: dmdtc,
			invid: invid,
			invname: invname,
			country: country,
			ageu: ageu,
			age: age})
			fmt.Println(*dm[i])
	}
}
