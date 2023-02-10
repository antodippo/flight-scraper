package main

import (
	"os"
	"strings"
	"time"
)

// SearchInput is a representation of the input request
type SearchInput struct {
	Departure string
	Arrival   string
	Date      string
}

const dataDir = "flight-data/"

func init() {
	if _, err := os.Stat(dataDir); os.IsNotExist(err) {
		err = os.Mkdir(dataDir, os.ModeDir|0755)
		LogFatalAndExitIfNotNull(err)
	}
}

func main() {
	searchInput, recipient := ParseFlags(os.Args[0], os.Args[1:])
	url := "https://www.kayak.it/flights/" + searchInput.Departure + "-" + searchInput.Arrival + "/" + searchInput.Date + "?sort=price_a&fs=stops=~0"
	now := time.Now().Format("20060102-150405")

	LogInfo("Fetching results...")
	doc := FetchAndStorePage(url, dataDir+now+"_response.html")

	LogInfo("Parsing results...")
	results := ParseKayakResults(doc)
	if len(results) == 0 {
		LogWarning("No results parsed.")
		os.Exit(0)
	}

	LogInfo("Storing JSON results...")
	StoreJSONResults(results, dataDir+now+"_results.json")

	LogInfo("Storing JSON queries...")
	StoreJSONQueries(searchInput, recipient, dataDir+now+"_results.json")

	LogInfo("Building results template...")
	mailTemplate := BuildResultsHTMLTemplate(results, searchInput)

	LogInfo("Sending results...")
	SendEmailNotification(&SMTPServer{}, strings.Split(recipient, ","), mailTemplate, searchInput)

	LogInfo("Done.")
}
