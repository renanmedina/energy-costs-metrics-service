package services

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/renanmedina/energy-costs-metrics-service/utils"
)

type PDFConverter interface {
	Convert(source_filepath string, output_filepath string) ([]string, error)
}

type PDFToImageConverter struct {
	logger *utils.ApplicationLogger
}

func (c PDFToImageConverter) Convert(source_filepath string, output_filepath string) ([]string, error) {
	output_filepath = strings.TrimSuffix(output_filepath, filepath.Ext(output_filepath))
	outputed_filename := fmt.Sprintf("%s.jpg", output_filepath)

	defer c.logger.Info("Finished converting pdf file to image", "source_filepath", source_filepath, "output_filepath", outputed_filename)

	c.logger.Info("Starting converting pdf file to image", "source_filepath", source_filepath, "output_filepath", outputed_filename)
	command := exec.Command("pdftoppm", "-jpeg", "-r", "300", "-singlefile", source_filepath, output_filepath)
	_, err := command.Output()

	if err != nil {
		c.logger.Error(fmt.Sprintf("Failed converting pdf file to image: %s", err.Error()), "source_filepath", source_filepath, "output_filepath", outputed_filename, "error", err)
		return make([]string, 0), err
	}

	c.logger.Info("Successfully converted pdf file to image", "source_filepath", source_filepath, "output_filename", outputed_filename)

	return []string{outputed_filename}, nil
}

func NewPDFToImageConverter() PDFToImageConverter {
	return PDFToImageConverter{
		utils.GetApplicationLogger(),
	}
}
