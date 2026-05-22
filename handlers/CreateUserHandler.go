package handlers

import (
	"encoding/json"
	"net/http"

	database "rest_api_users/db"
	"rest_api_users/models"

	"github.com/jackc/pgx/v5"
)

func CreateUserHandler(conn pgx.Conn) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var newUser models.User

		if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
			http.Error(w, "invalid JSON", http.StatusBadRequest)
			return
		}

		if err := database.CreateUser(r.Context(), conn, newUser.Name, newUser.Email); err != nil {
			http.Error(w, "failed to create user", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)

		json.NewEncoder(w).Encode(map[string]string{
			"message": "user created",
		})
	}
}
