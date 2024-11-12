
package main
import (
	"log"
	"github.com/signintech/gopdf"
)

func main() {

	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{ PageSize: *gopdf.PageSizeA4 })
	pdf.AddPage()
	err := pdf.AddTTFFont("wts11", "ttf/wts11.ttf")
	if err != nil {log.Fatalf("error -- addfont: %v\n", err)}

	err = pdf.SetFont("wts11", "", 14)
	if err != nil {log.Fatalf("error -- setfont: %v\n", err)}
	pdf.Cell(nil, "您好")
	pdf.WritePdf("hello.pdf")
}
