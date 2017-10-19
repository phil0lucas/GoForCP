package main

import (
	"fmt"
	"flag"
	"strconv"
	"github.com/phil0lucas/GoForCP/CPUtils"
)

var infile = flag.String("i", "../DM/dm2.csv", "Name of input file")

func main() {
// 	Read the input file into a struct of values
	dm := CPUtils.ReadFile(infile)
	
	for ii, v := range dm {
		fmt.Printf("%v. %s %s %s %s %s %s %s %s %s %s %s %s %s %s %s %s %s %s %s\n", 
				   ii, v.Studyid, v.Domain, v.Subjid, v.Siteid, v.Usubjid, 
				   CPUtils.Pdate2str(v.Rfstdtc),
				   CPUtils.Pdate2str(v.Rfendtc),
				   v.Dmdtc.Format("2006-01-02"),
				   v.Invid,
				   v.Invname,
				   v.Country,
				   CPUtils.Pint2str(v.Age),
				   v.Ageu,
			       CPUtils.Pdate2str(v.Brthdtc),
				   CPUtils.Ptr2str(v.Sex),
				   CPUtils.Ptr2str(v.Race),
				   CPUtils.Pint2str(v.Armcd),
				   CPUtils.Ptr2str(v.Arm),
				   strconv.Itoa(v.Dmdy),
				  )
	}
}
