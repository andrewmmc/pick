package function

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html"
	"log"
	"net/http"
)

type Body struct {
	Challenge string `json:"challenge"`
	Token     string `json:"token"`
	TeamID    string `json:"team_id"`
	APIAppID  string `json:"api_app_id"`
	Event     Event
}

type Event struct {
	Channel string `json:"channel"`
	Type    string `json:"type"`
	Text    string `json:"text"`
	User    string `json:"user"`
}

type Message struct {
	Text    string `json:"text"`
	Channel string `json:"channel"`
}

func ReceiveEvent(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var payload Body
	err := decoder.Decode(&payload)
	if err != nil {
		fmt.Fprint(w, "Error!")
		return
	}

	if payload.Challenge != "" {
		// respond with the challenge parameter value
		fmt.Fprint(w, html.EscapeString(payload.Challenge))
		return
	}

	if payload.Event.Type == "app_mention" {
		log.Println(payload.Token)
		log.Println(payload.Event.Text)
		log.Println(payload.Event.User)
		log.Println(payload.Event.Channel)
		SendMessage()
		return
	}
}

func SendMessage() {
	client := &http.Client{}
	m := Message{"testing", ""}
	b, _ := json.Marshal(m)
	req, err := http.NewRequest("POST", "https://slack.com/api/chat.postMessage", bytes.NewBuffer(b))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer ")
	if err != nil {
		log.Println("Error")
		return
	}
	_, err = client.Do(req)
	if err != nil {
		log.Println("Error")
		return
	}
}
