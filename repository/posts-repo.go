package repository

import (
	"context"
	"lpvdt/api/entity"
)

type PostRepository interface {
	Save(post *entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
}

type repo struct{}

func NewRepository() PostRepository {
	return &repo{}
}

func (*repo) Save(post *entity.Post) {
	ctx := context.Background()
}
