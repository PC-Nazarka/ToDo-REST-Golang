package service

import (
	"todo-list/internal/entity"
	"todo-list/internal/repository"
)

type CommentService struct {
	repo repository.Comment
}

func NewCommentService(repo repository.Comment) *CommentService {
	return &CommentService{repo: repo}
}

func (c *CommentService) Create(userId int, comment entity.CommentCreate) (int, error) {
	return c.repo.Create(userId, comment)
}

func (c *CommentService) GetById(id int) (entity.Comment, error) {
	return c.repo.GetById(id)
}

func (c *CommentService) GetByPostId(id int) ([]entity.Comment, error) {
	return c.repo.GetByPostId(id)
}

func (c *CommentService) Update(id int, comment entity.CommentUpdate) error {
	if err := comment.Validate(); err != nil {
		return err
	}
	return c.repo.Update(id, comment)
}

func (c *CommentService) Delete(id int) error {
	return c.repo.Delete(id)
}
