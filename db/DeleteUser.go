package database

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
)

func DeleteUser(ctx context.Context, conn pgx.Conn, id int) error {
	sqlQuery := `
	DELETE FROM users
	WHERE id = $1
	`

	result, err := conn.Exec(ctx, sqlQuery, id)
	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return errors.New("User is not found")
	}

	return nil
}
