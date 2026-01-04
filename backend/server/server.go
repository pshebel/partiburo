package server

import (
	"log"
	"net/http"
	"fmt"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
	"github.com/pshebel/partiburo/backend/transport"
	"github.com/pshebel/partiburo/backend/env"
)

var server *http.Server
var startTime = time.Now()

func init() {
	// Create a new router
	r := mux.NewRouter()
	// Apply middleware
	r.Use(loggingMiddleware)

	r.HandleFunc("/{token_hash}/party", transport.GetPartyHandler).Methods("GET")
	r.HandleFunc("/party", transport.CreatePartyHandler).Methods("POST")
	r.HandleFunc("/guests", transport.GetGuestsHandler).Methods("GET")
	r.HandleFunc("/{token_hash}/guest", transport.CreateGuestHandler).Methods("POST")

	r.HandleFunc("/post", transport.CreatePostHandler).Methods("POST")


	fmt.Println(env.AllowedOrigins)
	cors := handlers.CORS(
        handlers.AllowedOrigins(env.AllowedOrigins), // React dev server
        handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"}),
        handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
    )


	server = &http.Server{
		Handler:      cors(r),
		Addr:         fmt.Sprintf(":%s", env.Port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}
	log.Printf("server initialized")
}

// Middleware for logging requests
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s %v", r.Method, r.URL.Path, time.Since(start))
	})
}

func GetServer() (*http.Server) {
	return server
}


