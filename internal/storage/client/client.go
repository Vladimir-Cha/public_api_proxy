package client

import (
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"
)

type HTTPClient interface {
	Get(url string) ([]byte, error)
	Post(url string, body []byte) ([]byte, error)
}

type Client struct {
	client *resty.Client
}

func New(baseURL string, timeout time.Duration) *Client {
	return &Client{
		client: resty.New().
			SetBaseURL(baseURL).
			SetTimeout(timeout).
			SetHeader("Content-Type", "application/json"),
	}
}

func (c *Client) Get(endpoint string) ([]byte, error) {
	resp, err := c.client.R().
		Get(endpoint)

	if err != nil {
		return nil, fmt.Errorf("GET request failed: %w", err)
	}

	return resp.Body(), nil
}

// Запрос с телом
func (c *Client) Post(endpoint string, body []byte) ([]byte, error) {
	//Выполняем запрос
	resp, err := c.client.R().
		SetBody(body).
		Post(endpoint)

	if err != nil {
		return nil, fmt.Errorf("POST request failed: %w", err)
	}

	return resp.Body(), nil
}
