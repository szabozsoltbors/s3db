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

func (s *DeleteCommand) Execute(ctx context.Context) (string, error) {
	result, err := s.port.Execute(ctx)
	if err != nil {
		return "", err
	}

	return result, nil
}
