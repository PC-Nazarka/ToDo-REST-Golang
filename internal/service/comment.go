package service

import (
	"errors"
	"todo-list/internal/entity"
	"todo-list/internal/repository"
)

type CommentService struct {
	repo     repository.Comment
	postRepo repository.Post
}

func NewCommentService(repo repository.Comment, postRepo repository.Post) *CommentService {
	return &CommentService{repo: repo, postRepo: postRepo}
}

func (c *CommentService) Create(userId int, comment entity.CommentCreate) (int, error) {
	_, err := c.postRepo.GetById(comment.PostId)
	if err != nil {
		return 0, err
	}
	return c.repo.Create(userId, comment)
}

func (c *CommentService) GetById(id int) (entity.Comment, error) {
	return c.repo.GetById(id)
}

func (c *CommentService) GetByPostId(id int) ([]entity.Comment, error) {
	return c.repo.GetByPostId(id)
}

func (c *CommentService) Update(userId, commentId int, comment entity.CommentUpdate) error {
	commentExists, err := c.repo.GetById(commentId)
	if err != nil {
		return err
	}
	if userId != commentExists.UserId {
		return errors.New("you can't update comment another user")
	}
	if err = comment.Validate(); err != nil {
		return err
	}
	return c.repo.Update(commentId, comment)
}

func (c *CommentService) Delete(userId, commentId int) error {
	comment, err := c.repo.GetById(commentId)
	if err != nil {
		return err
	}
	if userId != comment.UserId {
		return errors.New("you can't delete comment another user")
	}
	return c.repo.Delete(commentId)
}
