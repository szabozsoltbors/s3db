package in

import "context"

type CreateObject interface {
	Upload(ctx context.Context, key string, data []byte) error
}

type DeleteObject interface {
	Delete(ctx context.Context, key string) error
}

type ListObjects interface {
	List(ctx context.Context, prefix string) ([]string, error)
}
