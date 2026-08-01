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

func (s *CreateCommand) Upload(ctx context.Context, key string, data []byte) error {
	return s.port.Upload(ctx, key, data)
}
