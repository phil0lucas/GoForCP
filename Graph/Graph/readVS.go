package Graph

import (
	"fmt"
	"bufio"
	"strings"
	"os"
	"strconv"
	"github.com/phil0lucas/GoForCP/CPUtils"
)

func ReadFile(infile *string) []*CPUtils.DMrec {
	// open the file and pass it to a Scanner object
	file, err := os.Open(*infile)
	if err != nil {
		panic(fmt.Sprintf("error opening %s: %v", *infile, err))
	}
	defer file.Close()
	
	// Pass the opened file to a scanner
	scanner := bufio.NewScanner(file)

	var dmx []*CPUtils.DMrec
	for i := 0; scanner.Scan(); i++ {
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "error reading from file:", err)
			os.Exit(3)
		}
		str := scanner.Text()
		usubjid := strings.Split(str, ",")[4]
		age := CPUtils.Str2Int(strings.Split(str, ",")[11])
		bday := CPUtils.Str2Date(strings.Split(str, ",")[13])
		sex := CPUtils.Str2Ptr(strings.Split(str, ",")[14])
		race := CPUtils.Str2Ptr(strings.Split(str, ",")[15])
		armcd, _ := strconv.Atoi(strings.Split(str, ",")[16])
		arm := strings.Split(str, ",")[17]
		dmx = append(dmx, &CPUtils.DMrec{
			Usubjid: usubjid,
			Age: age,
			Birthdate: bday,
			Sex: sex,
			Race: race,
			Armcd: armcd,
			Arm: arm,
		})
	}
	return dmx
}
