package main

import (
    "io/ioutil"
    "encoding/json"
    "net/http"
    "html/template"
    "strings"
    "log"
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

var defaultHandlerTmpl = `
<!DOCTYPE html>
<html>
  <head>
    <meta charset="utf-8">
    <title>Choose Your Own Adventure</title>
  </head>
  <body>
    <section class="page">
      <h1>{{.Title}}</h1>
      {{range .Paragraphs}}
        <p>{{.}}</p>
      {{end}}
      {{if .Options}}
        <ul>
        {{range .Options}}
          <li><a href="/{{.Chapter}}">{{.Text}}</a></li>
        {{end}}
        </ul>
      {{else}}
        <h3>The End</h3>
      {{end}}
    </section>
    <style>
      body {
        font-family: helvetica, arial;
      }
      h1 {
        text-align:center;
        position:relative;
      }
      .page {
        width: 80%;
        max-width: 500px;
        margin: auto;
        margin-top: 40px;
        margin-bottom: 40px;
        padding: 80px;
        background: #FFFCF6;
        border: 1px solid #eee;
        box-shadow: 0 10px 6px -6px #777;
      }
      ul {
        border-top: 1px dotted #ccc;
        padding: 10px 0 0 0;
        -webkit-padding-start: 0;
      }
      li {
        padding-top: 10px;
      }
      a,
      a:visited {
        text-decoration: none;
        color: #6295b5;
      }
      a:active,
      a:hover {
        color: #7792a2;
      }
      p {
        text-indent: 1em;
      }
    </style>
  </body>
</html>`



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
    tpl := template.Must(template.New("").Parse(defaultHandlerTmpl))

    return func(w http.ResponseWriter, r *http.Request) {
        path := strings.TrimSpace(r.URL.Path)
        if path == "" || path == "/" {
            path = "/intro"
        }
        path = path[1:]

        if story, ok := stories[path]; ok {
            err := tpl.Execute(w, story)
            if err != nil {
                log.Printf("%v", err)
                http.Error(w, "something went wrong ...", http.StatusInternalServerError)
            }
            return
        }

        http.Error(w, "Story not found.", http.StatusNotFound)
    }
}



