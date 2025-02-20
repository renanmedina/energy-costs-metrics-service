package services

import (
	"fmt"

	"github.com/otiai10/gosseract/v2"
	"github.com/renanmedina/energy-costs-metrics-service/utils"
)

type OCRImageService struct {
	client *gosseract.Client
	logger *utils.ApplicationLogger
}

func (ocr OCRImageService) ExtractTextFromImage(image_filepath string) (string, error) {
	defer ocr.logger.Info("Finished OCR image file", "image_filepath", image_filepath)

	ocr.logger.Info("Starting OCR image file", "image_filepath", image_filepath)
	ocr.client.SetImage(image_filepath)
	resultText, err := ocr.client.Text()

	if err != nil {
		ocr.logger.Error(fmt.Sprintf("Failed to OCR image file: %s", err.Error()), "image_filepath", image_filepath, "error", err)
		return "", err
	}

	ocr.logger.Info("Successfully OCR image file", "image_filepath", image_filepath)
	return resultText, nil
}

func NewOCRImageService() OCRImageService {
	return OCRImageService{
		gosseract.NewClient(),
		utils.GetApplicationLogger(),
	}
}
