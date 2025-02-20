package bills

import (
	"fmt"

	"github.com/renanmedina/energy-costs-metrics-service/internal/bills/providers"
	"github.com/renanmedina/energy-costs-metrics-service/utils"
)

type ParseBillFile struct {
	provider providers.EnergyCompanyProvider
	logger   *utils.ApplicationLogger
}

func (uc ParseBillFile) Execute(filepath string) {
	uc.logger.Info(fmt.Sprintf("Starting parsing bill with %s provider for file %s", uc.provider.CompanyName(), filepath))
	billFile := NewBillFile(filepath)
	uc.logger.Info(fmt.Sprintf("Reading bill file %s", filepath))
	fileText, err := billFile.Read()

	if err != nil {
		uc.logger.Error(err.Error())
		return
	}

	parsedInfo, err := uc.provider.ParseBillFileText(fileText)

	if err != nil {
		uc.logger.Error(err.Error())
		return
	}

	uc.logger.Info("Parsed successfully", "parsed_info", parsedInfo)
	uc.logger.Info(fmt.Sprintf("Finished parsing bill with %s provider for file %s", uc.provider.CompanyName(), filepath))
}

func NewParseBillFile(companyProvider providers.EnergyCompanyProvider) ParseBillFile {
	return ParseBillFile{
		companyProvider,
		utils.GetApplicationLogger(),
	}
}
