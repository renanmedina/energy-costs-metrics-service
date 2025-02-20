# energy-costs-metrics-service

A repository to ocr and parse energy bills and serve metrics to prometheus

### Dependencies
- Nodejs v22.14.0 (I advise to use a node version manager like [nvm](https://github.com/nvm-sh/nvm))
- [puppeteer](https://pptr.dev/)
- pdftoppm (sudo apt-get install poppler-utils)
- AWS S3 Bucket