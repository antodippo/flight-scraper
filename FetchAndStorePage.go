package main

import (
	"github.com/antchfx/htmlquery"
	"go.zoe.im/surferua"
	"golang.org/x/net/html"
	"io/ioutil"
	"net/http"
)

// FetchAndStorePage fetches a web page and stores its content
func FetchAndStorePage(url string, filename string) (doc *html.Node) {
	request, _ := http.NewRequest(http.MethodGet, url, nil)
	request.Header.Add("Cache-Control", "no-cache")
	request.Header.Add("Accept", "*/*")
	request.Header.Add("Connection", "keep-alive")
	request.Header.Add("User-Agent", surferua.New().Desktop().Chrome().String())

	response, err := http.DefaultClient.Do(request)
	LogFatalAndExitIfNotNull(err)

	body, err := ioutil.ReadAll(response.Body)
	LogFatalAndExitIfNotNull(err)

	ioutil.WriteFile(filename, []byte(body), 0644)
	doc, err = htmlquery.LoadDoc(filename)
	LogFatalAndExitIfNotNull(err)

	return
}
