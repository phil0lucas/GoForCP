// 

package main

import (
//	"github.com/jung-kurt/gofpdf"
	"fmt"
	"bufio"
	"flag"
	"strings"
	"os"
	"strconv"
//	"reflect"
)

type dmrec struct {
	usubjid string
	age		int
	sex		string
	race	string
	armcd	int
	arm		string
}

type Key struct {
	arm, sex	string
}

// Slice of pointers to structs - 
//	one element per input record
type dmrecs []*dmrec

// The program will be run with flags to specify the input & output files
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
func countd(dm dmrecs) map[Key]int {
	m := make(map[Key]int)
	for _, v := range dm {
		m[Key{v.arm, v.sex}]++
		m[Key{"All",v.sex}]++
	}
	return m
}


func main() {
	var dm dmrecs
	// fmt.Printf("Type of infile object %T\n", infile)
	dm = readFile(infile)
	// fmt.Println(*dm[99] )
	// fmt.Printf("Type of dm object %T\n", dm)
	m := countd(dm)
	fmt.Println(m )
}


	/*
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, "Hello, world")
	err := pdf.OutputFileAndClose("hello.pdf")
	fmt.Println(err)
	*/