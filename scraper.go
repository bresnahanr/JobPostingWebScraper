package main

import (
	//import Colly
	"fmt"

	"github.com/gocolly/colly"
)

func main() {

	collector := colly.NewCollector()
	var jobPostings []JobPosting

	collector.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting: ", r.URL)
	})

	collector.OnError(func(_ *colly.Response, err error) {
		fmt.Println("something went wrong: ", err)
	})

	collector.OnResponse(func(r *colly.Response) {
		fmt.Println("Page visited: ", r.Request.URL)
	})

	collector.OnHTML("a.lf-career-card", func(section *colly.HTMLElement) {
		//collector.OnHTML("section.css-y0qu21", func(section *colly.HTMLElement) {

		fmt.Println("Found css-y0qu21 section")

		jobPosting := JobPosting{}

		jobPosting.url = section.Attr("href")
		jobPosting.title = section.ChildText("h3")

		jobPostings = append(jobPostings, jobPosting)

	})

	collector.OnScraped(func(r *colly.Response) {
		fmt.Println(r.Request.URL, " scraped!")
	})

	collector.Visit("https://www.strathcona.ca/council-county/careers/")

	for _, job := range jobPostings {
		fmt.Println("URL: ", job.url)
		fmt.Println("Title: ", job.title)
	}
}

type JobPosting struct {
	url, title, description string
}
