package main

import (
	"fmt"
	"log"
	"time"

	"github.com/Vladimir-Cha/public_api_proxy/internal/storage/client"
)

func main() {
	//Инициализация HTTP-клиента
	httpClient := client.New(
		"https://jsonplaceholder.typicode.com",
		10*time.Second,
	)

	//GET-запрос
	rawPost, err := httpClient.Get("/posts/1")
	if err != nil {
		log.Fatalf("Error get post: %v", err)
	}
	fmt.Printf("Получен пост:\n%s\n", string(rawPost))

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
	fmt.Printf("Создан пост:\n%s\n", createdPost)
}
