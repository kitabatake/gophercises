package main

import (
    // "fmt"
    "log"
    "net/http"
)

func main () {
    // Parsing Yaml
    stories, err := parseStoriesFromFile("./story.json")
    if err != nil {
        log.Fatal(err)
    }

    // Make Http handler
    http.Handle("/", makeHandler(stories))
    
    // Start server
    log.Fatal(http.ListenAndServe(":8080", nil))
}