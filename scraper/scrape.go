package scraper

import (
	"strings"

	"github.com/gocolly/colly"
)

type Job struct {
	Title     string `json:"title"`
	Company   string `json:"company"`
	Location  string `json:"location"`
	Other     string `json:"other"`
	ApplyLink string `json:"link"`
}

var c = colly.NewCollector(colly.AllowedDomains(
	"indeed.com",
	"https://in.indeed.com/",
	"in.indeed.com",
))

func ScrapeIndeed(requestURL string) []Job {
	var jobs []Job

	currJob := Job{}

	c.OnHTML("td[class=resultContent]", func(h *colly.HTMLElement) {
		jobTitle := strings.TrimSpace(h.ChildText("div > h2 > a > span"))
		company := strings.Title(h.ChildText("span[class=companyName]"))
		location := strings.TrimSpace(h.ChildText("div[class=companyLocation]"))
		others := strings.TrimSpace(h.ChildText("div[class=attribute_snippet]"))
		applyLink := strings.TrimSpace(h.ChildAttr("div > h2 > a[role=button]", "href"))
		applyLink = h.Request.AbsoluteURL(applyLink)

		currJob.Title = jobTitle
		currJob.Company = company
		currJob.Location = location
		currJob.Other = others
		currJob.ApplyLink = applyLink

		jobs = append(jobs, currJob)
	})

	c.OnHTML("a[aria-label=Next]", func(h *colly.HTMLElement) {
		nextPage := h.Request.AbsoluteURL(h.Attr("href"))
		c.Visit(nextPage)
	})

	c.Visit(requestURL)

	return jobs
}
