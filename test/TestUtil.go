package test

import (
	"net/http/httptest"
	"strings"
	"log"
	"net/http"
	"github.com/mercadolibre/magneto-challenge-apicore/app"
)

func ExecuteRequest(method string, url string, body string) *httptest.ResponseRecorder {
	var request *http.Request
	var err error

	Router := app.SetupRouter()

	response := httptest.NewRecorder()

	if body != "" {
		request, err = http.NewRequest(method, url, strings.NewReader(body))
	} else {
		request, err = http.NewRequest(method, url, nil)
	}

	if err != nil {
		log.Printf("[TestUtil.ExecuteRequest] Error executing http request with body to : %s " + url, err)
	}

	Router.ServeHTTP(response, request)

	return response
}
