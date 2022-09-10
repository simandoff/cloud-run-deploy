package main

import (
	"fmt"
	"net/http"
	"os"
)

func httpRoot(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Home of Acuzio.watch!")
}

func httpHeaders(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {
	http.HandleFunc("/", httpRoot)
	http.HandleFunc("/headers", httpHeaders)

	panic(http.ListenAndServe(":"+getDefaultEnv("PORT", "8080"), nil))
}

func getDefaultEnv(env string, def string) (v string) {
	v = os.Getenv(env)
	if v == "" {
		return def
	}
	return v
}
