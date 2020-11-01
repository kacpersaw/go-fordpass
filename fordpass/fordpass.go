package fordpass

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	baseURL = "https://usapi.cv.ford.com/api"
	authURL = "https://fcis.ice.ibmcloud.com/v1.0/endpoint/default/token"
)

type Client struct {
	BaseURL string

	Username string
	Password string
	Vin      string

	Token     string
	ExpiresAt time.Time

	client *http.Client
}

func NewClient(username, password, vin string) *Client {
	return &Client{
		BaseURL: baseURL,

		Username: username,
		Password: password,
		Vin:      vin,

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

type ControlResponse struct {
	CommandId string `json:"commandId"`
	Status    int    `json:"status"`
}

// Authenticate to ford api
func (c *Client) auth() error {
	data := url.Values{}
	data.Set("client_id", "9fb503e0-715b-47e8-adfd-ad4b7770f73b")
	data.Set("grant_type", "password")
	data.Set("username", c.Username)
	data.Set("password", c.Password)

	req, err := http.NewRequest("POST", authURL, strings.NewReader(data.Encode()))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	var res AuthResponse
	if err := c.doRequest(req, &res); err != nil {
		return err
	}

	c.Token = res.AccessToken

	t := time.Now()
	t.Add(time.Duration(res.ExpiresIn) * time.Second)
	c.ExpiresAt = t

	return nil
}

// Fetch and refresh token if needed
func (c *Client) acquireToken() error {
	if c.Token == "" || time.Now().After(c.ExpiresAt) {
		return c.auth()
	}

	return nil
}

// Do request to the given URL
func (c *Client) doRequest(req *http.Request, v interface{}) error {
	c.setHeaders(req)

	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusBadRequest {
		var errRes map[string]string
		if err = json.NewDecoder(resp.Body).Decode(&errRes); err == nil {
			return errors.New(errRes["error_description"])
		}
	}

	if err = json.NewDecoder(resp.Body).Decode(&v); err != nil {
		return err
	}

	return nil
}

// Poll status of pending command
func (c *Client) pollStatus(url, id string) error {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s", url, id), nil)
	if err != nil {
		return err
	}
	c.setHeaders(req)

	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var data ControlResponse
	if err = json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return err
	}

	switch data.Status {
	case 552:
		time.Sleep(5 * time.Second)
		return c.pollStatus(url, id)
	case 200:
		return nil
	default:
		return errors.New("poll failed")
	}
}

// Set needed headers for request
func (c *Client) setHeaders(req *http.Request) {
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Language", "en-us")
	req.Header.Set("User-Agent", "fordpass-na/353 CFNetwork/1121.2.2 Darwin/19.3.0")
	req.Header.Set("Accept-Encoding", "gzip, deflate, br")

	if req.URL.String() != authURL {
		req.Header.Set("Application-Id", "71A3AD0A-CF46-4CCF-B473-FC7FE5BC4592")
		req.Header.Set("Content-Type", "application/json")
	}

	//set auth token header if set
	if c.Token != "" {
		req.Header.Set("auth-token", c.Token)
	}
}
