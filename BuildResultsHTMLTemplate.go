package main

import (
	"bytes"
	"html/template"
)

type TemplateData struct {
	Results     []FlightResult
	SearchInput SearchInput
}

// BuildResultsHTMLTemplate build an HTML template for results and return it as string
func BuildResultsHTMLTemplate(results []FlightResult, searchInput SearchInput) []byte {
	data := TemplateData{
		Results:     results,
		SearchInput: searchInput,
	}
	var tmplBuffer bytes.Buffer
	tmpl, err := template.ParseFiles("assets/layout.html")
	LogFatalAndExitIfNotNull(err)
	tmpl.Execute(&tmplBuffer, data)

	return tmplBuffer.Bytes()
}
