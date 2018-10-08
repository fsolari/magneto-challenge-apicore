package test

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"github.com/mercadolibre/magneto-challenge-apicore/app"
	"strings"
	"log"
)

func TestPingMustReturnPong(t *testing.T) {

	url := "/ping"
	r := executeRequest("GET", url, "")

	assert.Equal(t, 200, r.Code)
	assert.Equal(t, "pong", r.Body.String())

}

func executeRequest(method string, url string, body string) *httptest.ResponseRecorder {
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
		log.Printf("Error executing http request with body to %s" + url, err)
	}

	Router.ServeHTTP(response, request)

	return response
}
