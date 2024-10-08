package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Client struct {
	token  string
	client http.Client
}

func NewClient(token string) (*Client, error) {
	if token == "" {
		return nil, fmt.Errorf("No token specified, please provide token")
	}

	return &Client{
		token:  token,
		client: http.Client{Timeout: 5 * time.Second},
	}, nil
}

func (cl *Client) ContributorsList(repo string) ([]Contributor, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/contributors", repo)
	req, err := http.NewRequest("Get", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+cl.token)

	// Execute the request
	res, err := cl.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	// Ensure status code is 200
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API responded with a %d %s", res.StatusCode, res.Status)
	}

	// Decode results
	var cons []Contributor
	if err := json.NewDecoder(res.Body).Decode(&cons); err != nil {
		return nil, err
	}

	return cons, nil
}
