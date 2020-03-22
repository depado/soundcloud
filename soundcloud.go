package soundcloud

import (
	"github.com/go-resty/resty/v2"
)

const base = "https://api-v2.soundcloud.com"

type Client struct {
	c        *resty.Client
	clientID string
}

// NewClient will create a new SoundCloud client with the specified client ID
func NewClient(clientID string) *Client {
	return &Client{
		c:        resty.New().SetHostURL(base).SetQueryParam("client_id", clientID),
		clientID: clientID,
	}
}

// NewAutoIDClient will fetch a client ID from soundcloud's homepage and create
// a new client using this client ID
func NewAutoIDClient() (*Client, error) {
	t, err := NewClientIDFromPublicHTML()
	if err != nil {
		return nil, err
	}
	return &Client{
		c:        resty.New().SetHostURL(base).SetQueryParam("client_id", t),
		clientID: t,
	}, nil
}
