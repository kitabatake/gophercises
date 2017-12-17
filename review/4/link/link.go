package link

import (
	"io"
	"golang.org/x/net/html"
	"strings"
)

type Link struct {
	Href string
	Text string
}

func Parse(reader io.Reader) []Link {
	var links []Link

	doc, _ := html.Parse(reader)
	aNodes := extractANodes(doc)
	for _, aNode := range aNodes {
		links = append(links, buildLink(aNode))
	}
	return links
}

func extractANodes(node *html.Node) []*html.Node {
	if node.Type == html.ElementNode && node.Data == "a" {
		return []*html.Node {node}
	}

	var nodes []*html.Node
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		nodes = append(nodes, extractANodes(c)...)
	}
	return nodes
}

func buildLink(node *html.Node) Link {
	var link Link
	for _, attr := range node.Attr {
		if attr.Key == "href" {
			link.Href = attr.Val
		}
	}
	link.Text = strings.Join(extractTexts(node), " ")
	return link
}

func extractTexts(node *html.Node) []string {
	if (node.Type == html.TextNode) {
		return []string{strings.Replace(node.Data, "\n", " ", len(node.Data))}
	}
	var texts []string
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		texts = append(texts, extractTexts(c)...)
	}
	return texts
}