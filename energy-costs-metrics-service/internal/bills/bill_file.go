package bills

type BillFile interface {
	Read() (string, error)
}

func NewBillFile(filepath string) BillFile {
	return PDFBill{filepath}
}
