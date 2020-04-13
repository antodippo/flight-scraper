package main

import (
	"encoding/json"
	"io/ioutil"
)

// StoreJSONResults stores a list of results in a json file
func StoreJSONResults(results []FlightResult, filename string) {
	jsonString, err := json.MarshalIndent(results, "", "\t")
	LogErrorIfNotNull(err)
	err = ioutil.WriteFile(filename, jsonString, 0644)
	LogErrorIfNotNull(err)
}
