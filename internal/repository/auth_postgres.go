package repository

import (
	"fmt"

	"github.com/auth_test/internal/entity"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user entity.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (username, password_hash) values ($1, $2) RETURNING id", usersTable)
	row := r.db.QueryRow(query, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, nil
	}
	return id, nil
}

func (r *AuthPostgres) GetUser(username, password string) (entity.User, error) {
	var user entity.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE username=$1 AND password_hash=$2", usersTable)
	err := r.db.Get(&user, query, username, password)
	fmt.Println("User get is complete", user)

	return user, err
}
