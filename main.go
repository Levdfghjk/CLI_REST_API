package main

import (
	"context"
	"fmt"
	"net/http"
	"rest_api_users/handlers"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5"
)

func main() {
	ctx := context.Background()
	router := mux.NewRouter()
	conn, err := pgx.Connect(ctx, "postgres://postgres:1234@localhost:5432/postgres")

	if err != nil {
		panic(err)
	}
	defer conn.Close(ctx)

	router.HandleFunc("/users", handlers.CreateUserHandler(*conn)).Methods("POST")
	router.HandleFunc("/users/{id}", handlers.GetUserByIDHandler(*conn)).Methods("GET")
	router.HandleFunc("/users/{id}", handlers.DeleteUserHandler(*conn)).Methods("DELETE")
	router.HandleFunc("/users", handlers.GetTenLastHandler(*conn)).Methods("GET")

	fmt.Println("server started on :9091")

	if err := http.ListenAndServe(":9091", router); err != nil {
		fmt.Println("Ошибка запуска сервера")
	}
}
