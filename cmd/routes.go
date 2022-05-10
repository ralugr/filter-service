package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/ralugr/filter-service/internal/handlers"
	"net/http"
)

func routes(h *handlers.Handlers) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)

	mux.Get("/", h.Home)
	mux.Post("/filter_message", h.FilterMessage)
	mux.Get("/rejected_messages", h.RejectedMessages)
	mux.Get("/queued_messages", h.QueuedMessages)

	return mux
}
