package pulse

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

var (
	urlClassify = "https://api.pulsesecurity.org/api/classify"
)

type Client struct {
	siteKey    string
	siteSecret string
	http       *http.Client
}

func New(siteKey, siteSecret string) *Client {
	return &Client{
		siteKey:    siteKey,
		siteSecret: siteSecret,
		http:       http.DefaultClient,
	}
}

func (c *Client) Classify(token string) (bool, error) {
	payload, err := json.Marshal(&classifyPayload{
		SiteKey:   c.siteKey,
		SecretKey: c.siteSecret,
		Token:     token,
	})
	if err != nil {
		return false, err
	}

	req, err := http.NewRequest(http.MethodPost, urlClassify, bytes.NewReader(payload))
	if err != nil {
		return false, err
	}

	req.Header.Set("Content-Type", "application/json")

	res, err := c.http.Do(req)
	if err != nil {
		return false, err
	}

	defer res.Body.Close()

	var response classifyResponse
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return false, err
	}

	if err := response.Error(); err != nil {
		return false, err
	}

	if response.IsBot == nil {
		return false, fmt.Errorf("missing isBot field in response")
	}

	return *response.IsBot, nil
}

type classifyPayload struct {
	SiteKey   string `json:"siteKey"`
	SecretKey string `json:"secretKey"`
	Token     string `json:"token"`
}

type classifyResponse struct {
	errorResponse
	IsBot *bool `json:"isBot,omitempty"`
}
