package main

import (
	"../../link"
	"flag"
	"os"
	"fmt"
)

func main() {
	htmlFilePath := flag.String("html", "example1.html", "html file path")
	flag.Parse()

	htmlFile, err := os.Open(*htmlFilePath)
	if err != nil {
		fmt.Println("file isn't exists.")
		os.Exit(1)
	}
	links := link.Parse(htmlFile)
	fmt.Printf("%#v", links)
}
