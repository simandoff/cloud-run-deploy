package main

import (
	"fmt"
	"net/http"
	"os"
)

func httpRoot(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "<h1>Home of APIs!</h1>")
}

func httpHeaders(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

const (
	pathStatic = "static"
	portVar    = "PORT"
	portValue  = "8080"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/test", httpRoot)
	mux.HandleFunc("/api/h", httpHeaders)

	fs := http.FileServer(http.Dir("./" + pathStatic))
	mux.Handle("/", fs)
	
	panic(http.ListenAndServe(":"+getEnvDef(portVar, portValue), mux))
}

func getEnvDef(env string, def string) (r string) {
	r = os.Getenv(env)
	if r == "" {
		r = def
	}
	return r
}
