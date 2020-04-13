package main

import (
	"os"
	"time"
)

// SearchInput is a representation of the input request
type SearchInput struct {
	Departure string
	Arrival   string
	Date      string
}

func main() {
	searchInput, recipient := ParseFlags(os.Args[0], os.Args[1:])
	url := "https://www.kayak.it/flights/" + searchInput.Departure + "-" + searchInput.Arrival + "/" + searchInput.Date + "?sort=price_a&fs=stops=~0"
	now := time.Now().Format("20060102-150405")

	LogInfo("Fetching results...")
	doc := FetchAndStorePage(url, "data/"+now+"_response.html")

	LogInfo("Parsing results...")
	results := ParseKayakResults(doc)
	if len(results) == 0 {
		LogWarning("No results parsed.")
		os.Exit(0)
	}

	LogInfo("Storing JSON results...")
	StoreJSONResults(results, "data/"+now+"_results.json")

	LogInfo("Building results template...")
	mailTemplate := BuildResultsHTMLTemplate(results, searchInput)

	LogInfo("Sending results...")
	SendEmailNotification(&SMTPServer{}, recipient, mailTemplate, searchInput)

	LogInfo("Done.")
}
