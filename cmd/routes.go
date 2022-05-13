package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/ralugr/filter-service/internal/handlers"
	"net/http"
)

// routes is used for mapping the routes to their handlers
func routes(h *handlers.Handlers) http.Handler {
	// Lightweight router for http services
	mux := chi.NewRouter()
	//Gracefully absorb panics and prints the stack trace
	mux.Use(middleware.Recoverer)

	mux.Get("/", h.Home)
	mux.Post("/filter_message", h.FilterMessage)

	// Contains notification encapsulating the new banned word list provided by the language-service
	mux.Post("/notify", h.Notify)
	mux.Get("/rejected_messages", h.RejectedMessages)
	mux.Get("/queued_messages", h.QueuedMessages)

	return mux
}
