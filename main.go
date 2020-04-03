package main

import (
	"fmt"
	"goin"
	"net/http"
)

func main() {
	g := goin.New()
	g.GET("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "URL.Path: %s\n", req.URL.Path)
	})
	g.GET("/hello", func(w http.ResponseWriter, req *http.Request) {
		for k, v := range req.Header {
			fmt.Fprintf(w, "Header[%v]=%v\n", k, v)
		}
	})
	g.Run(":9999")
}
