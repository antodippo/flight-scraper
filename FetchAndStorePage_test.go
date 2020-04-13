package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/antchfx/htmlquery"
)

func TestFetchAndStorePage(t *testing.T) {
	filepath := "data_test/test_store"
	url := serverMock().URL + "/test-flights"

	t.Run("It fetches an html page and stores it on filesystem", func(t *testing.T) {
		doc := FetchAndStorePage(url, filepath)
		got := htmlquery.OutputHTML(doc, true)
		want := "<html><head></head><body><div>test</div></body></html>"
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}

func serverMock() *httptest.Server {
	handler := http.NewServeMux()
	handler.HandleFunc("/test-flights", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("<html><head></head><body><div>test</div></body></html>"))
	})
	server := httptest.NewServer(handler)

	return server
}
