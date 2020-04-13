package main

import (
	"bytes"
	"html/template"
)

type TemplateData struct {
	Results []FlightResult
}

// BuildResultsHTMLTemplate build an HTML template for results and return it as string
func BuildResultsHTMLTemplate(results []FlightResult) []byte {
	data := TemplateData{
		Results: results,
	}
	var tmplBuffer bytes.Buffer
	tmpl, err := template.ParseFiles("layout.html")
	LogFatalAndExitIfNotNull(err)
	tmpl.Execute(&tmplBuffer, data)

	return tmplBuffer.Bytes()
}
