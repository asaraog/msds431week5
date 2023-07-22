package main

import (
	"sort"
	"testing"
)

// From https://gosamples.dev/remove-duplicates-slice/
func unique(s []string) []string {
	inResult := make(map[string]bool)
	var result []string
	for _, str := range s {
		if _, ok := inResult[str]; !ok {
			inResult[str] = true
			result = append(result, str)
		}
	}
	return result
}

func TestScrape(t *testing.T) {
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

	//Scrape out Titles and Sort alphabetically
	var titles []string
	for i := range descs {
		titles = append(titles, descs[i].Title)
	}
	titles = unique(titles)
	sort.Strings(titles)

	//Read in correct titles (Ground Truth) and sort alphabetically
	pytitles := []string{"Android_(robot)", "Applications_of_artificial_intelligence", "Chatbot", "Intelligent_agent", "Reinforcement_learning", "Robot_Operating_System", "Robot", "Robotic_process_automation", "Robotics", "Software_agent"}
	sort.Strings(pytitles)

	//Checks Title matches and outputs which titles did not match
	for j := range titles {
		if titles[j] != pytitles[j] {
			t.Errorf("Error Scraping Title %v %v", pytitles[j], titles[j])
		}
	}

	//Scrape out URLs and Sort alphabetically
	var urls []string
	for k := range urls {
		urls = append(urls, descs[k].Url)
	}
	urls = unique(urls)
	sort.Strings(urls)

	//Read in correct URLs (Ground Truth) and sort alphabetically
	sort.Strings(pagesToScrape)

	//Checks URL matches and outputs which URL did not match
	for l := range urls {
		if urls[l] != pagesToScrape[l] {
			t.Errorf("Error Scraping URL %v %v", pagesToScrape[l], urls[l])
		}
	}
}
