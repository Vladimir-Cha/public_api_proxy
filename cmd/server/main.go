package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/Vladimir-Cha/public_api_proxy/internal/storage/client"
	"github.com/Vladimir-Cha/public_api_proxy/internal/storage/config"
	"github.com/joho/godotenv"
)

func main() {
	httpClient := loadConfig()
	if httpClient == nil {
		log.Fatal("Failed to load HTTP-config")
	}

	//GET-запрос
	rawPost, err := httpClient.Get("/posts/1")
	if err != nil {
		log.Fatalf("Error get post: %v", err)
	}
	log.Printf(
		"Get post:\nAnswer: %s\nResponse code: %d\nResponse time: %v\n",
		string(rawPost.Body),
		rawPost.StatusCode,
		rawPost.Duration,
	)

	//POST-запрос
	newPost := []byte(`{
        "title": "Новый пост",
        "body": "Содержание поста",
        "userId": 1
    }`)

	createdPost, err := httpClient.Post("/posts", newPost)
	if err != nil {
		log.Fatalf("Error create post: %v", err)
	}
	fmt.Printf("Created post:\nResponse code: %d\nResponse time: %v\n%s",
		createdPost.StatusCode,
		createdPost.Duration,
		string(createdPost.Body),
	)
}

func loadConfig() *client.Client {
	var cfg *config.Config

	//загрузка env файла
	if err := godotenv.Load(); err == nil {
		timeout, _ := strconv.Atoi(os.Getenv("API_TIMEOUT_SECONDS"))
		baseUrl := os.Getenv("API_BASE_URL")
		if baseUrl == "" {
			log.Fatalf("URL not found")
		}

		cfg = &config.Config{
			API: config.APIconfig{
				BaseURL: baseUrl,
				Timeout: time.Duration(timeout) * time.Second,
			},
			Logging: config.LoggingConfig{
				Enabled:  os.Getenv("LOG_ENABLED") == "true",
				LevelLog: os.Getenv("LOG_LEVEL"),
			},
		}
	} else {
		//загрузка yaml конфига, если нет файла .env
		cfg, err = config.Load("config.yaml")
		if err != nil {
			log.Fatalf("error loading config file: %v", err)
		}
	}
	return client.New(cfg)
}
