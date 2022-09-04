package slack

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

// refs:
// https://api.slack.com/messaging/webhooks

type body struct {
	Text string `json:"text"`
}

func Send(url, text string) error {
	body := body{
		Text: text,
	}
	bodyJson, err := json.Marshal(body)
	if err != nil {
		return err
	}
	bodyBytes := bytes.NewBuffer(bodyJson)

	req, err := http.NewRequest(http.MethodPost, url, bodyBytes)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	log.Println(string(respBytes))
	return nil
}
