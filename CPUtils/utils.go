package CPUtils

import (
	"fmt"
	"log"
	"time"
	"os"
	"strconv"
	"strings"
	"math/rand"
)	

//	Determine if a string is within a slice of strings
func StringInSlice(a string, list []string) bool {
    for _, b := range list {
        if b == a {
            return true
        }
    }
    return false
}

// Provide a timestamp for a program execution
func TimeStamp () string {
	t := time.Now()
	return t.Format("2006-01-02 15:04:05")	
}

// Determine the current running program
// This does not work with go run <program-name.go>.
// Use go build <program-name.go> and then ./program-name
func GetCurrentProgram () string {
	ex, err := os.Executable()
    if err != nil { log.Fatal(err) }
	return ex + ".go"
}

//	Here, the input string should be in the form of an int.
//	This is used when the input string can be missing i.e. a blank
//	value and is thus converted into a nil pointer.
func Str2Int (s string) *int {
	a, err := strconv.Atoi(s)
	if err == nil {
		return &a
	} else {
		return nil
	}
}

//	Could model string mising values with a zero-length string.
//	But decided to be consistent with numeric values and model
//	with a pointer.
func Str2Ptr (s string) *string {
	if s == "" {
		return nil
	} else {
		return &s
	}
}

// String version of a date changed into a pointer to a time.Time 
// Done this way in case the string is blank and reopresents a missing value
func Str2Date (s string) *time.Time {
    if s != "" {
        d, _ := time.Parse("2006-01-02", s)
        return &d
    } else {
        return nil
    }	
}

func Pdate2str(d *time.Time) string {
	if d != nil{
		return d.Format("2006-01-02")
	} else {
		return ""
	}
}

func Pint2str(d *int) string {
	if d != nil{
		return strconv.Itoa(*d)
	} else {
		return ""
	}
}

func Ptr2str(s *string) string {
	if s != nil{
		return *s
	} else {
		return ""
	}	
}

func Pfloat2Pstr(d *float64, dec int) *string {
	if d != nil{
		v := strconv.FormatFloat(*d, 'f', dec, 64)
		return &v
	} else {
		return nil
	}
}

func Pfloat2str(d *float64, dec int) string {
	if d != nil{
		v := strconv.FormatFloat(*d, 'f', dec, 64)
		return v
	} else {
		return ""
	}
}

//	Randomly selects a variable as having a missing value
func FlagMiss (threshold float64) bool {
	
	if threshold == 0.0 {
		threshold = 0.05
	}
	
	rand.Seed(time.Now().UTC().UnixNano())
	if rand.Float64() >= threshold {
		return false
	} else {
		return true
	}
}

// Select a random member from a slice of strings
func Choice(s []string) string {
	// Allocate seed for generating random numbers
	rand.Seed(time.Now().UTC().UnixNano())
	return s[rand.Intn(len(s))]
}

// This pads the string in the 1st arg to the length
// in the 3rd arg with the char in the 2nd arg
func LeftPad2Len(s string, padStr string, overallLen int) string {
	var padCountInt int
	padCountInt = 1 + ((overallLen - len(padStr)) / len(padStr))
	var retStr = strings.Repeat(padStr, padCountInt) + s
	return retStr[(len(retStr) - overallLen):]
}

func PrintDate(t time.Time) {
	fmt.Println(t.Format("2006-01-02"))
}
// Can these similar functions be rationalized via an interface??
func PrintDateP(t *time.Time) {
	if t == nil {
		fmt.Println("Missing Value")
	} else {
		fmt.Println(t.Format("2006-01-02"))
	}
}

func PrintPint(p *int) {
	if p == nil {
		fmt.Println("Missing Value")
	} else {
		fmt.Println(*p)
	}
}

func PrintPfloat(p *float64) {
	if p == nil {
		fmt.Println("Missing Value")
	} else {
		fmt.Println(*p)
	}
}

//	Select random key and value from a map
func RandItem (m map[int]string) (int, string) {
	rand.Seed(time.Now().UTC().UnixNano())
	key := rand.Intn(len(m))
	return key, m[key]
}

//	Select random value from a map or set to missing
func RandItemP (m map[int]string) *string {
	if FlagMiss(0) == false {
		rand.Seed(time.Now().UTC().UnixNano())
		key := rand.Intn(len(m))
		value := m[key]
		return &value
	} else {
		return nil
	}
}
