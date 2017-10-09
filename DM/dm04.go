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

// 	Screening Failure subjects will be included and have missing values
//	for some of their data fields
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
	"flag"
	"math/rand"
// 	"strconv"
// 	"log"
	
	"github.com/phil0lucas/GoForCP/CPUtils"
)

// This will mirror the metadata above with more natural types
type dmrec struct {
	studyid string
	domain  string
	usubjid string
	subjid  string
	siteid  string
	rfstdtc *time.Time
	rfendtc *time.Time
	dmdtc   time.Time
	invid	string
	invname string
	country string
	ageu	string
	age		*int
	brthdtc *time.Time
	sex		*string
	race	*string
	armcd	*int
	arm		*string
	dmdy    int
}

// The program will be run with flags to specify the input & output files
var infile = flag.String("i", "../SC/sc2.csv", "Name of input file")
var outfile = flag.String("o", "dm2.csv", "Name of output file")

// Various lookups for random selection
var invid = map[int]string{0:"AAA", 1:"BBB", 2:"CCC", 3:"DDD", 4:"EEE"}
var invnm = map[int]string{0:"Smith", 1:"Jones", 2:"Robinson", 3:"Brown", 4:"Green"}
var ctrymap = map[int]string{0:"GBR", 1:"USA", 2:"FRA", 3:"GER", 4:"SWE"}
var sexmp = map[int]string{0:"M", 1:"F"}
var racemp = map[int]string{0:"White", 1:"Black", 2:"Asian"}

const (
	domain 	= "DM"
	ageu	= "Years"
	dmdy	= 0
)

/*
// Randomly select the investigator ID and name 
func getInv() (string, string) {
	rand.Seed(time.Now().UTC().UnixNano())
	key := rand.Intn(len(invid))
	return invid[key], invnm[key]
}
*/




// min age = 20, max age = 80
func getAge() *int {
	if CPUtils.FlagMiss(0) == false {
		rand.Seed(time.Now().UTC().UnixNano())
		r := rand.Intn(59) + 20
		return &r
	} else {
		return nil
	}
}

// Generate a birth date based on the recorded age
func getBday(dmdtc time.Time, age *int) *time.Time {
	// Birth date is recorded at screening, which is DMDTC here.
	// Having randomly generated an age, calculate the last possible 
	// birthday at that age and then subtract a random number
	// of days between 0 and 364
	if age != nil {
		v := *age
		_bdate := dmdtc.AddDate(-v, 0, 0)
		rand.Seed(time.Now().UTC().UnixNano())
		offset := rand.Intn(364)
		b := _bdate.AddDate(0,0,-offset)
		return &b
	} else {
		return nil
	}
}

/*
func (s *birthdate) dateFmt() string {
    // The receiver s is a pointer to a RefEndDate
    // If the value of the pointer is not nil then
    // it points to a variable
    if s != nil {
        d := time.Time(*s)
        return d.Format("2006-01-02")
    } else {
        return ""
    }
}

func ageToString(a *age) string {
	if a != nil {
		return strconv.Itoa(int(*a))
	} else {
		return ""
	}
}

func sexToString(a *sex) string {
	if a != nil {
		return string(*a)
	} else {
		return ""
	}
}

func raceToString(a *race) string {
	if a != nil {
		return string(*a)
	} else {
		return ""
	}
}
*/

func main() {
	flag.Parse()
	
	// open the file and pass it to a Scanner object
	file, err := os.Open(*infile)
	if err != nil {
		panic(fmt.Sprintf("error opening %s: %v", infile, err))
	}
	defer file.Close()

	// Output slice of pointers to structs
// 	var dm []*dmrec

	// Pass the opened file to a scanner
	scanner := bufio.NewScanner(file)

	for i := 0; scanner.Scan(); i++ {
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "error reading from file:", err)
			os.Exit(3)
		}
		str := scanner.Text()
		studyid := strings.Split(str, ",")[0]
		fmt.Println(studyid)
		usubjid := strings.Split(str, ",")[3]
		fmt.Println(usubjid)
		subjid := strings.Split(str, ",")[1]
		fmt.Println(subjid)
		siteid := strings.Split(str, ",")[2]
		fmt.Println(siteid)
		
		// Screening date
		dmdtc, _ := time.Parse("2006-01-02", strings.Split(str, ",")[5])
		CPUtils.PrintDate(dmdtc)
		
		// First date of dosing for randomized subjects.
		// For screening failures this will be a nil pointer.
		rfstdtc := CPUtils.Str2Date(strings.Split(str, ",")[7])
		fmt.Printf("%T\n", rfstdtc)
		//	Last day of dosing.
		//	This will also be missing if the subject is a screening failure
		rfendtc := CPUtils.Str2Date(strings.Split(str, ",")[8])
		fmt.Printf("%T\n", rfendtc)
		
		
		iKey, invid := CPUtils.RandItem(invid)
		invname := invnm[iKey]
		fmt.Printf("%s, %s\n", invid, invname)
		_, country := CPUtils.RandItem(ctrymap)
		fmt.Printf("Country code: %s\n", country)
		age := getAge()
		CPUtils.PrintPint(age)
		brthdtc := getBday(dmdtc, age)
		CPUtils.PrintDateP(brthdtc)
		
		sex := CPUtils.RandItemP(sexmp)
		fmt.Println(CPUtils.Ptr2str(sex))
		race := CPUtils.RandItemP(racemp)
		fmt.Println(CPUtils.Ptr2str(race))		
		armcd := CPUtils.Str2Int(strings.Split(str, ",")[9])
		CPUtils.PrintPint(armcd)
		arm := strings.Split(str, ",")[10]
		fmt.Println(arm)
// 		
// 		dm = append(dm, &dmrec{
// 			studyid: studyid,
// 			domain:  domain,
// 			usubjid: usubjid,
// 			subjid: subjid,
// 			siteid: siteid,
// 			rfstdtc: rfstdtc,
// 			rfendtc: rfendtc,
// 			dmdtc: dmdtc,
// 			invid: invid,
// 			invname: invname,
// 			country: country,
// 			ageu: ageu,
// 			age: age,
// 			brthdtc: brthdtc,
// 			sex: sex,
// 			race: race,
// 			armcd: armcd,
// 			arm: arm,
// 			dmdy: dmdy})
// // 			fmt.Println(*dm[i])
	}
	
	// Injection of some 'missing values' into:
	// age (int)
	// race (string)
	// brthdtc (time.Time)
	// For each type a substitute value is chosen dependent upon
	// the type and range of values it may contain. (Is this the best / only way??)
	
// 	rand.Seed(time.Now().UTC().UnixNano())
// 	for ii, _ := range dm{
//         randno := rand.Intn(100)
//         //fmt.Println(randno)
//         if randno < 3 {
//             dm[ii].age = -999
//             // dm[ii].brthdtc = How to set date literal ??
//         }
//         randnoc := rand.Intn(100)
//         //fmt.Println(randno)
//         if randnoc < 3 {
//             dm[ii].race = "" 
//         }
//     }
	
/*	
// 	// Output file writing section
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
			ageToString(dm[ii].age) +  "," + 
			dm[ii].ageu +  "," + 
			dm[ii].brthdtc.dateFmt() + "," +
			sexToString(dm[ii].sex) +  "," +
			raceToString(dm[ii].race) +  "," +
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
*/
}
