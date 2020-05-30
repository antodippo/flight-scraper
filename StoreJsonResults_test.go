package main

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStoreJsonResults(t *testing.T) {
	t.Run("It stores a list of results into a json file", func(t *testing.T) {
		list := getInputResultList()
		filename := "data-test/test_store_json"
		StoreJSONResults(list, filename)

		want := getWantedJSON()
		got, _ := ioutil.ReadFile(filename)
		assert.Equal(t, want, string(got))
	})
}

func getInputResultList() []FlightResult {
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
			Price:         "123",
			Airline:       "KLM",
		},
	}
}

func getWantedJSON() string {
	return `[
	{
		"Departure": "AMS",
		"Arrival": "FCO",
		"DepartureTime": "16:30",
		"ArrivalTime": "18:50",
		"Price": "69",
		"Airline": "LEVEL"
	},
	{
		"Departure": "AMS",
		"Arrival": "FCO",
		"DepartureTime": "7:20",
		"ArrivalTime": "9:45",
		"Price": "123",
		"Airline": "KLM"
	}
]`
}
