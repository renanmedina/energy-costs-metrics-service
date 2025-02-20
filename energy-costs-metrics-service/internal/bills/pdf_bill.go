package bills

import (
	"fmt"
	"path/filepath"

	"github.com/renanmedina/energy-costs-metrics-service/internal/bills/services"
)

type PDFBill struct {
	filepath     string
	pdfConverter services.PDFConverter
}

func (f PDFBill) Read() (string, error) {
	imagePaths, err := f.pdfConverter.Convert(f.filepath, fmt.Sprintf("./storage/bill_files/converted/%s", filepath.Base(f.filepath)))

	if err != nil {
		return "", err
	}

	fmt.Println(imagePaths)

	return "", nil
}
