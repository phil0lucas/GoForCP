// 

package ReportSummary

import (
	"github.com/jung-kurt/gofpdf"
	"fmt"
)

func writeTitle(outputFile string, title1 string) {
	pdf := gofpdf.New("L", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Courier", "", 12)
	pdf.Cell(10, 10, title1)
	err := pdf.OutputFileAndClose(outputFile)
	fmt.Println(err)
}
