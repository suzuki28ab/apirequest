package api_request

import (
	"testing"
)

func TestGetYoutubeLiveData(t *testing.T) {
	isLive, _ := GetYoutubeLiveData("UCoNJkwUekNXHvuGuX93rAhw")
	if isLive == true {
		t.Error(isLive)
	}
}
