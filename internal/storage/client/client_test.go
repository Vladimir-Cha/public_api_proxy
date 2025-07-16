package client

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/Vladimir-Cha/public_api_proxy/internal/storage/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// тестовый сервер
func setupTestServer(status int, body string) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(status)
		w.Write([]byte(body))

	}))
}

// тест GET запроса
func TestGet(t *testing.T) {
	server := setupTestServer(http.StatusOK, `{"id": 1}`)
	defer server.Close()

	// конфиг клиента
	cfg := &config.Config{
		API: config.APIconfig{
			BaseURL: server.URL,
			Timeout: 1 * time.Second,
		},
	}
	client := New(cfg) // создан клиент
	resp, err := client.Get("/")

	require.NoError(t, err)                         // отсутствие ошибок
	assert.Equal(t, http.StatusOK, resp.StatusCode) // ответ должен быть 200 OK
	assert.Equal(t, `{"id": 1}`, string(resp.Body)) // сравнение тела ответа
}

func TestGetNotFound(t *testing.T) {
	server := setupTestServer(http.StatusNotFound, `Not Found`)
	defer server.Close()

	// конфиг клиента
	cfg := &config.Config{
		API: config.APIconfig{
			BaseURL: server.URL,
			Timeout: 1 * time.Second,
		},
	}
	client := New(cfg)                 // создан клиент
	resp, _ := client.Get("/notfound") // запрос к несуществующему ресурсу

	assert.Equal(t, http.StatusNotFound, resp.StatusCode) // есть ошибка
	assert.Contains(t, string(resp.Body), "Not Found")    // ответ должен быть 404
}

func TestPost(t *testing.T) {
	server := setupTestServer(http.StatusCreated, `{"id": 1}`)
	defer server.Close()

	// конфиг клиента
	cfg := &config.Config{
		API: config.APIconfig{
			BaseURL: server.URL,
			Timeout: 1 * time.Second,
		},
	}
	client := New(cfg) // создан клиент
	data := []byte(`{"title":"Test"}`)
	resp, err := client.Post("/posts", data)

	require.NoError(t, err)                              // отсутствие ошибок
	assert.Equal(t, http.StatusCreated, resp.StatusCode) // ответ должен быть 201
	assert.Contains(t, string(resp.Body), `{"id": 1}`)   // сравнение тела ответа
}
