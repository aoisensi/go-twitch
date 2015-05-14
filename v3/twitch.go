package twitch

import (
	"encoding/json"
	"net/http"
	"time"
)

const root = "https://api.twitch.tv/kraken"

type Client struct {
	client   *http.Client
	token    token
	Channels *ChannelsService
}

type token struct {
	Authorization struct {
		Scopes    []string  `json:"scopes"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	} `json:"authorization"`
	Username string `json:"user_name"`
	Valid    bool   `json:"valid"`
}

func NewClient(httpClient *http.Client) (*Client, error) {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	resp, err := httpClient.Get(root)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var body struct {
		Token token `json:"token"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		return nil, err
	}
	c := &Client{
		client: httpClient,
		token:  body.Token,
	}
	c.Channels = &ChannelsService{client: c}

	return c, nil
}

func (c *Client) httpGet(url string, body interface{}) error {
	resp, err := c.client.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	d := json.NewDecoder(resp.Body)
	if resp.StatusCode != http.StatusOK {
		var err Error
		d.Decode(&err)
		return err
	}
	d.Decode(body)
	return nil
}
