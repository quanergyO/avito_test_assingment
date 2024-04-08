package postgres

import (
	"avito_test_assingment/types"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type Auth struct {
	db *sqlx.DB
}

func NewAuth(db *sqlx.DB) *Auth {
	return &Auth{db: db}
}

func (r *Auth) CreateUser(user types.UserType) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (username, password_hash, role) VALUES ($1, $2, $3) RETURNING ID", userTable)
	row := r.db.QueryRow(query, user.Username, user.Password, user.Role)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *Auth) GetUser(username, password string) (types.UserType, error) {
	var user types.UserType
	query := fmt.Sprintf("SELECT id, role FROM %s WHERE username=$1 AND password_hash=$2", userTable)
	err := r.db.Get(&user, query, username, password)

	return user, err
}
