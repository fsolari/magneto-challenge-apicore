package test

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestPingMustReturnPong(t *testing.T) {

	url := "/ping"
	r := ExecuteRequest("GET", url, "")

	assert.Equal(t, 200, r.Code)
	assert.Equal(t, "pong", r.Body.String())

}


