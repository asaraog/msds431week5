package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/gocolly/colly"
)

type JSONoutput struct {
	Url   string `json:"url"`
	Title string `json:"title"`
	Text  string `json:"Text"`
}

func main() {
	os.Mkdir("wikipages", 0700)
	pagesToScrape := []string{
		"https://en.wikipedia.org/wiki/Robotics",
		"https://en.wikipedia.org/wiki/Robot",
		"https://en.wikipedia.org/wiki/Reinforcement_learning",
		"https://en.wikipedia.org/wiki/Robot_Operating_System",
		"https://en.wikipedia.org/wiki/Intelligent_agent",
		"https://en.wikipedia.org/wiki/Software_agent",
		"https://en.wikipedia.org/wiki/Robotic_process_automation",
		"https://en.wikipedia.org/wiki/Chatbot",
		"https://en.wikipedia.org/wiki/Applications_of_artificial_intelligence",
		"htts://en.wikipedia.orgwiki/Android_(robot)",
	}

	//HTTP client creation
	c := colly.NewCollector(
		colly.AllowedDomains("en.wikipedia.org"),
	)

	//Parsing Logic
	descs := make([]JSONoutput, 0)
	c.OnHTML(".mw-parser-output", func(e *colly.HTMLElement) {
		sepURL := strings.Split(e.Request.URL.String(), "/")
		desc := JSONoutput{
			Url:   e.Request.URL.String(),
			Title: sepURL[len(sepURL)-1],
			Text:  e.Text,
		}
		descs = append(descs, desc)
		e.Response.Save("./wikipages/" + sepURL[len(sepURL)-1] + ".html")
	})

	//Scraping each URL
	for _, pageToScrape := range pagesToScrape {
		c.Visit(pageToScrape)
		c.OnError(func(_ *colly.Response, err error) {
			fmt.Println("Bad url:", err)
		})
	}
	c.Wait()

	//Export Logic
	Marshaldescs, _ := json.Marshal(descs)
	os.WriteFile("items.jl", Marshaldescs, 0666)
}
