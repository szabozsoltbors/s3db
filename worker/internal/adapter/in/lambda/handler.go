package lambda

import (
	"context"
	"encoding/json"
	"fmt"

	uuid "github.com/google/uuid"
	port "github.com/szabozsoltbors/s3db/internal/port/in"
)

const (
	CommandInsertOne = "insertOne"
	CommandDeleteOne = "deleteOne"
	CommandFind      = "find"
)

type Handler struct {
	create port.CreateObject
	delete port.DeleteObject
	list   port.ListObjects
}

type Filter struct {
	ID string `json:"_id"`
}

func NewHandler(create port.CreateObject, delete port.DeleteObject, list port.ListObjects) *Handler {
	return &Handler{
		create: create,
		delete: delete,
		list:   list,
	}
}

func (h *Handler) Handle(ctx context.Context, event Event) (string, error) {

	if event.Command == CommandInsertOne {
		return h.handleInsertOne(ctx, event)
	}

	if event.Command == CommandDeleteOne {
		var filter Filter
		json.Unmarshal(event.Filter, &filter)

		return h.handleDelete(ctx, event, filter.ID)
	}

	if event.Command == CommandFind {
		return h.handleList(ctx, event)
	}

	return "Unsupported command provided: " + event.Command + "\n", nil
}

func (h *Handler) handleInsertOne(ctx context.Context, event Event) (string, error) {
	uuid := uuid.New()
	key := event.Collection + "/" + uuid.String()

	createErr := h.create.Upload(ctx, key, event.Data)
	if createErr != nil {
		return "", createErr
	}

	return "Command: " + event.Command + "\n" +
		"document inserted with name: " + key + "\n", nil
}

func (h *Handler) handleDelete(ctx context.Context, event Event, id string) (string, error) {
	if id == "" {
		return "", fmt.Errorf("missing id in filter for deleteOne command")
	}

	key := event.Collection + "/" + id

	deleteErr := h.delete.Delete(ctx, key)
	if deleteErr != nil {
		return "", deleteErr
	}

	return "Command: " + event.Command + "\n" +
		"document deleted with name: " + key + "\n", nil
}

func (h *Handler) handleList(ctx context.Context, event Event) (string, error) {
	listResult, listErr := h.list.List(ctx, event.Collection)
	if listErr != nil {
		return "", listErr
	}

	return "Command: " + event.Command + "\n" +
		"List result: " + fmt.Sprint(listResult) + "\n", nil
}
