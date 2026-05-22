package handlers

import (
	"net/http"
	database "rest_api_users/db"
	"strconv"
	"strings"

	"github.com/jackc/pgx/v5"
)

func DeleteUserHandler(conn pgx.Conn) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := strings.TrimPrefix(r.URL.Path, "/user/")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "problems with conversation", http.StatusMethodNotAllowed)
			return
		}

		if err := database.DeleteUser(r.Context(), conn, id); err != nil {
			http.Error(w, "problems with DB", http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("user deleted successfully"))
	}

}
