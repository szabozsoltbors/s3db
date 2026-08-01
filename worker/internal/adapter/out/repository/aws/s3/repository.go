package s3

import (
	"context"
)

type Repository struct {
}

func NewRepository() *Repository {
	return &Repository{
	}
}

func (r *Repository) Execute(ctx context.Context) (string, error) {
	return "Object handled successfully by the repository layer!", nil
}
