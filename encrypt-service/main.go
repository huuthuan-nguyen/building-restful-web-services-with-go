package main

import (
	"log"
	"net/http"
	httptransport "github.com/go-kit/kit/transport/http"
	"encryptservice/helpers"
	kitlog "github.com/go-kit/kit/log"
	"os"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
)

func main() {
	// init logger
	logger := kitlog.NewLogfmtLogger(os.Stderr)

	// instrumenting metrics
	fieldKeys := []string{"method", "error"}
	requestCount := kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
		Namespace: "encryption",
		Subsystem: "my_service",
		Name: "request_count",
		Help: "Number of requests received.",
	}, fieldKeys)
	requestLatency := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "encryption",
		Subsystem: "my_service",
		Name: "request_latency_microseconds",
		Help: "Total duration of requests in microseconds.",
	}, fieldKeys)

	// encrypt service
	var svc helpers.EncryptService
	svc = helpers.EncryptServiceInstance{}

	// attach logging middleware
	svc = helpers.LoggingMiddleware{Logger: logger, Next: svc}

	// attach instrumenting midlleware
	svc = helpers.InstrumentingMiddleware{RequestCount: requestCount, RequestLatency: requestLatency, Next: svc}
	encryptHandler := httptransport.NewServer(helpers.MakeEncryptEndpoint(svc), helpers.DecodeEncryptRequest, helpers.EncodeResponse)
	decryptHandler := httptransport.NewServer(helpers.MakeDecryptEndpoint(svc), helpers.DecodeDecryptRequest, helpers.EncodeResponse)

	http.Handle("/encrypt", encryptHandler)
	http.Handle("/decrypt", decryptHandler)

	// handle metrics route
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}