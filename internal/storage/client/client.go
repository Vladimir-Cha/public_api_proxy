package client

import (
	"context"
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"
)

// структура для возврата метрик
type ResponseMetrics struct {
	Body       []byte
	StatusCode int
	Duration   time.Duration
}

type HTTPClient interface {
	Get(url string) ([]byte, error)
	Post(url string, body []byte) (*ResponseMetrics, error)
}

type Client struct {
	client *resty.Client
}

func New(baseURL string, timeout time.Duration) *Client {
	return &Client{
		client: resty.New().
			SetBaseURL(baseURL).
			SetTimeout(timeout).
			SetHeader("Content-Type", "application/json").
			OnBeforeRequest(func(c *resty.Client, req *resty.Request) error {
				req.SetContext(context.WithValue(req.Context(), "startTime", time.Now()))
				return nil
			}),
	}
}

func (c *Client) Get(endpoint string) (*ResponseMetrics, error) {
	startTime := time.Now()

	resp, err := c.client.R().
		Get(endpoint)

	if err != nil {
		return nil, fmt.Errorf("GET request failed: %w", err)
	}

	return &ResponseMetrics{
		Body:       resp.Body(),
		StatusCode: resp.StatusCode(),
		Duration:   time.Since(startTime),
	}, nil
}

// Запрос с телом
func (c *Client) Post(endpoint string, body []byte) (*ResponseMetrics, error) {
	startTime := time.Now()
	//Выполняем запрос
	resp, err := c.client.R().
		SetBody(body).
		Post(endpoint)

	if err != nil {
		return nil, fmt.Errorf("POST request failed: %w", err)
	}

	return &ResponseMetrics{
		Body:       resp.Body(),
		StatusCode: resp.StatusCode(),
		Duration:   time.Since(startTime),
	}, nil
}
