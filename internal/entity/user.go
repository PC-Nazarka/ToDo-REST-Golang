package entity

import "errors"

type User struct {
	Id        int    `json:"id" db:"id"`
	Username  string `json:"username" db:"username" binding:"required"`
	FirstName string `json:"first_name" db:"first_name"`
	LastName  string `json:"last_name" db:"last_name"`
	Email     string `json:"email" db:"email"`
}

type UserCreate struct {
	Username  string `json:"username" db:"username" binding:"required"`
	FirstName string `json:"first_name" db:"first_name"`
	LastName  string `json:"last_name" db:"last_name"`
	Email     string `json:"email" db:"email" binding:"required"`
	Password  string `json:"password,writeOnly,omitempty" binding:"required"`
}

type UserUpdate struct {
	Username  *string `json:"username"`
	FirstName *string `json:"first_name"`
	LastName  *string `json:"last_name"`
	Email     *string `json:"email"`
}

func (u *UserUpdate) Validate() error {
	if u.Email == nil && u.FirstName == nil && u.LastName == nil && u.Username == nil {
		return errors.New("update structure has no values")
	}
	return nil
}
