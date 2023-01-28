package turnstile

import (
	"encoding/json"
	"net/http"
	"net/url"
	"time"
)

const (
	API_ENDPOINT    = "https://challenges.cloudflare.com/turnstile/v0/siteverify"
	DEFAULT_TIMEOUT = 15
)

type Turnstile struct {
	SecretKey string
	Timeout   time.Duration
}

func New(secretKey string, timeout int) *Turnstile {
	return &Turnstile{
		SecretKey: secretKey,
		Timeout:   time.Duration(timeout) * time.Second,
	}
}

func (t *Turnstile) Verify(responseToken string, remoteIP string) (*Response, error) {
	data := url.Values{
		"secret":   {t.SecretKey},
		"response": {responseToken},
	}
	if remoteIP == "" {
		data.Add("remoteip", remoteIP)
	}

	request, err := http.PostForm(API_ENDPOINT, data)
	if err != nil {
		return nil, err
	}
	defer request.Body.Close()

	response := &Response{}
	if err := json.NewDecoder(request.Body).Decode(response); err != nil {
		return nil, err
	}

	return response, nil
}
