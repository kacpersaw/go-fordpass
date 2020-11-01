package fordpass

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	defaultBaseURL = "https://usapi.cv.ford.com/api"
	defaultAuthURL = "https://fcis.ice.ibmcloud.com/v1.0/endpoint/default/token"
)

type Client struct {
	BaseURL string

	username string
	password string
	vin      string

	token     string
	expiresAt time.Time

	client *http.Client
}

func NewClient(username, password, vin string) *Client {
	return &Client{
		BaseURL: defaultBaseURL,

		username: username,
		password: password,
		vin:      vin,

		client: &http.Client{
			Timeout: time.Minute,
		},
	}
}

type AuthResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	GrantId      string `json:"grant_id"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
}

func (c *Client) Authenticate() error {
	data := url.Values{}
	data.Set("client_id", "9fb503e0-715b-47e8-adfd-ad4b7770f73b")
	data.Set("grant_type", "password")
	data.Set("username", c.username)
	data.Set("password", c.password)

	req, err := http.NewRequest("POST", defaultAuthURL, strings.NewReader(data.Encode()))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	var res AuthResponse
	if err := c.doRequest(req, &res); err != nil {
		return err
	}

	c.token = res.AccessToken

	t := time.Now()
	t.Add(time.Duration(res.ExpiresIn) * time.Second)
	c.expiresAt = t

	return nil
}

func (c *Client) doRequest(req *http.Request, v interface{}) error {
	setHeaders(req)
	res, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		var errRes map[string]string
		if err = json.NewDecoder(res.Body).Decode(&errRes); err == nil {
			return errors.New(errRes["error_description"])
		}
	}

	if err = json.NewDecoder(res.Body).Decode(&v); err != nil {
		return err
	}
	return nil
}

func setHeaders(req *http.Request) {
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Language", "en-us")
	req.Header.Set("User-Agent", "fordpass-na/353 CFNetwork/1121.2.2 Darwin/19.3.0")
	req.Header.Set("Accept-Encoding", "gzip, deflate, br")

	if req.URL.String() != defaultAuthURL {
		req.Header.Set("Application-Id", "71A3AD0A-CF46-4CCF-B473-FC7FE5BC4592")
		req.Header.Set("Content-Type", "application/json")
	}
}
