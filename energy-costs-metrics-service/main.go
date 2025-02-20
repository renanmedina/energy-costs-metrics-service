package main

import (
	"github.com/renanmedina/energy-costs-metrics-service/internal/bills"
	"github.com/renanmedina/energy-costs-metrics-service/internal/bills/providers"
)

func main() {
	// ctx := context.Background()
	// tracer := InitTracer()
	// defer tracer.Shutdown(ctx)

	// http.HandleFunc("/metrics/process-from-filepath", handlers.ProcessMetricFromFilePathHandler)
	// http.HandleFunc("/metrics", handlers.PrometheusMetrics)
	// http.ListenAndServe(fmt.Sprintf(":%s", utils.GetConfigs().PORT), nil)

	uc := bills.NewParseBillFile(providers.NewEnelProvider())
	uc.Execute("./storage/bill_files/test.pdf")
}
