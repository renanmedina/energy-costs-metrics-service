package services

import (
	"fmt"
	"os/exec"

	"github.com/renanmedina/energy-costs-metrics-service/utils"
)

type PDFConverter interface {
	Convert(filepath string, output string) ([]byte, error)
}

type PDFToImageConverter struct {
	logger *utils.ApplicationLogger
}

func (c PDFToImageConverter) Convert(filepath string, output string) ([]byte, error) {
	c.logger.Info("Starting converting pdf file to image", "source_filepath", filepath, "output_filepath", output)
	command := exec.Command("pdftoppm", "-jpeg", "-r", "300", filepath, output)
	result, err := command.Output()

	if err != nil {
		c.logger.Error(fmt.Sprintf("Failed converting pdf file to image: %s", err.Error()), "source_filepath", filepath, "output_filepath", output, "error", err)
		return make([]byte, 0), err
	}

	c.logger.Info("Successfully converted pdf file to image", "source_filepath", filepath, "output_filepath", output)
	c.logger.Info("Finished converting pdf file to image", "source_filepath", filepath, "output_filepath", output)
	return result, nil
}

func NewPDFToImageConverter() PDFToImageConverter {
	return PDFToImageConverter{
		utils.GetApplicationLogger(),
	}
}
