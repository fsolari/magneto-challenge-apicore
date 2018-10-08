package test

import (
	"testing"
	"fmt"
)

func TestGetStatsMustReturnStats(t *testing.T) {

	url := "/stats"
	r := executeRequest("GET", url, "")

	fmt.Printf(r.Body.String())

}
