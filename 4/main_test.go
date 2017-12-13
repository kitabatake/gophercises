package main

import (
	"testing"
	"fmt"
)

func TestParseHTML(t *testing.T) {
	result, err := parseHTML("<div><a href='google.com'>aaa<span>hogepuni</span></a><p>aa</p></div>")
	if err!= nil {
		t.Fatalf("failed test %#v", err)
	}

	fmt.Printf("%+v\n", result)
}