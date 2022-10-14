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
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("./static"))

	mux.Handle("/static", http.StripPrefix("/static", fs))

	mux.HandleFunc("/", httpRoot)
	mux.HandleFunc("/headers", httpHeaders)
	mux.HandleFunc("/h", httpHeaders)
	
	panic(http.ListenAndServe(":"+os.GetEnvDev("PORT", "8080"), mux))
}


func (o os) GetEnvDef(env string, def string) (r string) {
	r = o.Getenv(env)
	if r == "" {
		r = def
	}
	return r
}
