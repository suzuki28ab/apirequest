package api_request

import (
	"testing"
)

func TestGetMixerResponse(t *testing.T) {
	resp := getMixerResponse("test")
	if resp.StatusCode == 404 {
		t.Fatal("Error request")
	}
}
