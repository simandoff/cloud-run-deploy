package main

import (
	"fmt"
	"net/http"
	"os"
)

func httpRoot(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "<h1>Home of index!</h1>")
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

	panic(http.ListenAndServe(":"+getDefEnv("PORT", "8080"), nil))
}

func getDefEnv(env string, def string) (res string) {
	res = os.Getenv(env)
	if res == "" {
		res = def
	}
	return
}
