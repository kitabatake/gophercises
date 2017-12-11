package main

import (
    "io/ioutil"
    "encoding/json"
    "net/http"
    "fmt"
)

type Story struct {
    Title   string   `json:"title"`
    Paragraph   []string `json:"story"`
    Options []Option `json:"options"`
}

type Option struct {
    Text string `json:"text"`
    Arc  string `json:"arc"`
}



func parseStoriesFromFile (filepath string) (map[string]Story, error) {
    storiesBytes, err := ioutil.ReadFile(filepath)
    if err != nil {
        return nil, err
    }

    var stories map[string]Story
    if err := json.Unmarshal(storiesBytes, &stories); err != nil {
        return nil, err
    }

    return stories, nil
}

func makeHandler (stories map[string]Story) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, %s", r.URL.Path)
    }
}