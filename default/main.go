package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ryutah/operations-suites-load-test/pkg/env"
	"github.com/ryutah/operations-suites-load-test/pkg/frontend"
)

func main() {
	log.Fatalln(http.ListenAndServe(
		fmt.Sprintf(":%s", env.Appengine().Port), frontend.NewHandler(),
	))
}
