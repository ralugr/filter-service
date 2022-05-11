package handlers

import (
	"io"
	"net/http"

	"github.com/ralugr/filter-service/internal/adapter"
	"github.com/ralugr/filter-service/internal/logger"
	"github.com/ralugr/filter-service/internal/model"
	"github.com/ralugr/filter-service/internal/processor"
	"github.com/ralugr/filter-service/internal/respond"
)

type Handlers struct {
	processor processor.Processor
}

func New(p processor.Processor) *Handlers {
	logger.Info.Println("Creating handlers")

	return &Handlers{
		processor: p,
	}
}

func (h *Handlers) Home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to the filter service!!</h1>"))
}

func (h *Handlers) FilterMessage(w http.ResponseWriter, r *http.Request) {
	msg, err := io.ReadAll(r.Body)
	if err != nil {
		respond.Error(w, 500, "Encountered internal error")
		return
	}

	// Should check if the message contains an ID, otherwise the message is not valid
	message, err := adapter.ConvertByteArrayToMessage(msg)
	if err != nil {
		respond.Error(w, 400, "Failed to decode payload")
		return
	}

	if message.UID == "" {
		respond.Error(w, 400, "Id field is required")
		return
	}

	h.processor.FilterMessage(&message)

	respond.Success(w, message.State)
}

func (h *Handlers) RejectedMessages(w http.ResponseWriter, r *http.Request) {
	msg, err := h.processor.GetMessages(model.Rejected)
	if err != nil {
		logger.Warning.Printf("Could not get rejected messages %v", err)
		respond.Error(w, 200, "Could not get rejected messages")
		return
	}

	respond.Success(w, msg)
}

func (h *Handlers) QueuedMessages(w http.ResponseWriter, r *http.Request) {
	msg, err := h.processor.GetMessages(model.Queued)
	if err != nil {
		logger.Warning.Printf("Could not get queued messages %v", err)
		respond.Error(w, 500, "Could not get queued messages")
		return
	}

	respond.Success(w, msg)
}

func (h *Handlers) Notify(w http.ResponseWriter, r *http.Request) {
	msg, err := io.ReadAll(r.Body)
	if err != nil {
		respond.Error(w, 500, "Encountered internal error")
		return
	}

	// Should check if the message contains a Token, otherwise the request is not valid
	message, err := adapter.ConvertByteArrayToNotifyPayload(msg)
	if err != nil {
		respond.Error(w, 400, "Failed to decode payload")
		return
	}

	if message.Token != "jkhfkashdk1e1jh76t@5$" {
		respond.Error(w, 401, "Invalid token")
		return
	}

	logger.Info.Println(message)
	respond.Success(w, "Updated words")
}
