package main

import (
	"encoding/json"
	"io/ioutil"
)

type Query struct {
	Departure     string
	Arrival       string
	DepartureTime string
	Email         string
}

// StoreJSONResults stores a list of results in a json file
func StoreJSONResults(results []FlightResult, filename string) {
	jsonString, err := json.MarshalIndent(results, "", "\t")
	LogErrorIfNotNull(err)
	err = ioutil.WriteFile(filename, jsonString, 0644)
	LogErrorIfNotNull(err)
}

// StoreJSONQueries stores a list of queries in a json file
func StoreJSONQueries(searchInput SearchInput, recipient string, filename string) {
	query := Query{
		Departure:     searchInput.Departure,
		Arrival:       searchInput.Arrival,
		DepartureTime: searchInput.Date,
		Email:         recipient,
	}
	jsonString, err := json.MarshalIndent(query, "", " ")
	LogErrorIfNotNull(err)
	err = ioutil.WriteFile(filename, jsonString, 0644)
	LogErrorIfNotNull(err)

}
