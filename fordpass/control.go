package fordpass

import (
	"errors"
	"fmt"
	"net/http"
)

func (c *Client) Lock() error {
	return c.issueCommand("lock")
}

func (c *Client) Unlock() error {
	return c.issueCommand("unlock")
}

func (c *Client) Start() error {
	return c.issueCommand("start")
}

func (c *Client) Stop() error {
	return c.issueCommand("stop")
}

func (c *Client) issueCommand(command string) error {
	if err := c.acquireToken(); err != nil {
		return err
	}

	var method string
	var url string

	switch command {
	case "lock":
		method = "PUT"
		url = fmt.Sprintf("%s/vehicles/v2/%s/doors/lock", baseURL, c.Vin)
	case "unlock":
		method = "DELETE"
		url = fmt.Sprintf("%s/vehicles/v2/%s/doors/lock", baseURL, c.Vin)
	case "start":
		method = "PUT"
		url = fmt.Sprintf("%s/vehicles/v2/%s/engine/start", baseURL, c.Vin)
	case "stop":
		method = "DELETE"
		url = fmt.Sprintf("%s/vehicles/v2/%s/engine/start", baseURL, c.Vin)
	default:
		return errors.New("no command specified")
	}

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return err
	}

	var res ControlResponse
	if err := c.doRequest(req, &res); err != nil {
		return err
	}

	return c.pollStatus(url, res.CommandId)
}
