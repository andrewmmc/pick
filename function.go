package function

import (
	"encoding/json"
	"fmt"
	"html"
	"net/http"
)

type body struct {
	Challenge string `json:"challenge"`
}

// respond with the challenge parameter value
func ReceiveEvent(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var d body
	err := decoder.Decode(&d)
	if err != nil {
		fmt.Fprint(w, "Error!")
		return
	}
	fmt.Fprint(w, html.EscapeString(d.Challenge))
}
