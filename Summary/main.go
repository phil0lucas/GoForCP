// 

package main

import (
//	"github.com/jung-kurt/gofpdf"
	"github.com/phil0lucas/GoForCP/Summary"
	"fmt"
	"bufio"
	"flag"
	"strings"
	"os"
	"strconv"
)

type dmrec struct {
	usubjid string
	age		int
	sex		string
	race	string
	armcd	int
	arm		string
}

var infile = flag.String("i", "../DM/dm.csv", "Name of input file")
//var outfile = flag.String("o", "summary.pdf", "Name of output file")

func readFile(infile *string) dmrecs {
	// open the file and pass it to a Scanner object
	file, err := os.Open(*infile)
	if err != nil {
		panic(fmt.Sprintf("error opening %s: %v", *infile, err))
	}
	defer file.Close()
	
	// Pass the opened file to a scanner
	scanner := bufio.NewScanner(file)

	var dmx dmrecs
	for i := 0; scanner.Scan(); i++ {
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "error reading from file:", err)
			os.Exit(3)
		}
		str := scanner.Text()
		usubjid := strings.Split(str, ",")[4]
		age, _ := strconv.Atoi(strings.Split(str, ",")[11])
		sex := strings.Split(str, ",")[14]
		race := strings.Split(str, ",")[15]
		armcd, _ := strconv.Atoi(strings.Split(str, ",")[16])
		arm := strings.Split(str, ",")[17]
		dmx = append(dmx, &dmrec{
			usubjid: usubjid,
			age: age,
			sex: sex,
			race: race,
			armcd: armcd,
			arm: arm,
		})
	}
	return dmx
}

//func countd(dm dmrecs), element string, by string) map[string]int {
func countd(dm dmrecs) map[bigN]int {
	m := make(map[bigN]int)
	for _, v := range dm {
		m[bigN{v.arm}]++
		m[bigN{"All"}]++
	}
	return m
}

func nmiss(dm dmrecs) map[bigN]int {
	m := make(map[bigN]int)
	for _, v := range dm {
		if v.age != -999 {
            m[bigN{v.arm}]++
            m[bigN{"All"}]++
        }
	}
	return m    
}




func main() {
	var dm dmrecs
	// Read the file and dump into the slice of structs
	dm = readFile(infile)
    // Determine the numbers of subjects by TG including an 'All' level
    
    /*
	m := countd(dm)
    fmt.Println(m)
    // Determine number of non-missing values of Age per TG
    n := nmiss(dm)
	fmt.Println(n)
	*/
}


	/*
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, "Hello, world")
	err := pdf.OutputFileAndClose("hello.pdf")
	fmt.Println(err)
	*/
