package providers

type EletricBillInfo struct{}

type EnergyCompanyProvider interface {
	CompanyName() string
	FetchFileUrl() string
	ParseBillFile(filepath string) (EletricBillInfo, error)
	ParseBillFileText(fileText string) (EletricBillInfo, error)
}
