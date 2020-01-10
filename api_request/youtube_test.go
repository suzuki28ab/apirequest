package api_request

import (
	"testing"
)
func TestGetYoutubeResponse(t *testing.T) {
	resp := getYoutubeResponse("")
	if resp.StatusCode == 404 {
		t.Fatal("Error request")
	}
}
