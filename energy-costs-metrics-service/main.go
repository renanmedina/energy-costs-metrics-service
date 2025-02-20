package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/renanmedina/energy-costs-metrics-service/handlers"
)

func main() {
	ctx := context.Background()
	tracer := InitTracer()
	defer tracer.Shutdown(ctx)

	http.HandleFunc("/metrics/process-from-filepath", handlers.ProcessMetricFromFilePathHandler)
	http.HandleFunc("/metrics", handlers.PrometheusMetrics)
	http.ListenAndServe(fmt.Sprintf(":%s", GetConfigs().PORT), nil)
}
