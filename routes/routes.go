package routes

import (
	"net/http"

	"GO_Auth/handlers"
	"GO_Auth/middleware"
)

func RegisterRoutes() http.Handler {
	mux := http.NewServeMux()

	// Health-check endpoint used by fly.io to verify the app is running.
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	mux.HandleFunc("/auth/signup", handlers.Signup)
	mux.HandleFunc("/auth/login", handlers.Login)

	mux.Handle(
		"/profile",
		middleware.AuthMiddleware(http.HandlerFunc(handlers.GetProfile)),
	)
	return mux
}