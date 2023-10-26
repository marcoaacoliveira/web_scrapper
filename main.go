package main

import (
	"fmt"
	"os"

	"github.com/gocolly/colly"
)

const STORAGE_PATH = "storage/"

func main() {
	c := colly.NewCollector()

	// Find and visit all links
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnResponse(SaveFile)

	c.Visit("https://g1.com.br/")
}

func SaveFile(r *colly.Response) {
	filename := STORAGE_PATH + r.Request.URL.Host + r.Request.URL.Path + ".html"

	os.MkdirAll(STORAGE_PATH+r.Request.URL.Host, 0755)
	savingPath := filename
	if r.Request.URL.Path == "/" || r.Request.URL.Path == "" {
		savingPath = STORAGE_PATH + r.Request.URL.Host + "/index.html"
	}
	err := r.Save(savingPath)

	if err != nil {
		fmt.Println(err)
	}
}
