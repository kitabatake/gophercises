package main

import (
    "fmt"
    "log"
)

func main () {
    // Parsing Yaml
    stories, err := parseStoriesFromFile("./story.json")
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(stories["intro"].Title)

    // Adjust to struct

    // Make Http handler

    // Start server
}