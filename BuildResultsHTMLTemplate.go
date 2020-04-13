package main

import (
	"bytes"
	"html/template"
)

type TemplateData struct {
	Results     []FlightResult
	SearchInput SearchInput
}

var layout = `<html>
  <body>
    Flights from {{.SearchInput.Departure}} to {{.SearchInput.Arrival}} on {{.SearchInput.Date}}
    <table border='1' cellpadding='3px' style='border-collapse: collapse;'>
        <tr>
          <th>Route</th><th>Time</th><th>Airline</th><th>Price</th>
        </tr>
        {{range .Results}}
        <tr>
          <td>{{.Departure}} - {{.Arrival}}</td>
          <td>{{.DepartureTime}} - {{.ArrivalTime}}</td>
          <td>{{.Airline}}</td>
          <td>{{.Price}}</td>
        </tr>
        {{end}}
    </table>
  </body>
</html>`

// BuildResultsHTMLTemplate build an HTML template for results and return it as string
func BuildResultsHTMLTemplate(results []FlightResult, searchInput SearchInput) []byte {
	data := TemplateData{
		Results:     results,
		SearchInput: searchInput,
	}
	var tmplBuffer bytes.Buffer
	tmpl, err := template.New("results").Parse(layout)
	LogFatalAndExitIfNotNull(err)
	err = tmpl.Execute(&tmplBuffer, data)
	LogFatalAndExitIfNotNull(err)

	return tmplBuffer.Bytes()
}
