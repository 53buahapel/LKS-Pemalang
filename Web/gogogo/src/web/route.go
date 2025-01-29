package web

import "net/http"

func InitializeRoute(mux *http.ServeMux, handler *Handler) {
	mux.HandleFunc("GET /signin", handler.SignIn)
	mux.HandleFunc("POST /signin", handler.SignedIn)
	mux.HandleFunc("GET /signup", handler.SignUp)
	mux.HandleFunc("POST /signup", handler.SignedUp)
	mux.HandleFunc("GET /dashboard", handler.Dashboard)
	mux.HandleFunc("POST /todos", handler.CreateTodo)
	mux.HandleFunc("POST /fetch", handler.Fetching)
	mux.HandleFunc("GET /debug", handler.Debug)
}
