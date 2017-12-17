package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
)

type Story struct {
	Title   string   `json:"title"`
	Paragraphs   []string `json:"story"`
	Options []Option `json:"options"`
}

type Option struct {
	Text string `json:"text"`
	Chapter  string `json:"arc"`
}

func parseJson(reader io.Reader) (map[string]Story, error) {
	var stories map[string]Story
	storiesBytes, _ := ioutil.ReadAll(reader)
	err := json.Unmarshal(storiesBytes, &stories)
	if err != nil {
		return nil, err
	}
	return stories, nil
}