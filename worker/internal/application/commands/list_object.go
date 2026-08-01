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

func (s *ListCommand) Execute(ctx context.Context) (string, error) {
	result, err := s.port.Execute(ctx)
	if err != nil {
		return "", err
	}

	return result, nil
}
