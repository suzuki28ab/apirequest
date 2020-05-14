package api_request

import (
	"testing"
)

func TestGetTwitchResponse(t *testing.T) {
	resp := getTwitchResponse("", "")
	if resp.StatusCode == 404 {
		t.Fatal("Error request")
	}
}
