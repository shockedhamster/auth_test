package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type EditPostgres struct {
	db *sqlx.DB
}

func NewEditPostgres(db *sqlx.DB) *EditPostgres {
	return &EditPostgres{db: db}
}

func (r *EditPostgres) DeleteUser(username string) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE username=$1 RETURNING id", usersTable)
	row := r.db.QueryRow(query, username)

	if err := row.Scan(); err != nil {
		logrus.Errorf("Error deleting a user: %s", err.Error())
	}

	return nil
}

func (r *EditPostgres) UpdateUsername(usernameOld, usernameNew string) (int, error) {
	var id int
	query := fmt.Sprintf("UPDATE %s SET username=$1 WHERE username=$2 RETURNING id", usersTable)
	row := r.db.QueryRow(query, usernameNew, usernameOld)
	if err := row.Scan(&id); err != nil {
		logrus.Errorf("Error updating username: %s", err.Error())
		return 0, err
	}
	return id, nil

}
