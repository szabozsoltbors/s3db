package lambda

import (
	"context"
	"fmt"

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

func (h *Handler) Handle(ctx context.Context, event Event) (string, error) {
	createErr := h.create.Upload(ctx, event.Key, event.Data)
	if createErr != nil {
		return "", createErr
	}

	listResultBeforeDelete, listErr := h.list.List(ctx, "")
	if listErr != nil {
		return "", listErr
	}

	deleteErr := h.delete.Delete(ctx, event.Key)
	if deleteErr != nil {
		return "", deleteErr
	}

	listResultAfterDelete, listErr := h.list.List(ctx, "")
	if listErr != nil {
		return "", listErr
	}

	return "Hello, s3db! \n" +
		"File created with name: " + event.Key + "\n" +
		"List result before delete: " + fmt.Sprint(listResultBeforeDelete) + "\n" +
		"Delete result: " + event.Key + "\n" +
		"List result after delete: " + fmt.Sprint(listResultAfterDelete) + "\n", nil
}
