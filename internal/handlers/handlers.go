package handlers

import (
	"github.com/ralugr/filter-service/internal/config"
	"io"
	"net/http"

	"github.com/ralugr/filter-service/internal/adapter"
	"github.com/ralugr/filter-service/internal/logger"
	"github.com/ralugr/filter-service/internal/model"
	"github.com/ralugr/filter-service/internal/processor"
	"github.com/ralugr/filter-service/internal/respond"
)

type Handlers struct {
	processor *processor.Processor
	cfg       *config.Config
}

func New(p *processor.Processor, c *config.Config) *Handlers {
	logger.Info.Println("Creating handlers")

	return &Handlers{
		processor: p,
		cfg:       c,
	}
}

func (h *Handlers) Home(w http.ResponseWriter, r *http.Request) {
	if _, err := w.Write([]byte("Welcome to the filter service!!")); err != nil {
		logger.Warning.Printf("Could not write welcome message")
		respond.Error(w, 500, "Encountered internal error")
		return
	}
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
	n, err := io.ReadAll(r.Body)
	if err != nil {
		respond.Error(w, 500, "Encountered internal error")
		return
	}

	// Should check if the notification contains a Token, otherwise the request is not valid
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
	h.processor.UpdateBannedWords(bw)

	respond.Success(w, "Updated words")
}
