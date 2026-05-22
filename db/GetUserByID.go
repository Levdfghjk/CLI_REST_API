package database

import (
	"context"
	"rest_api_users/models"

	"github.com/jackc/pgx/v5"
)

func GetUserByID(ctx context.Context, conn pgx.Conn, id int) (models.User, error) {
	sqlQuery := `
	SELECT id, name, email, created_at FROM users
	WHERE id = $1
	`

	var user models.User

	if err := conn.QueryRow(ctx, sqlQuery, id).Scan(&user.ID, &user.Name, &user.Email, &user.Created_At); err != nil {
		return models.User{}, err
	}

	return user, nil
}
