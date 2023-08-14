package service

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

func ScrapeMain() map[string]string {
	data := make(map[string]string, 1)

	c := colly.NewCollector()

	c.OnHTML(`a[href]`, func(h *colly.HTMLElement) {
		links := h.Attr("href")

		if !strings.Contains(links, "://") {
			data[strings.ToLower(h.Text)] = fmt.Sprintf("%s%s", "https://gobyexample.com/", links)
		}
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.Visit("https://gobyexample.com/")

	return data
}

func ScrapeSubMain(lq string) (data string) {
	c := colly.NewCollector()

	c.OnHTML(`div`, func(h *colly.HTMLElement) {
		data = h.Text
	})

	c.Visit(lq)

	return data
}
