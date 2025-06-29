package service

import (
	"fmt"

	"github.com/Vladimir-Cha/public_api_proxy/internal/storage/client"
)

type JSONPlaceholder struct {
	client client.HTTPClient
}

func New(httpClient client.HTTPClient) *JSONPlaceholder {
	return &JSONPlaceholder{
		client: httpClient,
	}
}

func (s *JSONPlaceholder) GetPost(id int) ([]byte, error) {
	endpoint := fmt.Sprintf("/posts/%d", id)
	return s.client.Get(endpoint)
}

func (s *JSONPlaceholder) Create(post []byte) ([]byte, error) {
	return s.client.Post("/posts", post)
}
