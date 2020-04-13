package main

func main() {
	LogInfo("Fetching results...")
	doc := FetchAndStorePage("https://www.kayak.it/flights/AMS-FCO/2020-07-05?sort=price_a&fs=stops=~0", "response.html")

	LogInfo("Parsing results...")
	results := ParseKayakResults(doc)

	LogInfo("Storing JSON results...")
	StoreJSONResults(results, "data/results.json")

	LogInfo("Building results template...")
	mailTemplate := BuildResultsHTMLTemplate(results)

	LogInfo("Sending results...")
	SendEmailNotification(&SMTPServer{}, []string{"antonellodippolito@gmail.com"}, mailTemplate)

	LogInfo("Done.")
}
