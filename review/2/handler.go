package main

import (
	"net/http"
	"gopkg.in/yaml.v2"
	"fmt"
)

func mapRedirectHandleFunc(m map[string]string, fallback *http.ServeMux) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		if url, ok := m[request.URL.Path]; ok {
			http.Redirect(writer, request, url, http.StatusTemporaryRedirect)
		}
		fallback.ServeHTTP(writer, request)
	}
}

func yamlRedirectHandlerFunc(yamlBytes []byte, fallback *http.ServeMux) http.HandlerFunc {
	pathUrls, _ := parseYaml(yamlBytes)
	pathUrlMap := map[string]string{}
	for _, pathUrl := range pathUrls {
		pathUrlMap[pathUrl.Path] = pathUrl.Url
	}
	return mapRedirectHandleFunc(pathUrlMap, fallback)
}

func parseYaml(yamlBytes []byte) ([]pathUrl, error) {
	var pathUrls []pathUrl
	err := yaml.Unmarshal(yamlBytes, &pathUrls)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return pathUrls, nil
}

type pathUrl struct {
	Path string `yaml:"path"`
	Url string `yaml:"url"`
}