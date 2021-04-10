package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/gomarkdown/markdown"
)

func main() {

	if len(os.Args) < 2 {
		log.Println("usage: go run md2html [filename]")
		os.Exit(2)
	}

	file := os.Args[1]
	fileprefix := strings.TrimSuffix(file, ".md")

	content, err := ioutil.ReadFile(file)

	if err != nil {
		log.Fatalf("%s cannot be found", file)
	}

	html := markdown.ToHTML(content, nil, nil)

	htmlFileName := fileprefix + ".html"

	err2 := ioutil.WriteFile(htmlFileName, html, 0755)

	if err2 != nil {
		log.Fatalf("Couldn't write to %s", htmlFileName)
	}

	fmt.Println("HTML created at: ", htmlFileName)
}
