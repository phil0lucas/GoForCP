package CPUtils

import (
	"log"
	"time"
	"os"
	"strconv"
	"strings"
)	

//	Capital letter ensures this will be exported in the package
//	Those variables which can be missing are modelled with pointers.
type DMrec struct {
	Usubjid 	string
	Age			*int
	Sex			*string
	Race		*string
	Armcd		int
	Arm			string
	Birthdate 	*time.Time
}

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

//	From the slice of pointers to th
func UniqueTG(dm []*DMrec) []string {
	var s []string
	for _, v := range dm{
		if v.Arm != "" && !StringInSlice(v.Arm, s) {
			s = append(s, v.Arm)
		}
	}
	s = append(s, "Overall")
	return s
}

func SubsetByArm (dm []*DMrec, value string) []*DMrec {
	var subdm []*DMrec
	for _, v := range dm {
		if v.Arm == value || value == "Overall"{
			subdm = append(subdm, v)
		}
	}
	return subdm
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

func CountByTG (dm []*DMrec) map[string]int {
	m := make(map[string]int)
	m["Screened"] = len(dm)
	for _, v := range dm {
		if v.Arm != "" {
			m[v.Arm]++
		} else {
			m["SF"]++
		}
	}
	
	total := 0
	for k, v := range m{
		if k != "SF" && k != "Screened" {
			total += v
		}
	}
	m["Overall"] = total
	return m
}

//	Usubjid is displayed with leading studyid removed in
//	the style siteid-subjidr
func SiteSubj (usubjid string) string {
	sl := strings.Split(usubjid,"-")
	return strings.Join(sl[1:],"-")
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
