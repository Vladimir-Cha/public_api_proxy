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
	httpConfig := loadConfig()
	if httpConfig == nil {
		log.Fatal("Failed to load HTTP-config")
	}

	//GET-запрос
	rawPost, err := httpConfig.Get("/posts/1")
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

	createdPost, err := httpConfig.Post("/posts", newPost)
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
	return client.New(cfg)
}
