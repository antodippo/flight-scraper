package main

import (
	"strings"

	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
)

// FlightResult is a representation of a single result
type FlightResult struct {
	Departure     string
	Arrival       string
	DepartureTime string
	ArrivalTime   string
	Price         string
	Airline       string
}

// ParseKayakResults parses results in an html page
func ParseKayakResults(doc *html.Node) (flightResults []FlightResult) {
	resultNodes := htmlquery.Find(doc, "//*[@class=\"Base-Results-HorizonResult Flights-Results-FlightResultItem phoenix-rising get-in-formation extra-vertical-spacing hover-actions get-in-formation sleek rp-contrast \"]")
	for _, node := range resultNodes {
		flightResult := FlightResult{
			Departure:     getNthElementInDoc(node, "//*[@class=\"bottom-airport js-airport \"]", 1),
			Arrival:       getNthElementInDoc(node, "//*[@class=\"bottom-airport js-airport \"]", 2),
			DepartureTime: getNthElementInDoc(node, "//*[@class=\"depart-time base-time\"]", 1),
			ArrivalTime:   getNthElementInDoc(node, "//*[@class=\"arrival-time base-time\"]", 1),
			Price:         getNthElementInDoc(node, "//*[@class=\"price-text\"]", 1),
			Airline:       getNthElementInDoc(node, "//*[@class=\"section times\"]/*[@class=\"bottom \"]", 1),
		}
		flightResults = append(flightResults, flightResult)
	}

	return
}

func getNthElementInDoc(doc *html.Node, xpath string, nth int) (element string) {
	list := htmlquery.Find(doc, xpath)
	if len(list) == 0 {
		return ""
	}
	element = strings.Trim(htmlquery.InnerText(list[nth-1]), "\t \n")
	element = stripCtlAndExtFromUTF8(element)

	return
}

func stripCtlAndExtFromUTF8(str string) string {
	return strings.Map(func(r rune) rune {
		if r >= 32 && r < 127 {
			return r
		}
		return -1
	}, str)
}
