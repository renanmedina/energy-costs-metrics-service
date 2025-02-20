package providers

import "github.com/renanmedina/energy-costs-metrics-service/utils"

type EnelProvider struct {
	logger *utils.ApplicationLogger
}

func (eprovider EnelProvider) CompanyName() string {
	return "Enel"
}

func (eprovider EnelProvider) FetchFileUrl() string {
	return ""
}

func (eprovider EnelProvider) ParseBillFileText(fileText string) (EletricBillInfo, error) {
	return EletricBillInfo{}, nil
}

func (eprovider EnelProvider) ParseBillFile(filepath string) (EletricBillInfo, error) {
	return EletricBillInfo{}, nil
}

func NewEnelProvider() EnelProvider {
	return EnelProvider{
		utils.GetApplicationLogger(),
	}
}
