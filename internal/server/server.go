package server

import (
	"FindVibeGo/internal/database"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	_ "github.com/joho/godotenv/autoload"
)

type Server struct {
	port int

	db database.Service
}

func NewServer() *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	NewServer := &Server{
		port: port,

		db: database.New(),
	}

	err := NewServer.db.CreateTables()
	if err != nil {
		return nil
	}

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
