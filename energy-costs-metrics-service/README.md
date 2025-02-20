# energy-costs-metrics-service

A repository to ocr and parse energy bills and serve metrics to prometheus

## Dependencies
- Golang 1.24
- [Tesseract OCR](https://github.com/tesseract-ocr/tesseract): using the C linked library libtesseract-dev into golang program
- [qpdf](https://github.com/qpdf/qpdf): Used to handle encrypted pdf files if needed
- pdftoppm (sudo apt-get install poppler-utils)
- AWS S3 Bucket

## Running the project

#### Docker
We provide a Dockerfile prepared to be builted which already installs the dependencies for you
```bash
docker run --build
```

#### Manually
Install the tesseract dependency before running the code

```bash
apt-get update && apt-get install -y tesseract-ocr tesseract-ocr-por libtesseract-dev poppler-utils
```

Run the application
```bash
go run main.go
```