
package main
import (
	"log"
	"github.com/signintech/gopdf"
)

func main() {

	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{ PageSize: *gopdf.PageSizeA4 })
	pdf.AddPage()
	err := pdf.AddTTFFont("cal", "ttf/Calibri.ttf")
	if err != nil {log.Fatalf("error -- addfont: %v\n", err)}

	err = pdf.SetFont("cal", "", 14)
	if err != nil {log.Fatalf("error -- setfont: %v\n", err)}
	pdf.Cell(nil, "hello pdf")
	pdf.WritePdf("helloV2.pdf")
}
