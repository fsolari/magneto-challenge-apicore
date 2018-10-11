package handler

import (
	"testing"
	"net/http/httptest"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"strings"
)

func TestPingMustReturnPong(t *testing.T) {

	url := "/ping"
	r := handlerTest("GET", url, "")

	assert.Equal(t, 200, r.Code)
	assert.Equal(t, "pong", r.Body.String())

}

func handlerTest(method string, url string, body string) *httptest.ResponseRecorder {
	router := gin.Default()
	MapUrls(router)

	var req *http.Request

	if body != "" {
		req, _ = http.NewRequest(method, url, strings.NewReader(body))
	} else {
		req, _ = http.NewRequest(method, url, nil)
	}

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	return rr
}