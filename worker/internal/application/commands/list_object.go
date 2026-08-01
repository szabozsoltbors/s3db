package commands

import (
	"context"

	port "github.com/szabozsoltbors/s3db/internal/port/out"
)

type ListCommand struct {
    port port.List
}

func NewListCommand(port port.List) *ListCommand {
	return &ListCommand{
		port: port,
	}
}

func (s *ListCommand) List(ctx context.Context, prefix string) ([]string, error) {
	result, err := s.port.List(ctx, prefix)
	if err != nil {
		return nil, err
	}

	return result, nil
}
