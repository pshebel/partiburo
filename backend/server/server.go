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

	r.HandleFunc("/api/party", transport.CreatePartyHandler).Methods("POST")
	r.HandleFunc("/api/party", transport.GetPartyHandler).Methods("GET")
	r.HandleFunc("/api/home", transport.GetHomeHandler).Methods("GET")

	r.HandleFunc("/api/post", transport.CreatePostHandler).Methods("POST")

	r.HandleFunc("/api/guest", transport.CreateGuestHandler).Methods("POST")
	r.HandleFunc("/api/guest", transport.UpdateGuestHandler).Methods("PUT")
	r.HandleFunc("/api/guests", transport.GetGuestsHandler).Methods("GET")

	r.HandleFunc("/api/unsubscribe", transport.CreateUnsubscribeHandler).Methods("POST")
	r.HandleFunc("/api/confirm", transport.CreateConfirmHandler).Methods("POST")



	fmt.Println(env.AllowedOrigins)
	cors := handlers.CORS(
        handlers.AllowedOrigins(env.AllowedOrigins), // React dev server
        handlers.AllowedMethods([]string{"GET", "POST", "PUT", "OPTIONS"}),
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


