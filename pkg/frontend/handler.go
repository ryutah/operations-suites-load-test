package frontend

import (
	"fmt"
	"net/http"
)

func NewHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", SayHello)
	return mux
}

func SayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World")
}
