package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type parseArgumentsTestCase struct {
	args        []string
	searchInput SearchInput
	recipient   string
}

func TestParseArguments(t *testing.T) {
	testCases := getTestCases()
	for _, testCase := range testCases {
		t.Run("It parses the argumentents provided", func(t *testing.T) {
			searchInput, recipient := ParseFlags("test", testCase.args)
			assert.Equal(t, testCase.searchInput, searchInput)
			assert.Equal(t, testCase.recipient, recipient)
		})
	}
}

func getTestCases() []parseArgumentsTestCase {
	return []parseArgumentsTestCase{
		{
			[]string{"-departure=FCO", "-arrival=AMS", "-date=2020-02-02", "-recipient=test@example.com"},
			SearchInput{"FCO", "AMS", "2020-02-02"},
			"test@example.com",
		},
	}
}
