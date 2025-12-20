package routes

import (
	"net/http"

	"GO_Auth/handlers"
	"GO_Auth/middleware"
)

func RegisterRoutes() http.Handler{
	mux := http.NewServeMux()

	mux.HandleFunc("/auth/signup", handlers.Signup)
	mux.HandleFunc("/auth/login", handlers.Login)

	mux.Handle(
		"/profile",
		 middleware.AuthMiddleware(http.HandlerFunc(handlers.GetProfile)),
	)
	return mux
}