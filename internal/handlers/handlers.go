package handlers

import (
	"io"
	"net/http"

	"github.com/ralugr/filter-service/internal/config"

	"github.com/ralugr/filter-service/internal/adapter"
	"github.com/ralugr/filter-service/internal/logger"
	"github.com/ralugr/filter-service/internal/model"
	"github.com/ralugr/filter-service/internal/processor"
	"github.com/ralugr/filter-service/internal/respond"
)

// Handlers defines all the available handlers and redirect calls to the processor
type Handlers struct {
	processor *processor.Processor
	cfg       *config.Config
}

// New constructor
func New(p *processor.Processor, c *config.Config) *Handlers {
	logger.Info.Println("Creating handlers")

	return &Handlers{
		processor: p,
		cfg:       c,
	}
}

// Home handler, used for testing connection to the service
func (h *Handlers) Home(w http.ResponseWriter, r *http.Request) {
	if _, err := w.Write([]byte("Welcome to the filter service!!")); err != nil {
		logger.Warning.Printf("Could not write welcome message")
		respond.Error(w, 500, "Encountered internal error")
		return
	}
}

// FilterMessage receives a message and redirects it to the processor
func (h *Handlers) FilterMessage(w http.ResponseWriter, r *http.Request) {
	msg, err := io.ReadAll(r.Body)
	if err != nil {
		respond.Error(w, 500, "Encountered internal error")
		return
	}

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

// RejectedMessages communicates with the processor for retrieving the rejected messages
func (h *Handlers) RejectedMessages(w http.ResponseWriter, r *http.Request) {
	msg, err := h.processor.GetMessages(model.Rejected)
	if err != nil {
		logger.Warning.Printf("Could not get rejected messages %v", err)
		respond.Error(w, 200, "Could not get rejected messages")
		return
	}

	respond.Success(w, msg)
}

// QueuedMessages communicates with the processor for retrieving the queued messages
func (h *Handlers) QueuedMessages(w http.ResponseWriter, r *http.Request) {
	msg, err := h.processor.GetMessages(model.Queued)
	if err != nil {
		logger.Warning.Printf("Could not get queued messages %v", err)
		respond.Error(w, 500, "Could not get queued messages")
		return
	}

	respond.Success(w, msg)
}

// Notify handler for receiving banned list updates
func (h *Handlers) Notify(w http.ResponseWriter, r *http.Request) {
	n, err := io.ReadAll(r.Body)
	if err != nil {
		respond.Error(w, 500, "Encountered internal error")
		return
	}

	notification, err := adapter.ConvertByteArrayToNotifyPayload(n)
	if err != nil {
		respond.Error(w, 400, "Failed to decode payload")
		return
	}

	if notification.Token != h.cfg.Token {
		respond.Error(w, 401, "Invalid token")
		return
	}

	logger.Info.Println("Received ", notification)

	bw := model.NewBannedWords(notification.Words)
	if err = h.processor.UpdateBannedWords(bw); err != nil {
		respond.Error(w, 401, "Failed to update banned words")
	}

	respond.Success(w, "Updated words")
}
