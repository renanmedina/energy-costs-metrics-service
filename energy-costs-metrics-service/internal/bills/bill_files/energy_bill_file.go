package bill_files

import (
	"path/filepath"
)

type EnergyBillFile interface {
	Read() (string, error)
}

func NewEnergyBillFile(file_path string) EnergyBillFile {
	switch extension := filepath.Ext(file_path); extension {
	case ".pdf":
		return NewPDFBill(file_path)
	default:
		return NewImageBill(file_path)
	}
}
