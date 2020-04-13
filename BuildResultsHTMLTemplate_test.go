package main

import (
	"io/ioutil"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildResultsHTMLTemplate(t *testing.T) {
	t.Run("It parses and returns flights results", func(t *testing.T) {
		fileContent, _ := ioutil.ReadFile("data_test/test_results_html")
		searchInput := SearchInput{"FCO", "RTM", "2020-02-02"}
		want := trimWhitespaces(string(fileContent))
		got := trimWhitespaces(string(BuildResultsHTMLTemplate(getResultList(), searchInput)))
		assert.Equal(t, want, got)
	})
}

func trimWhitespaces(initialString string) (strippedString string) {
	strippedString = strings.Replace(initialString, " ", "", -1)
	strippedString = strings.Replace(strippedString, "\n", "", -1)

	return
}

func getResultList() []FlightResult {
	return []FlightResult{
		FlightResult{
			Departure:     "AMS",
			Arrival:       "FCO",
			DepartureTime: "16:30",
			ArrivalTime:   "18:50",
			Price:         "69",
			Airline:       "LEVEL",
		},
		FlightResult{
			Departure:     "AMS",
			Arrival:       "FCO",
			DepartureTime: "7:20",
			ArrivalTime:   "9:45",
			Price:         "123",
			Airline:       "KLM",
		},
	}
}
