package test

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestHomeMustReturn200(t *testing.T) {

	url := "/"
	r := ExecuteRequest("GET", url, "")

	assert.Equal(t, 200, r.Code)
}
