package handlers

import (
	"encoding/json"
	"net/http"
	database "rest_api_users/db"

	"github.com/jackc/pgx/v5"
)

func GetTenLastHandler(conn pgx.Conn) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := database.GetTenLast(r.Context(), conn)
		if err != nil {
			http.Error(w, "failed to get users", http.StatusInternalServerError)
			return
		}

		if err := json.NewEncoder(w).Encode(users); err != nil {
			http.Error(w, "invalid JSON", http.StatusBadRequest)
		}
	}
}
