package function

import (
	"math/rand"
	"net/http"
	"net/http/httptest"
	"testing"
)

func init() {
	rand.Seed(1)
}

func TestInstall(t *testing.T) {
	req, err := http.NewRequest("GET", "/install", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Install)

	handler.ServeHTTP(rr, req)
	status := rr.Code
	if status != http.StatusFound {
		t.Errorf("Expected statusCode: %v, but got %v", http.StatusFound, status)
	}
}

func TestRollChoice(t *testing.T) {
	choices := []string{"choice1", "choice2", "choice3"}
	answer := rollChoice(choices)
	if answer != "choice3" {
		t.Errorf("Expected rollChoice: %s, but got: %s", "choice3", answer)
	}
}
