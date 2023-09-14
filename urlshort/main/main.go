package main

import (
	"fmt"
	"net/http"

	"github.com/urlshort"
)

func main() {
	mux := defaultMux()
	//_ = mux

	// Build the MapHandler using the mux as the fallback
	// pathsToUrls := map[string]string{
	// 	"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
	// 	"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	// 	"/1":              "https://www.programiz.com/golang/pointers#google_vignette",
	// }
	//mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	//Build the YAMLHandler using the mapHandler as the
	//fallback
	yaml := `
- path: /google
  url: https:www.google.com
- path: /yaml
  url: https://godoc.org/gopkg.in/yaml.v2
  `
	yamlHandler, err := urlshort.YAMLHandler([]byte(yaml), mux)
	if err != nil {
		//panic(err)
	}
	if yamlHandler == nil {
		fmt.Println(" nil")
	}
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", yamlHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	mux.HandleFunc("/redirect", redirect)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}

//redirect to new url
func redirect(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://www.google.com", http.StatusSeeOther)
}
