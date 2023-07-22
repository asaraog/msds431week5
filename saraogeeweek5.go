package main

import (
	"encoding/json"
	"os"
	"strings"

	"github.com/gocolly/colly"
)

// Output file type and fields for URL, Title and Text for each page
type JSONoutput struct {
	Url   string `json:"url"`
	Title string `json:"title"`
	Text  string `json:"Text"`
}

func main() {
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
		"https://en.wikipedia.org/wiki/Android_(robot)",
	}

	//Perform Scraping
	descs := Scrape(pagesToScrape)
	//Marshal description into JSON encoding
	marshaldescs, _ := json.Marshal(descs)
	//Write JSON line file
	os.WriteFile("items.jl", marshaldescs, 0666)
}

// Function for scraping and testing
func Scrape(pagesToScrape []string) (descs []JSONoutput) {
	//Make seperate directory to save .html files
	os.Mkdir("wikipages", 0700)

	//HTTP client creation using colly
	c := colly.NewCollector(
		colly.AllowedDomains("en.wikipedia.org"),
	)

	//Parsing Logic for each HTML page for URL, Title and Text
	//Initialize empty file for output
	descs = make([]JSONoutput, 0)
	//Wikipedia pages have a .mw-parser-output style heading in the Body
	c.OnHTML(".mw-parser-output", func(e *colly.HTMLElement) {
		//This seperates the URL using '/' to seperate the last part of the URL for Title
		sepURL := strings.Split(e.Request.URL.String(), "/")

		desc := JSONoutput{
			Url: e.Request.URL.String(),
			//This takes the last value after the URL is split for the page Title
			Title: sepURL[len(sepURL)-1],
			Text:  e.Text,
		}
		descs = append(descs, desc)
		e.Response.Save("./wikipages/" + sepURL[len(sepURL)-1] + ".html")
	})

	//Scraping each URL
	for _, pageToScrape := range pagesToScrape {
		c.Visit(pageToScrape)
	}
	c.Wait() //Waits to ensure time to visit each URL

	return descs
}
