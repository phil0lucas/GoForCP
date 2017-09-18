package SummaryReport

import (
	"os"
	"fmt"
	"bufio"
	"strings"
	"strconv"
)

type tAge *int

type DMrec struct {
	usubjid string
	Age		tAge
	sex		string
	race	string
	armcd	int
	Arm		string
}

func setAge (s string) tAge {
	a, err := strconv.Atoi(s)
	if err == nil {
		return &a
	} else {
		return nil
	}
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
		usubjid := strings.Split(str, ",")[4]
		age := setAge(strings.Split(str, ",")[11])
		sex := strings.Split(str, ",")[14]
		race := strings.Split(str, ",")[15]
		armcd, _ := strconv.Atoi(strings.Split(str, ",")[16])
		arm := strings.Split(str, ",")[17]
		dmx = append(dmx, &DMrec{
			usubjid: usubjid,
			Age: age,
			sex: sex,
			race: race,
			armcd: armcd,
			Arm: arm,
		})
	}
	return dmx
}
