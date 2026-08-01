package lambda

import (
	"context"

	port "github.com/szabozsoltbors/s3db/internal/port/in"
)

type Handler struct {
    create port.CreateObject
    delete port.DeleteObject
    list   port.ListObjects
}

func NewHandler(create port.CreateObject, delete port.DeleteObject, list port.ListObjects) *Handler {
	return &Handler{
		create: create,
		delete: delete,
		list:   list,
	}
}

func (h *Handler) Handle(ctx context.Context) (string, error) {
	createResult, createErr := h.create.Execute(ctx)
	if createErr != nil {
		return "", createErr
	}

	deleteResult, deleteErr := h.delete.Execute(ctx)
	if deleteErr != nil {
		return "", deleteErr
	}

	listResult, listErr := h.list.Execute(ctx)
	if listErr != nil {
		return "", listErr
	}

	return "Hello, s3db! \n" +
		"Create result: " + createResult + "\n" +
		"Delete result: " + deleteResult + "\n" +
		"List result: " + listResult + "\n", nil
}
