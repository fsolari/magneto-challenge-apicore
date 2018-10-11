package handler

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestHomeMustReturn200(t *testing.T) {

	url := "/"
	r := handlerTest("GET", url, "")

	assert.Equal(t, 200, r.Code)
}
