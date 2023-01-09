package repository

import (
	"crypto/sha1"
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"todo-list/internal/entity"

	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (u *UserRepository) generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash)
}

func (u *UserRepository) GetByUsernameAndPassword(username, password string) (int, error) {
	var id int
	query := fmt.Sprintf("SELECT id FROM %s WHERE username=$1 AND password_hash=$2;", userTable)
	if err := u.db.Get(&id, query, username, u.generatePasswordHash(password)); err != nil {
		switch err {
		case sql.ErrNoRows:
			return 0, errors.New("user with this username and password not found")
		default:
			return 0, errors.New(fmt.Sprintf("error: %s", err.Error()))
		}
	}
	return id, nil
}

func (u *UserRepository) GetById(id int) (entity.User, error) {
	var user entity.User
	query := fmt.Sprintf("SELECT id, username, first_name, last_name, email FROM %s WHERE id=$1;", userTable)
	if err := u.db.Get(&user, query, id); err != nil {
		switch err {
		case sql.ErrNoRows:
			return user, errors.New("user with this id not found")
		default:
			return user, errors.New(fmt.Sprintf("error: %s", err.Error()))
		}
	}
	return user, nil
}

func (u *UserRepository) Create(user entity.UserCreate) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (username, first_name, last_name, email, password_hash) VALUES ($1, $2, $3, $4, $5) RETURNING id;", userTable)
	row := u.db.QueryRow(query, user.Username, user.FirstName, user.LastName, user.Email, u.generatePasswordHash(user.Password))
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (u *UserRepository) Update(id int, user entity.UserUpdate) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1
	if user.Username != nil {
		setValues = append(setValues, fmt.Sprintf("username=$%d", argId))
		args = append(args, user.Username)
		argId++
	}
	if user.FirstName != nil {
		setValues = append(setValues, fmt.Sprintf("first_name=$%d", argId))
		args = append(args, user.FirstName)
		argId++
	}
	if user.LastName != nil {
		setValues = append(setValues, fmt.Sprintf("last_name=$%d", argId))
		args = append(args, user.LastName)
		argId++
	}
	if user.Email != nil {
		setValues = append(setValues, fmt.Sprintf("email=$%d", argId))
		args = append(args, user.Email)
		argId++
	}
	queryValues := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE %s SET %s WHERE id=$%d;", userTable, queryValues, argId)
	args = append(args, id)
	_, err := u.db.Exec(query, args...)
	return err
}

func (u *UserRepository) Delete(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1;", userTable)
	_, err := u.db.Exec(query, id)
	return err
}
