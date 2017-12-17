package main

import (
	"flag"
	"os"
	"fmt"
	"net/http"
	"log"
	"strings"
)

func main() {
	storiesJsonFilePath := flag.String("json", "stories.json", "stories JSON file path")
	flag.Parse()

	storiesJsonFile, err := os.Open(*storiesJsonFilePath)
	if err != nil {
		fmt.Printf("%s isn't exists.\n", *storiesJsonFilePath)
		os.Exit(1)
	}
	stories, err := parseJson(storiesJsonFile)
	if err != nil {
		panic(err)
	}

	log.Fatal(http.ListenAndServe(":8080", handler(stories)))
}

func handler(stories map[string]Story) http.HandlerFunc {
	initTpl()
	return func(writer http.ResponseWriter, request *http.Request) {
		if story, ok := stories[path(request)]; ok {
			renderTpl(writer, story)
			return
		}
		http.Error(writer, "story is not found", http.StatusNotFound)
	}
}

func path(r *http.Request) string {
	path := strings.TrimSpace(r.URL.Path)
	if path == "" || path == "/" {
		path = "/intro"
	}
	return path[1:]
}
