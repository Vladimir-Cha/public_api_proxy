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

	switch {
	case resp.StatusCode >= http.StatusOK && resp.StatusCode < http.StatusMultipleChoices:
		return body, nil
	case resp.StatusCode >= http.StatusBadRequest && resp.StatusCode < http.StatusInternalServerError:
		return nil, fmt.Errorf("client error %d: %s", resp.StatusCode, http.StatusText(resp.StatusCode))
	case resp.StatusCode >= http.StatusInternalServerError:
		return nil, fmt.Errorf("server error %d: %s", resp.StatusCode, http.StatusText(resp.StatusCode))
	default:
		return nil, fmt.Errorf("unexpected status code %d: %s", resp.StatusCode, http.StatusText(resp.StatusCode))
	}
}
