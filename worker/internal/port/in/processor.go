package in

import "context"

type CreateObject interface {
	Execute(ctx context.Context) (string, error)
}

type DeleteObject interface {
	Execute(ctx context.Context) (string, error)
}

type ListObjects interface {
	Execute(ctx context.Context) (string, error)
}
