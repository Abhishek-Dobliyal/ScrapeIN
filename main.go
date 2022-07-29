package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"scrapein/scraper"
)

const indeedQueryURL = "https://in.indeed.com/jobs?q="

func displayJobs(jobs []scraper.Job) {
	if len(jobs) == 0 {
		fmt.Println("\nNo jobs matching the specified args found")
		return
	}
	fmt.Printf("\nTotal number of matching jobs found: %d\n", len(jobs))
	for idx, job := range jobs {
		fmt.Println(idx+1, job)
		fmt.Println()
	}
}

func main() {
	fmt.Printf("=============> Welcome to ScrapeIn <==============\n")
	// Fetch based on tags
	fetchAllByTag := flag.String("tag", "", "Fetch all jobs that match the specified tag")
	// Fetch the name of the JSON file
	jsonFileName := flag.String("name", "", "Output JSON file name")
	flag.Parse()

	if *fetchAllByTag == "" {
		fmt.Printf("\nPlease specify the '-tag' flag!")
		os.Exit(-1)
	}

	// Scrape according to the flag provided
	toScrape := indeedQueryURL + strings.ReplaceAll(*fetchAllByTag, " ", "%20")
	indeedJobs := scraper.ScrapeIndeed(toScrape)
	displayJobs(indeedJobs)

	// Save to JSON file if flag provided
	if *jsonFileName != "" {
		file, _ := json.MarshalIndent(indeedJobs, "", " ")
		_ = ioutil.WriteFile(*jsonFileName, file, 0644)
	}

}
