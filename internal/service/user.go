package service

import (
	"todo-list/internal/entity"
	"todo-list/internal/repository"
)

type UserService struct {
	repo repository.User
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{repo: repo}
}

func (u *UserService) GetByUsernameAndPassword(username, password string) (int, error) {
	return u.repo.GetByUsernameAndPassword(username, password)
}

func (u *UserService) Create(user entity.UserCreate) (int, error) {
	return u.repo.Create(user)
}

func (u *UserService) GetById(id int) (entity.User, error) {
	return u.repo.GetById(id)
}

func (u *UserService) Update(id int, user entity.UserUpdate) error {
	if err := user.Validate(); err != nil {
		return err
	}
	return u.repo.Update(id, user)
}

func (u *UserService) Delete(id int) error {
	return u.repo.Delete(id)
}
