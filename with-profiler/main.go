package main

import (
	"fmt"
	"log"
	"net/http"

	"cloud.google.com/go/profiler"

	"github.com/ryutah/operations-suites-load-test/pkg/env"
	"github.com/ryutah/operations-suites-load-test/pkg/frontend"
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

	log.Fatalln(http.ListenAndServe(
		fmt.Sprintf(":%s", env.Appengine().Port), frontend.NewHandler(),
	))
}
