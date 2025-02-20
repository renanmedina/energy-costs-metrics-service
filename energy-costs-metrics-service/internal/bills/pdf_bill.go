package bills

type PDFBill struct {
	filepath string
}

func (f PDFBill) Read() (string, error) {
	return "", nil
}
