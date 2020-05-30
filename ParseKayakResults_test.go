package main

import (
	"testing"

	"github.com/antchfx/htmlquery"
	"github.com/stretchr/testify/assert"
)

func TestParseKayakResults(t *testing.T) {
	t.Run("It parses and returns flights results", func(t *testing.T) {
		doc, _ := htmlquery.LoadDoc("data-test/test_response")
		got := ParseKayakResults(doc)
		want := getWantedResultList()
		assert.Equal(t, want, got)
	})
}

func getWantedResultList() []FlightResult {
	return []FlightResult{
		{
			Departure:     "AMS",
			Arrival:       "FCO",
			DepartureTime: "16:30",
			ArrivalTime:   "18:50",
			Price:         "69",
			Airline:       "LEVEL",
		},
		{
			Departure:     "AMS",
			Arrival:       "FCO",
			DepartureTime: "7:20",
			ArrivalTime:   "9:45",
			Price:         "69",
			Airline:       "LEVEL",
		},
		{
			Departure:     "AMS",
			Arrival:       "FCO",
			DepartureTime: "12:30",
			ArrivalTime:   "14:55",
			Price:         "104",
			Airline:       "easyJet",
		},
		{
			Departure:     "AMS",
			Arrival:       "FCO",
			DepartureTime: "11:55",
			ArrivalTime:   "14:10",
			Price:         "104",
			Airline:       "Alitalia",
		},
		{
			Departure:     "AMS",
			Arrival:       "FCO",
			DepartureTime: "17:25",
			ArrivalTime:   "19:40",
			Price:         "122",
			Airline:       "Alitalia",
		},
		{
			Departure:     "AMS",
			Arrival:       "FCO",
			DepartureTime: "14:25",
			ArrivalTime:   "16:35",
			Price:         "135",
			Airline:       "KLM",
		},
		{
			Departure:     "AMS",
			Arrival:       "FCO",
			DepartureTime: "20:35",
			ArrivalTime:   "22:45",
			Price:         "135",
			Airline:       "KLM",
		},
		{
			Departure:     "AMS",
			Arrival:       "FCO",
			DepartureTime: "16:35",
			ArrivalTime:   "18:45",
			Price:         "135",
			Airline:       "KLM",
		},
		{
			Departure:     "AMS",
			Arrival:       "FCO",
			DepartureTime: "7:15",
			ArrivalTime:   "9:30",
			Price:         "135",
			Airline:       "KLM",
		},
		{
			Departure:     "AMS",
			Arrival:       "FCO",
			DepartureTime: "9:35",
			ArrivalTime:   "11:50",
			Price:         "135",
			Airline:       "KLM",
		},
	}
}
