package main

import (
	"fmt"
	"net/http"
	"github.com/ramyaDhanush/learning-go/gophercises/urlshort"
)

func main() {
	mux := defaultMux()
	
	pathToUrls := map[string]string {
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	
	mapHandler := urlshort.MapHandler(pathToUrls, mux)
	
	// yaml := `
	// -path: /urlshort
	//  url: https://github.com/gophercises/urlshort
	// -path: /urlshort-final
	//  url: https://github.com/gophercises/urlshort/tree/final
	// `
	// yamlHandler, err := urlshort.YAMLHandler([]byte(yaml), mapHandler)
	
	// if err != nil {
	// 	panic(err)
	// }
	
	fmt.Println("Starting the server on: 8080")
	
	http.ListenAndServe(":8080", mapHandler)
	
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World!!")
}