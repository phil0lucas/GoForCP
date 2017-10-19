package CPUtils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type DMrec struct {
	Studyid string
	Domain  string
	Subjid  string
	Siteid  string
	Usubjid string
	Rfstdtc *time.Time
	Rfendtc *time.Time
	Dmdtc   time.Time
	Invid   string
	Invname string
	Country string
	Age     *int
	Ageu    string
	Brthdtc *time.Time
	Sex     *string
	Race    *string
	Armcd   *int
	Arm     *string
	Dmdy    int
}

func ReadFile(infile *string) []*DMrec {
	// open the file and pass it to a Scanner object
	file, err := os.Open(*infile)
	if err != nil {
		panic(fmt.Sprintf("error opening %s: %v", *infile, err))
	}
	defer file.Close()

	// Pass the opened file to a scanner
	scanner := bufio.NewScanner(file)

	var dmx []*DMrec
	for i := 0; scanner.Scan(); i++ {
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "error reading from file:", err)
			os.Exit(3)
		}
		str := scanner.Text()

		studyid := strings.Split(str, ",")[0]
		domain := strings.Split(str, ",")[1]
		subjid := strings.Split(str, ",")[2]
		siteid := strings.Split(str, ",")[3]
		usubjid := strings.Split(str, ",")[4]
		rfstdtc := Str2Date(strings.Split(str, ",")[5])
		rfendtc := Str2Date(strings.Split(str, ",")[6])
		dmdtc, _ := time.Parse("2006-01-02", strings.Split(str, ",")[7])
		invid := strings.Split(str, ",")[8]
		invname := strings.Split(str, ",")[9]
		country := strings.Split(str, ",")[10]
		age := Str2Int(strings.Split(str, ",")[11])
		ageu := strings.Split(str, ",")[12]
		bday := Str2Date(strings.Split(str, ",")[13])
		sex := Str2Ptr(strings.Split(str, ",")[14])
		race := Str2Ptr(strings.Split(str, ",")[15])
		armcd := Str2Int(strings.Split(str, ",")[16])
		arm := Str2Ptr(strings.Split(str, ",")[17])
		dmdy, _ := strconv.Atoi(strings.Split(str, ",")[18])

		dmx = append(dmx, &DMrec{
			Studyid: studyid,
			Domain:  domain,
			Subjid:  subjid,
			Siteid:  siteid,
			Usubjid: usubjid,
			Rfstdtc: rfstdtc,
			Rfendtc: rfendtc,
			Dmdtc:   dmdtc,
			Invid:   invid,
			Invname: invname,
			Country: country,
			Age:     age,
			Ageu:    ageu,
			Brthdtc: bday,
			Sex:     sex,
			Race:    race,
			Armcd:   armcd,
			Arm:     arm,
			Dmdy:    dmdy,
		})
	}
	return dmx
}
