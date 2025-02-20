package bills

import "github.com/renanmedina/energy-costs-metrics-service/internal/bills/services"

type BillFile interface {
	Read() (string, error)
}

func NewBillFile(filepath string) BillFile {
	return PDFBill{
		filepath,
		services.NewPDFToImageConverter(),
	}
}
