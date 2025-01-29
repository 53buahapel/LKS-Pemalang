package config

import (
	"gogogo/database"
	"gogogo/web"
	"net/http"

	"github.com/gorilla/sessions"
)

type BootstrapConfig struct {
	Mux     *http.ServeMux
	Session *sessions.Session
}

func Bootstrap(config *BootstrapConfig) {
	repo := web.NewRepository(database.DB)
	handler := web.NewHandler(repo, config.Session)
	web.InitializeRoute(config.Mux, handler)
}
