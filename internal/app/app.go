package app

import (
	"github.com/MoXcz/personal-webiste/internal/middleware"
	"net/http"
	"time"
)

func NewServer() *http.Server {
	mux := http.NewServeMux()

	// Define routes
	RegisterRoutes(mux)

	// Apply middleware
	secureMux := middleware.SecurityHeaders(mux)
	loggedMux := middleware.Logging(secureMux)
	limitedMux := middleware.RateLimit(loggedMux)

	return &http.Server{
		Addr:         ":8080",
		Handler:      limitedMux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}
}
