package main

import (
	"flag"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"time"
)

// ParseFlags parses input flags
func ParseFlags(progname string, args []string) (searchInput SearchInput, recipient string) {
	flags := flag.NewFlagSet(progname, flag.ExitOnError)
	departurePtr := flags.String("departure", "", "Airport code (ex. FCO, AMS, RTM, JFK...)")
	arrivalPtr := flags.String("arrival", "", "Airport code (ex. FCO, AMS, RTM, JFK...)")
	datePtr := flags.String("date", "", "Date of the flight, YYYY-MM-DD format")
	recipientPtr := flags.String("recipient", "", "Email address to notify")
	flags.Parse(args)

	if *departurePtr == "" || *arrivalPtr == "" || *datePtr == "" || *recipientPtr == "" {
		handleWarning(flags, "Not enough parameters:")
	}
	if !isIATAValid(departurePtr) || !isIATAValid(arrivalPtr) {
		handleWarning(flags, "Invalid IATA:")
	}
	if !isDateValid(datePtr) {
		handleWarning(flags, "Invalid date format:")
	}
	if !IsEmailValid(recipientPtr) {
		handleWarning(flags, "Invalid email format:")
	}

	return SearchInput{*departurePtr, *arrivalPtr, *datePtr}, *recipientPtr
}

func handleWarning(flags *flag.FlagSet, message string) {
	LogWarning(message)
	flags.PrintDefaults()
	os.Exit(0)
}

func isDateValid(datePtr *string) bool {
	_, err := time.Parse("2006-01-02", *datePtr)
	if err != nil {
		return false
	} else {
		return true
	}
}

func isIATAValid(departurePtr *string) bool {
	//api_key has only 1000 requesst available per month.
	var api_key string = "b6a88e00-390d-4187-930f-550ea564f9b3"
	url := "https://airlabs.co/api/v9/airports.csv?iata_code=" + *departurePtr + "&api_key=" + api_key
	req, _ := http.NewRequest("GET", url, nil)
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	if string(body) == "" {
		return false
	} else {
		return true
	}
}

func IsEmailValid(email *string) bool {
	emailPattern := regexp.MustCompile(`^[A-Za-z0-9]+[@]+[A-Za-z0-9]+[.]+com$`)
	if !emailPattern.MatchString(*email) || len(*email) > 50 {
		return false
	}
	return true
}
