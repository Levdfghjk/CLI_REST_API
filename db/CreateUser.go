package database

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func CreateUser(ctx context.Context, conn pgx.Conn, name string, email string) error {
	sqlQuery := `
	INSERT INTO users (name, email) VALUES ($1, $2)
	`

	_, err := conn.Exec(ctx, sqlQuery, name, email)
	if err != nil {
		return err
	}

	return nil
}
