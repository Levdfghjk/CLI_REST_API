package handlers

import (
	"encoding/json"
	"net/http"
	database "rest_api_users/db"
	"strconv"
	"strings"

	"github.com/jackc/pgx/v5"
)

func GetUserByIDHandler(conn pgx.Conn) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := strings.TrimPrefix(r.URL.Path, "/user/")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "problems with conversation", http.StatusMethodNotAllowed)
			return
		}

		user, err := database.GetUserByID(r.Context(), conn, id)
		if err != nil {
			http.Error(w, "cant get user", http.StatusInternalServerError)
			return
		}

		if err := json.NewEncoder(w).Encode(user); err != nil {
			http.Error(w, "invalid JSON", http.StatusBadRequest)
			return
		}
	}
}
