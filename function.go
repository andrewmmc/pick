package function

import (
	"context"
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"cloud.google.com/go/datastore"
)

type AuthResponse struct {
	AccessToken string `json:"access_token"`
	Scope       string `json:"scope"`
	TeamName    string `json:"team_name"`
	TeamID      string `json:"team_id"`
}

type Client struct {
	TeamName    string `datastore:"team_name,noindex"`
	AccessToken string `datastore:"access_token,noindex"`
	Scope       string `datastore:"scope,noindex"`
}

type Response struct {
	ResponseType string        `json:"response_type"`
	Text         string        `json:"text"`
	Attachments  []Attachments `json:"attachments"`
}

type Attachments struct {
	Text string `json:"text"`
}

var clientID = os.Getenv("SLACK_APP_CLIENT_ID")
var clientSecret = os.Getenv("SLACK_APP_CLIENT_SECRET")
var projectID = os.Getenv("PROJECT_ID")
var redirectURI = "https://" + os.Getenv("FUNCTION_REGION") + "-" + projectID + ".cloudfunctions.net/authCallback"

func Install(w http.ResponseWriter, r *http.Request) {
	scope := "commands"
	http.Redirect(w, r, "https://slack.com/oauth/authorize?client_id="+clientID+"&scope="+scope+"&redirect_uri="+redirectURI, 302)
}

func AuthCallback(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	code := q.Get("code")
	error := q.Get("error")

	if error == "access_denied" {
		log.Println("Error: access_denied")
		// redirect to error page
		return
	}

	form := url.Values{}
	form.Add("client_id", clientID)
	form.Add("client_secret", clientSecret)
	form.Add("code", code)
	form.Add("redirect_uri", redirectURI)

	req, err := http.NewRequest("POST", "https://slack.com/api/oauth.access", strings.NewReader(form.Encode()))
	req.PostForm = form
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	c := &http.Client{}
	res, err := c.Do(req)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	var payload AuthResponse
	err = decoder.Decode(&payload)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	createClient(payload.TeamID, payload.TeamName, payload.AccessToken, payload.Scope)
	// redirect to success page
	return
}

func GetAnswer(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	command := r.FormValue("command")
	userID := r.FormValue("user_id")
	text := r.FormValue("text")
	text = strings.TrimSpace(text)

	if command != "/pick" { // TODO: fix
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}

	message := ""
	var attachments []Attachments
	if text == "" {
		message = ":wave: Hi, I'm `pick`. Here are some tips to get you started!"
		attachments = append(attachments, Attachments{"*Basic*\nYou can ask me to randomly pick one choice for you from the list. For example:\n`/pick Chicken Pizza Kebab Pasta Rice`"})
	} else {
		// randomly pick one from the choices received
		choices := strings.Split(text, " ")
		rand.Seed(time.Now().Unix())
		message = "<@" + userID + ">, " + choices[rand.Intn(len(choices))] + "!"
	}

	body := Response{"in_channel", message, attachments}
	res, err := json.Marshal(body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func createClient(teamID string, teamName string, accessToken string, scope string) {
	kind := os.Getenv("DATA_STORE_KIND")
	ctx := context.Background()
	c, err := datastore.NewClient(ctx, projectID)

	key := datastore.NameKey(kind, teamID, nil)
	client := &Client{teamName, accessToken, scope}

	_, err = c.Put(ctx, key, client)
	if err != nil {
		log.Fatalf("Failed to save: %v", err)
	}
	return
}
