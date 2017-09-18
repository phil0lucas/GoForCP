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


// Need to change this to set a small random percentage of values to 'missing'.
// Missing will need to be a specific value per type e.g. -999 for an age.
// Perhaps this also implies a type of 'Age' which implements a method such
// as 'setMissing'.
// If multiple variables of different types need to be set in this way, 
// then can define an interface as a way of collecting the types which can be 
// set in this way.

package main

import (
    //"linux_amd64/github.com/phil0lucas/GoForCP/InjectMV"
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

type birthdate time.Time
type age int
type sex string
type race string

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
	age		*age
	brthdtc *birthdate
	sex		*sex
	race	*race
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
var sexmp = map[int]string{0:"M", 1:"F"}
var racemp = map[int]string{0:"White", 1:"Black", 2:"Asian"}
var arm = map[int]string{0:"Placebo", 1:"Active"}

const (
	domain 	= "DM"
	ageu	= "Years"
	dmdy	= 0
)

// Randomly select the investigator ID and name 
func getInv() (string, string) {
	rand.Seed(time.Now().UTC().UnixNano())
	key := rand.Intn(len(invid))
	return invid[key], invnm[key]
}

// Randomly select country from a slice
func getCty() string{
	rand.Seed(time.Now().UTC().UnixNano())
	key := rand.Intn(len(country))
	return country[key]
}

func flagMiss () bool {
	rand.Seed(time.Now().UTC().UnixNano())
	if rand.Float64() >= 0.05 {
		return false
	} else {
		return true
	}
}

// min age = 20, max age = 80
func getAge() *age {
	if flagMiss() == false {
		rand.Seed(time.Now().UTC().UnixNano())
		r := age(rand.Intn(59) + 20)
		return &r
	} else {
		return nil
	}
}

// Randomly assign a gender to the subject
func getSex() *sex {
	if flagMiss() == false {
		rand.Seed(time.Now().UTC().UnixNano())
		s := sex(sexmp[rand.Intn(len(sexmp))])
		return &s
	} else {
		return nil
	}
}

// Randomly assign a race to the subject
func getRace() *race {
	if flagMiss() == false {	
		rand.Seed(time.Now().UTC().UnixNano())
		r := race(racemp[rand.Intn(len(racemp))])
		return &r
	} else {
		return nil
	}
}

// Generate a birth date based on the recorded age
func getBday(dmdtc time.Time, age *age) *birthdate {
	// Birth date is recorded at screening, which is DMDTC here.
	// Having randomly generated an age, calculate the last possible 
	// birthday at that age and then subtract a random number
	// of days between 0 and 364
	if age != nil {
		v := int(*age)
		_bdate := dmdtc.AddDate(-v, 0, 0)
		rand.Seed(time.Now().UTC().UnixNano())
		offset := rand.Intn(364)
		b := birthdate(_bdate.AddDate(0,0,-offset))
		return &b
	} else {
		return nil
	}
}

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
		brthdtc := getBday(dmdtc, age)
		sex := getSex()
		race := getRace()
		armcd, _ := strconv.Atoi(strings.Split(str, ",")[9])
		arm := strings.Split(str, ",")[10]
		
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
// 			fmt.Println(*dm[i])
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
}
