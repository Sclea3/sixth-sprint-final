package main

import (
	"fmt"
	"net/http"

	"github.com/Yandex-Practicum/go-rest-api-homework/handlers"
	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	r.Get("/tasks", handlers.GetTasks)
	r.Post("/tasks", handlers.PostTask)
	r.Get("/tasks/{id}", handlers.GetTaskByID)
	r.Delete("/tasks/{id}", handlers.DeleteTaskByID)

	if err := http.ListenAndServe(":8080", r); err != nil {
		fmt.Printf("Ошибка при запуске сервера: %s", err.Error())
		return
	}
}
