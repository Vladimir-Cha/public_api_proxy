package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Vladimir-Cha/public_api_proxy/internal/storage/client"
	"github.com/Vladimir-Cha/public_api_proxy/internal/storage/config"
	"github.com/joho/godotenv"
)

func main() {
	cfg := loadConfig()
	if cfg == nil {
		log.Fatal("Failed to load HTTP-config")
	}
	httpClient := client.New(cfg)

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

func loadConfig() *config.Config {
	var cfg *config.Config

	//загрузка yaml конфига
	cfg, err := config.Load("config.yaml")
	if err != nil {
		log.Fatalf("error loading config file: %v", err)
		cfg = &config.Config{} // пустой конфиг
	}

	//загрузка env файла
	if err := godotenv.Load(); err == nil {
		if baseUrl := os.Getenv("API_BASE_URL"); baseUrl != "" {
			cfg.API.BaseURL = baseUrl
		}
		if logLevel := os.Getenv("LOG_LEVEL"); logLevel != "" {
			cfg.Logging.LevelLog = logLevel
		}
	}

	if cfg.API.BaseURL == "" {
		log.Fatalf("URL not found in config.yaml and .env")
	}

	log.Printf("Final config: %v", cfg)
	return cfg
}
