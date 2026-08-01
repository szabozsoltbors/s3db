package out

import "context"

type Save interface {
	Execute(ctx context.Context) (string, error)
}

type Delete interface {
	Execute(ctx context.Context) (string, error)
}

type List interface {
	Execute(ctx context.Context) (string, error)
}
