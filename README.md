# Crawling and Scraping the Web

## Project Summary

Go will be help power our backend web and database servers and distributed service offerings on the cloud. However, data science operations remain a key concern as Python remain popular. This project aims to implement web crawling and scraping of 10 webpages in Go using the [Colly](github.com/gocolly/colly). The Go implementation is benchmarked for runtime using 'time' before commands in the command line to compare with [Python's](./WebFocusedCrawlWorkV001) scrapy package. The results are subsequently written to a JSONL (.jl) file.

Python was significantly faster compared Go implementations with 'real' runtimes of 2.27s, 4.10s, 5.02s for Python and Go respectively. While Python was less verbose and a bit faster than Go, Go is more scalable and has concurrency support to allow for parallel processing using [Colly](https://go-colly.org/docs/examples/parallel/). The Data Science team sees test-driven development as an asset in Go and with equivalent web  crawling and scraping, we strongly recommend using Go as the primary programming language accross the company.

## Files

*saraogeeweek5.go:* \
Main routine loads descriptions from webpages using the Scrape function built using [Colly](github.com/gocolly/colly) with output as a .jl (JSONL) file with URL, Title and Text fields. It also saves a .html of the webpage in a new directory 'wikipages'.

*saraogeeweek5_test.go:* \
Unit test for Scrape function. This testing routine ensures equivalence with expected Titles and URLs from chosen websites.

*Week5* \
Unix executable file of cross-compiled Go code for Mac/Windows. 

## Installation

Download or git clone this project onto local machine into folder on local machine.

```
git clone https://github.com/asaraog/msds431week5.git
cd msds431week5
time ./Week5

cd ..
cd WebFocusedCrawlWorkV001
time python3 run-articles-spider.py
```
## References

Miller, Thomas. 2015. “Modeling Techniques in Predictive Analytics Chapter 3.” 2015. https://github.com/mtpa/mtpa/tree/master/MTPA_Chapter_3. \
Saha, Amit. 2022. Practical Go: Building Scalable Network & Non-network Applications. New York: Wiley. ISBN-13: 978-1-119-77381-8. Chapter 3 Writing HTTP Clients (pages 57–80) and Chapter 4 Advanced HTTP Clients (pages 81–104).
