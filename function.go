package function

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html"
	"log"
	"net/http"
	"strings"
)

type Body struct {
	Challenge   string   `json:"challenge"`
	Token       string   `json:"token"`
	TeamID      string   `json:"team_id"`
	APIAppID    string   `json:"api_app_id"`
	AuthedUsers []string `json:"authed_users"`
	Event       Event
}

type Event struct {
	Channel string `json:"channel"`
	Type    string `json:"type"`
	Text    string `json:"text"`
	User    string `json:"user"`
}

type Message struct {
	Channel string `json:"channel"`
	Text    string `json:"text"`
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
		log.Printf("%v", payload.AuthedUsers)
		text := payload.Event.Text
		// remove bot name from received text
		for _, user := range payload.AuthedUsers {
			bot := strings.Join([]string{"<@", user, ">"}, "")
			text = strings.Replace(text, bot, "", -1)
		}
		log.Println(text)
		SendMessage(payload.Event.Channel, text)
		return
	}
}

func SendMessage(channel string, text string) {
	m := Message{channel, text}
	b, _ := json.Marshal(m)
	req, _ := http.NewRequest("POST", "https://slack.com/api/chat.postMessage", bytes.NewBuffer(b))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer") // token here
	c := &http.Client{}
	_, err := c.Do(req)
	if err != nil {
		log.Println("Error")
		return
	}
	return
}
