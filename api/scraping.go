package api

import (
	"fmt"

	"github.com/gocolly/colly"
)

func GetTitle(url string) string {
	getExplore()
	c := colly.NewCollector()

	result := ""
	//ここでタイトルを取得
	c.OnHTML("title", func(e *colly.HTMLElement) {
		fmt.Println(e.Text)
		result = e.Text
	})

	err := c.Visit(url)
	if err != nil {
		fmt.Println(err)
	}

	return result
}

func getExplore() {

	c := colly.NewCollector()
	i := 0
	print("Explore")
	c.OnHTML("article", func(e *colly.HTMLElement) {

		i++
		book := e.DOM.Find("a > h3").Text()
		author := e.DOM.Find("div > a").Text()
		if author == "" {
			author = e.ChildText(".BookLink_userName__avtjq")
		}

		fmt.Printf("%d 著者: %s / タイトル: %s\n", i, author, book)
	})

	err := c.Visit("https://zenn.dev/books/explore")
	if err != nil {
		fmt.Println(err)
	}

}
