# Crawling and Scraping the Web

## Project Summary

Go will be help power our backend web and database servers and distributed service offerings on the cloud. However, data science operations remain a key concern as Python remain popular. This project aims to implement web crawling and scraping of Wikipedia [(described by Chanda 2021)](https://www.scrapingbee.com/blog/web-scraping-go/#building-a-basic-scraper) webpages in Go using [Colly](https://go-colly.org/). The Go implementation is benchmarked for runtime using 'time' before commands in the command line to compare with [Python's implementation for the same 10 webpages using scrapy](./WebFocusedCrawlWorkV001). The results are subsequently written to a JSONL (items.jl) file based on previous implementations ([Div Rhino 2020](https://divrhino.com/articles/build-webscraper-with-go-and-colly/)).

Python was significantly slower compared Go implementations with 'real' runtimes of 15.9s and 0.6s for Python and Go respectively. While Python was less verbose than Go, Go is more scalable and has concurrency support to allow for even faster processing using [Colly](https://go-colly.org/docs/examples/parallel/). The Data Science team sees test-driven development as an asset in Go and with the difference in processing time, we strongly recommend using Go as the primary programming language accross the company.

## Files

*saraogeeweek5.go:* \
Main routine loads descriptions from webpages using the Scrape function built using Colly with output as a .jl (JSONL) file with URL, Title and Text fields. It also saves a .html of the webpage in a new directory ['wikipages'](./wikipages).

*saraogeeweek5_test.go:* \
Unit test for Scrape function. This testing routine ensures equivalence with expected Titles and URLs from chosen websites. It uses a function by others to remove duplicates [(Gosamples 2022)](https://gosamples.dev/remove-duplicates-slice/).

*items.jl* \
JSONL output file with URL, Title and Text fields for each website on a different line.

*Week5* \
Unix executable file of cross-compiled Go code for Mac/Windows. 

## Installation

Download or git clone this project onto local machine into folder on local machine.
```
git clone https://github.com/asaraog/msds431week5.git
cd msds431week5
time ./Week5

cd WebFocusedCrawlWorkV001
time python3 run-articles-spider.py
```

## References
Chanda, Subha. 2021. “Web Scraping with Go.” ScrapingBee. 2021. https://www.scrapingbee.com/blog/web-scraping-go/. \
Colly Team. 2018. “Gocolly Package.” Colly. 2018. https://go-colly.org/. \
Div Rhino. 2020. “How to Build a Web Scraper with Go and Colly.” DivRhino. 2020. https://divrhino.com/articles/build-webscraper-with-go-and-colly/. \
Gosamples. 2022. “Remove Duplicates from a Slice in Go (Golang).” February 4, 2022. https://gosamples.dev/remove-duplicates-slice/.

