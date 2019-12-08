package soundcloud

import (
	"github.com/go-resty/resty/v2"
)

const base = "https://api-v2.soundcloud.com"

type Client struct {
	c        *resty.Client
	clientID string
}

func NewClient(clientID string) *Client {
	return &Client{
		c:        resty.New().SetHostURL(base).SetQueryParam("client_id", clientID),
		clientID: clientID,
	}
}
