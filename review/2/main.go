package main

import (
	"net/http"
	"log"
	"fmt"
)

func main() {
	mux := defaultMux()
	//urlMap := map[string]string{
	//	"hoge": "http://google.com",
	//	"fuga": "http://yahoo.co.jp",
	//}

	yamlString := `
- path: /urlshort
  url: https://github.com/gophercises/urlshort
- path: /urlshort-final
  url: https://github.com/gophercises/urlshort/tree/solution
`

	handler := yamlRedirectHandlerFunc([]byte(yamlString), mux)
	log.Fatal(http.ListenAndServe(":8080", handler))
}


func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintln(w, "Hello World!")
}
