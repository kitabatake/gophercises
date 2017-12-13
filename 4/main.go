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

	anodes := extractANodes(node)
	var links []Link
	for _, anode := range anodes {
		links = append(links, buildLink(anode))
	}
	return links, nil
}

func buildLink(node *html.Node) Link {
	var link Link
	link.Href = extractHref(node.Attr)
	link.Text = strings.Join(extractTexts(node), " ")
	return link
}

func extractANodes(node *html.Node) []*html.Node {
	if node.Type == html.ElementNode && node.Data == "a" {
		return []*html.Node {node}
	}
	var ret []*html.Node
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		ret = append(ret, extractANodes(c)...)
	}
	return ret
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

func extractTexts(node *html.Node) []string {
	if node.Type == html.TextNode {
		return []string {node.Data}
	}
	var ret []string
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		ret = append(ret, extractTexts(c)...)
	}
	return ret
}