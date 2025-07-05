package client

import (
	"fmt"
	"time"

	"github.com/Vladimir-Cha/public_api_proxy/internal/storage/config"
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
	config *config.Config
}

func New(cfg *config.Config) *Client {

	c := resty.New().
		SetBaseURL(cfg.API.BaseURL).
		SetTimeout(cfg.API.Timeout)

	if cfg.Logging.Enabled {
		c.SetDebug(true)
	}

	return &Client{
		client: c,
		//config: cfg,
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
