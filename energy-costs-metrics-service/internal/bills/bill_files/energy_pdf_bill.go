package bill_files

import (
	"fmt"
	"path/filepath"

	"github.com/renanmedina/energy-costs-metrics-service/internal/bills/services"
)

type EnergyPDFBill struct {
	filepath     string
	pdfConverter services.PDFConverter
}

func (f EnergyPDFBill) Read() (string, error) {
	imagePaths, err := f.pdfConverter.Convert(f.filepath, fmt.Sprintf("storage/bill_files/converted/%s", filepath.Base(f.filepath)))

	if err != nil {
		return "", err
	}

	imageBill := NewImageBill(imagePaths[0])
	return imageBill.Read()
}

func NewPDFBill(file_path string) EnergyPDFBill {
	return EnergyPDFBill{
		file_path,
		services.NewPDFToImageConverter(),
	}
}
