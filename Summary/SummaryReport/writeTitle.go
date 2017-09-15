// 

package SummaryReport

import (
	"github.com/jung-kurt/gofpdf"
	"fmt"
// 	"os"
// 	"time"
// 	"path/filepath"
// 	"log"
)

type headers struct {
	line1a	string
}

// func WriteReport(outputFile *string, title1 string, title2 string,
// 				title3 string, title4 string, title5 string, title6 string) error {
					
func WriteReport(outputFile *string, h *headers) error {					
					
	pdf := gofpdf.New("L", "mm", "A4", "")
	pdf.SetHeaderFunc(func() {
		pdf.SetFont("Courier", "", 10)
		pdf.CellFormat(0, 10, (*h).line1a, "0", 0, "L", false, 0, "")
// 		pdf.CellFormat(0, 10, "CONFIDENTIAL", "0", 0, "R", false, 0, "")
// 		pdf.Ln(4)
// 		pdf.CellFormat(0, 10, title2, "0", 0, "L", false, 0, "")
// 		pdf.CellFormat(0, 10, "Draft", "0", 0, "R", false, 0, "")
// 		pdf.Ln(4)
// 		pdf.CellFormat(0, 10, title3, "0", 0, "L", false, 0, "")
// 		pdf.Ln(4)
// 		pdf.CellFormat(0, 10, title4, "0", 0, "C", false, 0, "")
// 		pdf.Ln(4)
// 		pdf.CellFormat(0, 10, title5, "0", 0, "C", false, 0, "")
// 		pdf.Ln(4)
// 		pdf.CellFormat(0, 10, title6, "0", 0, "C", false, 0, "")
// 		pdf.Ln(10)
	})
	
// 	pdf.SetFooterFunc(func() {
// 		pdf.SetY(-15)
// 		pdf.SetFont("Courier", "", 10)
// 		pdf.CellFormat(0, 10, fmt.Sprintf("Page %d of {nb}", pdf.PageNo()),
// 			"", 0, "L", false, 0, "")
// 		
// 		dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
// 		if err != nil {
//             log.Fatal(err)
// 		}
// 		pdf.CellFormat(0, 10, dir, "", 0, "C", false, 0, "")	
// 		t := time.Now()
// 		tt:= t.Format("2006-01-02 15:04:05")
// 		pdf.CellFormat(0, 10, "Run: " + tt, "", 0, "R", false, 0, "")			
// 	})
// 	pdf.AliasNbPages("")
	
// 	// Simple table
// 	basicTable := func() {
// 		colHeader := []string{"Characteristic", "Statistic", "Placebo", "Active", "Overall"}
// 		for _, str := range colHeader {
// 			pdf.CellFormat(55, 8, str, "TB", 0, "L", false, 0, "")
// 		}
// 		pdf.Ln(-1)
// 		for _, c := range countryList {
// 			pdf.CellFormat(40, 6, c.nameStr, "1", 0, "", false, 0, "")
// 			pdf.CellFormat(40, 6, c.capitalStr, "1", 0, "", false, 0, "")
// 			pdf.CellFormat(40, 6, c.areaStr, "1", 0, "", false, 0, "")
// 			pdf.CellFormat(40, 6, c.popStr, "1", 0, "", false, 0, "")
// 			pdf.Ln(-1)
// 		}
// 	}	

	pdf.AddPage()
// 	basicTable()
// 
	err := pdf.OutputFileAndClose(*outputFile)
	fmt.Println(err)
	return err
} 
