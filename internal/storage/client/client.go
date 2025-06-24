package client

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"time"
)

type HTTPClient interface {
	Get(url string) ([]byte, error)
	Post(url string, body []byte) ([]byte, error)
}

type Client struct {
	baseURL string
	client  *http.Client
}

func New(baseURL string, timeout time.Duration) *Client {
	return &Client{
		baseURL: baseURL,
		client: &http.Client{
			Timeout: timeout,
		},
	}
}

func (c *Client) Get(endpoint string) ([]byte, error) {
	fullURL := c.baseURL + endpoint

	resp, err := c.client.Get(fullURL)

	if err != nil {
		return nil, fmt.Errorf("GET request failed: %w", err)
	}
	defer resp.Body.Close()

	return handleResponse(resp)
}

// Запрос с телом
func (c *Client) Post(endpoint string, body []byte) ([]byte, error) {
	fullURL := c.baseURL + endpoint

	//Выполняем запрос
	resp, err := http.Post(fullURL, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("POST request failed: %w", err)
	}
	defer resp.Body.Close()

	return handleResponse(resp)
}

func handleResponse(resp *http.Response) ([]byte, error) {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode < 200 && resp.StatusCode >= 300 && resp.StatusCode < 400 {
		return nil, fmt.Errorf("error: %d", resp.StatusCode)
	} else if resp.StatusCode >= 400 && resp.StatusCode < 500 {
		return nil, fmt.Errorf("client error: %d", resp.StatusCode)
	} else if resp.StatusCode >= 500 {
		return nil, fmt.Errorf("server error: %d", resp.StatusCode)
	}

	return body, nil
}
