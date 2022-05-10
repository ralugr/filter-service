package handlers

import (
	"encoding/json"
	"github.com/ralugr/filter-service/internal/adapter"
	"github.com/ralugr/filter-service/internal/logger"
	"github.com/ralugr/filter-service/internal/model"
	"github.com/ralugr/filter-service/internal/processor"
	"io"
	"net/http"
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
		w.WriteHeader(500)
		return
	}

	// Should check if the message contains an ID, otherwise the message is not valid
	message, err := adapter.ConvertByteArrayToMessage(msg)
	if err != nil {
		w.WriteHeader(400)
		return
	}
	h.processor.FilterMessage(message)
}

func (h *Handlers) RejectedMessages(w http.ResponseWriter, r *http.Request) {
	msg, err := h.processor.GetMessages(model.Rejected)
	if err != nil {
		logger.Warning.Printf("Could not get rejected messages %v", err)
	}

	rsp, err := json.Marshal(msg)
	if err != nil {
		logger.Warning.Printf("Could not convert message %v to json %v", msg)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(rsp)

	if err != nil {
		logger.Warning.Printf("Could not write data to response %v", err)
		return
	}
}

func (h *Handlers) QueuedMessages(w http.ResponseWriter, r *http.Request) {
	msg, err := h.processor.GetMessages(model.Queued)
	if err != nil {
		logger.Warning.Printf("Could not get queued messages %v", err)
	}

	rsp, err := json.Marshal(msg)
	if err != nil {
		logger.Warning.Printf("Could not convert message %v to json %v", msg)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(rsp)

	if err != nil {
		logger.Warning.Printf("Could not write data to response %v", err)
		return
	}
}
