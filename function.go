package function

import (
	"encoding/json"
	"fmt"
	"html"
	"log"
	"net/http"
)

type Body struct {
	Challenge string `json:"challenge"`
	Event     Event
}

type Event struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

// respond with the challenge parameter value
func ReceiveEvent(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var payload Body
	err := decoder.Decode(&payload)
	if err != nil {
		fmt.Fprint(w, "Error!")
		return
	}

	if payload.Challenge != "" {
		fmt.Fprint(w, html.EscapeString(payload.Challenge))
		return
	}

	if payload.Event.Type == "app_mention" {
		log.Println(payload.Event.Text)
		return
	}
}
