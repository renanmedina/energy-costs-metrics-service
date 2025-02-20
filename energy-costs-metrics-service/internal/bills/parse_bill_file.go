package bills

import (
	"fmt"

	"github.com/renanmedina/energy-costs-metrics-service/internal/bills/bill_files"
	"github.com/renanmedina/energy-costs-metrics-service/internal/bills/providers"
	"github.com/renanmedina/energy-costs-metrics-service/utils"
)

type ParseBillFile struct {
	provider providers.EnergyCompanyProvider
	logger   *utils.ApplicationLogger
}

func (uc ParseBillFile) Execute(filepath string) {
	defer uc.logger.Info(fmt.Sprintf("Finished parsing bill with %s provider for file %s", uc.provider.CompanyName(), filepath))

	uc.logger.Info(fmt.Sprintf("Starting parsing bill with %s provider for file %s", uc.provider.CompanyName(), filepath))
	billFile := bill_files.NewEnergyBillFile(filepath)
	uc.logger.Info(fmt.Sprintf("Reading bill file %s", filepath))
	fileText, err := billFile.Read()

	if err != nil {
		uc.logger.Error(err.Error())
		return
	}

	parsedInfo, err := uc.provider.ParseBillFileText(fileText)

	if err != nil {
		uc.logger.Error(fmt.Sprintf("Failed to parse bill file with %s provider for file %s: %s", uc.provider.CompanyName(), filepath, err.Error()), "error", err)
		return
	}

	uc.logger.Info(fmt.Sprintf("Successfully parsed bill file with %s provider for file %s", uc.provider.CompanyName(), filepath), "parsed_info", parsedInfo)
}

func NewParseBillFile(companyProvider providers.EnergyCompanyProvider) ParseBillFile {
	return ParseBillFile{
		companyProvider,
		utils.GetApplicationLogger(),
	}
}
