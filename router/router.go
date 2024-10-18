package router

import (
	"context"
	"net/http"

	"github.com/akthrmsx/todo-app-go/config"
	"github.com/go-chi/chi/v5"
)

func NewRouter(ctx context.Context, cfg *config.Config) (http.Handler, error) {
	r := chi.NewRouter()

	r.Get("/health", health)

	return r, nil
}
