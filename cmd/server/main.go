package main

import (
	"fmt"
	"log"
	"time"

	"github.com/Vladimir-Cha/public_api_proxy/internal/storage/service"

	"github.com/Vladimir-Cha/public_api_proxy/internal/storage/client"
)

func main() {
	//Инициализация HTTP-клиента
	httpClient := client.New(
		"https://jsonplaceholder.typicode.com",
		10*time.Second,
	)

	//Создание сервиса
	placeholderService := service.New(httpClient)

	//GET-запрос
	rawPost, err := placeholderService.GetPost(1)
	if err != nil {
		log.Fatalf("Error get post: %v", err)
	}
	fmt.Printf("Получен пост:\n%s\n", rawPost)

	//POST-запрос
	newPost := []byte(`{
        "title": "Новый пост",
        "body": "Содержание поста",
        "userId": 1
    }`)

	createdPost, err := placeholderService.Create(newPost)
	if err != nil {
		log.Fatalf("Error create post: %v", err)
	}
	fmt.Printf("Создан пост:\n%s\n", createdPost)
}
