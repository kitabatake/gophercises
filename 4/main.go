package main

import (
	"golang.org/x/net/html"
	"strings"
	)

type Link struct {
	Href string
	Text string
}

func parseHTML(htmlText string) ([]Link, error) {
	node, err := html.Parse(strings.NewReader(htmlText))
	if err != nil {
		return nil, err
	}

	var links []Link
	parseNode(node, &links)
	return links, nil
}

func parseNode(node *html.Node, links *[]Link) {
	if node.Type == html.ElementNode && node.Data == "a" {
		var texts []string
		extractTexts(node, &texts)
		*links = append(*links, Link {
			Href: extractHref(node.Attr),
			Text: strings.Join(texts, " "),
		})
	} else {
		for c := node.FirstChild; c != nil; c = c.NextSibling {
			parseNode(c, links)
		}
	}
}

func extractHref(attributes []html.Attribute) string {
	href := ""
	for _, attr := range attributes {
		if attr.Key == "href" {
			href = attr.Val
		}
	}
	return href
}

func extractTexts(node *html.Node, texts *[]string) {
	if node.Type == html.TextNode {
		*texts = append(*texts, node.Data)
	}
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		extractTexts(c, texts)
	}
}