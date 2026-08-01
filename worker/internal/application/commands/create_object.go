package commands

import (
	"context"

	port "github.com/szabozsoltbors/s3db/internal/port/out"
)

type CreateCommand struct {
    port port.Save
}

func NewCreateCommand(port port.Save) *CreateCommand {
	return &CreateCommand{
		port: port,
	}
}

func (s *CreateCommand) Execute(ctx context.Context) (string, error) {
	result, err := s.port.Execute(ctx)
	if err != nil {
		return "", err
	}

	return result, nil
}
