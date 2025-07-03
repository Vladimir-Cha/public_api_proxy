package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/Vladimir-Cha/public_api_proxy/internal/storage/client"
	"github.com/joho/godotenv"
)

func main() {
	//загрузка env файла
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("error loading .env file")
	}

	timeout, _ := strconv.Atoi(os.Getenv("API_TIMEOUT"))
	baseUrl := os.Getenv("API_BASE_URL")
	if baseUrl == "" {
		log.Printf("URL not found")
	}
	//Инициализация HTTP-клиента
	httpClient := client.New(
		baseUrl,
		time.Duration(timeout)*time.Second,
	)

	//GET-запрос
	rawPost, err := httpClient.Get("/posts/1")
	if err != nil {
		log.Fatalf("Error get post: %v", err)
	}
	fmt.Printf(
		"Получен пост:\nОтвет: %s\nКод ответа: %d\nВремя ответа: %v\n",
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
	fmt.Printf("Создан пост:\nКод ответа: %d\nВремя ответа: %v\n%s",
		createdPost.StatusCode,
		createdPost.Duration,
		string(createdPost.Body),
	)
}
