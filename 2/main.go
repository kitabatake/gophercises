package main

import (
    "fmt"
    "net/http"
    yaml "gopkg.in/yaml.v2"
)

func main () {
    mux := defaultMux()
    // pathsToUrls := map[string]string{
    //     "/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
    //     "/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
    // }
    // mapHandler := mapHandler(pathsToUrls, mux)

    yamlText := `
- path: /urlshort
  url: https://github.com/gophercises/urlshort
- path: /urlshort-final
  url: https://github.com/gophercises/urlshort/tree/solution
`
    yamlHandler, err := yamlHandler([]byte(yamlText), mux)
    if err != nil {
        panic(err)
    }

    fmt.Println("Starting the server on :8080")
    http.ListenAndServe(":8080", yamlHandler)
}

func defaultMux() *http.ServeMux {
    mux := http.NewServeMux()
    mux.HandleFunc("/", hello)
    return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello, world!")
}



func mapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        if url, ok := pathsToUrls[r.URL.Path]; ok {
            http.Redirect(w, r, url, http.StatusFound)
            return
        }
        fallback.ServeHTTP(w, r)
    }
}

func yamlHandler(yamlBytes []byte, fallback http.Handler) (http.HandlerFunc, error) {
    var pathUrls []pathUrl
    err := yaml.Unmarshal(yamlBytes, &pathUrls)
    if err != nil {
        return nil, err
    }

    fmt.Println(string(yamlBytes))
    fmt.Println(pathUrls)
    var pathToUrls map[string]string
    for _, pathUrl := range(pathUrls) {
        pathToUrls[pathUrl.path] = pathUrl.url
    }

    fmt.Println(pathToUrls)

    return mapHandler(pathToUrls, fallback), nil
}

type pathUrl struct {
    path string `yaml:"path"`
    url string `yaml:"url"`
}