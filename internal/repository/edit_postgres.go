package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type EditPostgres struct {
	db *sqlx.DB
}

func NewEditPostgres(db *sqlx.DB) *EditPostgres {
	return &EditPostgres{db: db}
}

func (r *EditPostgres) DeleteUser(username string) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE username=$1", usersTable)
	r.db.QueryRow(query, username)
	fmt.Println("User delete is complete", username)
	return nil
}
