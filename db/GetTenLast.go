package database

import (
	"context"
	"rest_api_users/models"

	"github.com/jackc/pgx/v5"
)

func GetTenLast(ctx context.Context, conn pgx.Conn) ([]models.User, error) {
	sqlQuery := `
	SELECT * FROM users
	ORDER BY id DESC
	LIMIT 10
	`

	var users []models.User

	rows, err := conn.Query(ctx, sqlQuery)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var user models.User

		if err := rows.Scan(&user.ID,
			&user.Name,
			&user.Email,
			&user.Created_At,
		); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}
