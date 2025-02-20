package bill_files

import (
	"github.com/renanmedina/energy-costs-metrics-service/internal/bills/services"
)

type EnergyImageBill struct {
	filepath   string
	ocrService services.OCRImageService
}

func (f EnergyImageBill) Read() (string, error) {
	extractedText, err := f.ocrService.ExtractTextFromImage(f.filepath)

	if err != nil {
		return "", err
	}

	return extractedText, nil
}

func NewImageBill(filepath string) EnergyImageBill {
	return EnergyImageBill{
		filepath,
		services.NewOCRImageService(),
	}
}
