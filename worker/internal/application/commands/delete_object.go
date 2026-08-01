package commands

import (
	"context"

	port "github.com/szabozsoltbors/s3db/internal/port/out"
)

type DeleteCommand struct {
    port port.Delete
}

func NewDeleteCommand(port port.Delete) *DeleteCommand {
	return &DeleteCommand{
		port: port,
	}
}

func (s *DeleteCommand) Delete(ctx context.Context, key string) error {
	return s.port.Delete(ctx, key)
}
