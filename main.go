package main

import (
	"log"
	"net/http"
	"todolist-service/config"
	"todolist-service/handlers"
	"todolist-service/repositories"
	"todolist-service/services"

	"github.com/gorilla/mux"
)

func main() {
	config.InitDB()
	repo := &repositories.TaskRepository{DB: config.DB}
	service := &services.TaskService{Repo: repo}
	handler := &handlers.TaskHandler{Service: service}

	router := mux.NewRouter()
	router.HandleFunc("/tasks", handler.CreateTask).Methods("POST")
	router.HandleFunc("/tasks", handler.GetTasks).Methods("GET")
	router.HandleFunc("/tasks/{id}", handler.UpdateTask).Methods("PUT")
	router.HandleFunc("/tasks/{id}", handler.DeleteTask).Methods("DELETE")

	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
