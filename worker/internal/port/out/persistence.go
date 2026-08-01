package out

import "context"

type Save interface {
	Upload(ctx context.Context, key string, data []byte) error
}

type Delete interface {
	Delete(ctx context.Context, key string) error
}

type List interface {
	List(ctx context.Context, prefix string) ([]string, error)
}
