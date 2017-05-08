// Program to generate SDTM data for a fictitious study.
// Domain DM
// Metadata :
// - STUDYID Char 6  (constant) Study Identifier
// - DOMAIN  Char 2  (constant) Domain abbreviation
// - USUBJID Char 18 STUDYID-SITEID-SUBJID Unique Subject Identifier
// - SUBJID  Char 6  Subject Identifier
// - SITEID  Char 4  Site Identifier
// - RFSTDTC Date 10 ISO8601 First date of study med exposure
// - RFENDTC Date 10 ISO8601 Last date of study med exposure
// - DMDTC   Date 10 ISO8601 Date/Time of Collection
// - INVID   Char 3  Investigator code
// - INVNAME Char 8  Investigator Name
// - COUNTRY Char 3  ISO3166 Country code
// - BRTHDTC Date 10 ISO8601 Subjects date of birth
// - AGE	 Num     Subject's age (min 20, Max 80)
// - AGEU    Char 5  (constant) Age units
// - SEX     Char 1  Subject's gender ((M/F)
// - RACE    Char 5  Subject's race (White, Black, Asian)
// - ARMCD   Num     Treatment Arm code
// - ARM     Char 7  Treatment Arm
// - DMDY    Num     Study Day of collection

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
	"log"
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
	brthdtc time.Time
	sex		string
	race	string
	armcd	int
	arm		string
	dmdy    int
}

// The program will be run with flags to specify the input & output files
var infile = flag.String("i", "../SC/sc.csv", "Name of input file")
var outfile = flag.String("o", "dm.csv", "Name of output file")
// Various lookups for random selection
var invid = map[int]string{0:"AAA", 1:"BBB", 2:"CCC", 3:"DDD", 4:"EEE"}
var invnm = map[int]string{0:"Smith", 1:"Jones", 2:"Robinson", 3:"Brown", 4:"Green"}
var country = map[int]string{0:"GBR", 1:"USA", 2:"FRA", 3:"GER", 4:"SWE"}
var sex = map[int]string{0:"M", 1:"F"}
var race = map[int]string{0:"White", 1:"Black", 2:"Asian"}
var arm = map[int]string{0:"Placebo", 1:"Active"}

const (
	domain 	= "DM"
	ageu	= "Years"
	dmdy	= 0
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

func getSex() string {
	rand.Seed(time.Now().UTC().UnixNano())
	return sex[rand.Intn(len(sex))]
}

func getRace() string {
	rand.Seed(time.Now().UTC().UnixNano())
	return race[rand.Intn(len(race))]	
}

func getArm() (int, string) {
	rand.Seed(time.Now().UTC().UnixNano())
	armcd := rand.Intn(len(arm))
	arm := arm[armcd]
	return armcd, arm
}

func getBday(dmdtc time.Time, age int) time.Time {
	// Birth date is recorded at screening, which is DMDTC here.
	// Having randomly generated an age, calculate the last possible 
	// birthday at that age and then subtract a random number
	// of days between 0 and 364
	_bdate := dmdtc.AddDate(-age, 0, 0)
	rand.Seed(time.Now().UTC().UnixNano())
	offset := rand.Intn(364)
	return _bdate.AddDate(0,0,-offset)
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
		
		// Need to trap the error output.
		rfstdtc, _ := time.Parse("2006-01-02", strings.Split(str, ",")[7])
		rfendtc, _ := time.Parse("2006-01-02", strings.Split(str, ",")[8])
		dmdtc, _ := time.Parse("2006-01-02", strings.Split(str, ",")[5])
		
		invid, invname := getInv()
		country := getCty()
		age := getAge()
		race := getRace()
		
		//_bdate := dmdtc.AddDate(-age, 0, 0).Format("2006-01-02")
		//_dmdtc := dmdtc.Format("2006-01-02")
		//_age := strconv.Itoa(age)
		//fmt.Println("Last possible birthdate: " + _bdate + "Age " + _age + " at " + _dmdtc)
		brthdtc := getBday(dmdtc, age)
		sex := getSex()
		armcd, arm := getArm()
		
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
			age: age,
			brthdtc: brthdtc,
			sex: sex,
			race: race,
			armcd: armcd,
			arm: arm,
			dmdy: dmdy})
			fmt.Println(*dm[i])
	}
	
	fo, err := os.Create(*outfile)
	if err != nil {
		log.Fatal(err)
	}
	defer fo.Close()
	
	// Create a buffered writer from the file
	w := bufio.NewWriter(fo)
	
	for ii, _ := range dm{
		bytesWritten, err := w.WriteString(
			dm[ii].studyid + "," +
			dm[ii].domain + "," +
			dm[ii].subjid + "," +
			dm[ii].siteid + "," +
			dm[ii].usubjid +  "," +
			dm[ii].rfstdtc.Format("2006-01-02") + "," +
			dm[ii].rfendtc.Format("2006-01-02") + "," +
			dm[ii].dmdtc.Format("2006-01-02") + "," +
			dm[ii].invid +  "," + 
			dm[ii].invname +  "," +
			dm[ii].country +  "," +
			strconv.Itoa(dm[ii].age) +  "," + 
			dm[ii].ageu +  "," + 
			dm[ii].brthdtc.Format("2006-01-02") + "," +
			dm[ii].sex +  "," +
			dm[ii].race +  "," +
			strconv.Itoa(dm[ii].armcd) +  "," + 
			dm[ii].arm +  "," +
			strconv.Itoa(dm[ii].dmdy) +
			"\n")

		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Bytes written: %d\n", bytesWritten)
	}

	// Write to disk
	w.Flush()		
}