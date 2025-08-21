package main

import (
	"fmt"
	"net/http"
)

type Engine struct{}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/index":
		fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)

	case "/hello":
		for k, v := range req.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}

	default:
		fmt.Fprintf(w, "404 Not Found\n")
	}
}

func main() {
	engine := new(Engine)
	http.ListenAndServe(":8080", engine)
}
