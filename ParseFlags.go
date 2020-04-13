package main

import (
	"flag"
	"os"
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
		LogWarning("Not enough parameters:")
		flags.PrintDefaults()
		os.Exit(0)
	}

	return SearchInput{*departurePtr, *arrivalPtr, *datePtr}, *recipientPtr
}
