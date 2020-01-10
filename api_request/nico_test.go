package api_request

import (
	"testing"
)

func TestGetNicoResponse(t *testing.T) {
	resp := getNicoResponse("", "")
	if resp.StatusCode == 404 {
		t.Fatal("Error request")
	}
}
