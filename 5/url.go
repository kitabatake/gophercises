package main

import (
	"regexp"
)

type URL struct {
	original string
}

func (url *URL) extractDomain() string {
	re := regexp.MustCompile(`^https?://([^/]+)`)
	m := re.FindStringSubmatch(url.original)
	if m != nil {
		return m[1]
	}
	return ""
}

func (url *URL) isSameDomain(href string) bool {
	domain := url.extractDomain()
	m, _ := regexp.MatchString(`https?://` + domain + `/`, href)

	// TODO: evalutes beginning at /
	return m
}

func (url *URL) normalizeURL(domain string) string {
	if m,_ := regexp.MatchString(`^/`, url.original); m {
		return "https://" + domain + url.original
	}
	return url.original
}