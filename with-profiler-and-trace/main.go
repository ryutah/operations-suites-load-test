package main

import (
	"fmt"
	"log"
	"net/http"

	"cloud.google.com/go/profiler"
	"contrib.go.opencensus.io/exporter/stackdriver"
	"contrib.go.opencensus.io/exporter/stackdriver/propagation"
	"github.com/ryutah/operations-suites-load-test/pkg/env"
	"github.com/ryutah/operations-suites-load-test/pkg/frontend"
	"go.opencensus.io/plugin/ochttp"
	"go.opencensus.io/trace"
)

func main() {
	cfg := profiler.Config{
		ProjectID:      env.Appengine().ProjectID,
		Service:        env.Appengine().Service,
		ServiceVersion: env.Appengine().Version,
	}
	if err := profiler.Start(cfg); err != nil {
		panic(err)
	}

	exporter, err := stackdriver.NewExporter(stackdriver.Options{
		ProjectID: env.Appengine().ProjectID,
	})
	if err != nil {
		panic(err)
	}
	trace.RegisterExporter(exporter)
	trace.ApplyConfig(trace.Config{
		DefaultSampler: trace.AlwaysSample(),
	})

	httpHandler := &ochttp.Handler{
		Handler:     frontend.NewHandler(),
		Propagation: &propagation.HTTPFormat{},
	}
	log.Fatal(
		http.ListenAndServe(fmt.Sprintf(":%s", env.Appengine().Port), httpHandler),
	)
}
