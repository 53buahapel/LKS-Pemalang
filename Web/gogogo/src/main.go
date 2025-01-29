package main

import (
	"crypto/rand"
	"encoding/hex"
	"gogogo/config"
	"gogogo/database"
	"log"
	"net/http"

	"github.com/gorilla/sessions"
)

func main() {
	mux := http.NewServeMux()
	store := sessions.NewCookieStore([]byte(genRandKey()))
	session := sessions.NewSession(store, "session")

	database.InitDB()
	config.Bootstrap(&config.BootstrapConfig{
		Mux:     mux,
		Session: session,
	})

	server := http.Server{
		Addr:    ":1337",
		Handler: mux,
	}

	log.Println("Server is running on port 1337")
	server.ListenAndServe()
}

func genRandKey() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatal(err)
	}
	return hex.EncodeToString(b)
}
