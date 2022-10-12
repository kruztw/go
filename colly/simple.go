package main

import (
	"fmt"
	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector()

	// tag
	c.OnHTML("title", func(e *colly.HTMLElement) {
		fmt.Printf("title: %v\n", e.Text)
	})

	// self-defined tag (meta with name)
	c.OnHTML("meta[name]", func(e *colly.HTMLElement) {
		fmt.Printf("meta[name]: %v\n", e)
	})

	// self-defined tag (meta with name='theme-color')
	c.OnHTML("meta[name='theme-color']", func(e *colly.HTMLElement) {
		fmt.Printf("content: %v\n", e.Attr("content"))
	})

	// CSS
	c.OnHTML(".GettingStartedGo-headerH2", func(e *colly.HTMLElement) {
		fmt.Printf("css: %v\n", e.Text)
	})

	// ID
	c.OnHTML("#wtf", func(e *colly.HTMLElement) {
		fmt.Printf("id: ", e.Text)
	})

	c.OnResponse(func(r *colly.Response) {
		//fmt.Println(string(r.Body))
	})

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.75 Safari/537.36")
	})

	c.Visit("https://go.dev/") 
}
