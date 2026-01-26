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

	r.HandleFunc("/api/titles", transport.PostTitlesHandler).Methods("POST")
	r.HandleFunc("/api/home/{code}", transport.GetHomeHandler).Methods("GET")

	r.HandleFunc("/api/party", transport.CreatePartyHandler).Methods("POST")
	r.HandleFunc("/api/party/{code}", transport.GetPartyHandler).Methods("GET")

	r.HandleFunc("/api/post/{code}", transport.CreatePostHandler).Methods("POST")
	// r.HandleFunc("/api/post/{code}", transport.UpdatePostHandler).Methods("PUT")
	// r.HandleFunc("/api/post/{code}", transport.DeletePostHandler).Methods("DELETE")

	r.HandleFunc("/api/guest/{code}", transport.CreateGuestHandler).Methods("POST")
	r.HandleFunc("/api/guest/{code}", transport.UpdateGuestHandler).Methods("PUT")
	r.HandleFunc("/api/guests/{code}", transport.GetGuestsHandler).Methods("GET")

	r.HandleFunc("/api/unsubscribe", transport.CreateUnsubscribeHandler).Methods("POST")
	r.HandleFunc("/api/confirm", transport.CreateConfirmHandler).Methods("POST")

	// admin
	r.HandleFunc("/api/announcement/{code}", transport.CreateAnnouncementHandler).Methods("POST")
	r.HandleFunc("/api/admin/party/{code}", transport.UpdatePartyHandler).Methods("PUT")


	fmt.Println(env.AllowedOrigins)
	cors := handlers.CORS(
        handlers.AllowedOrigins(env.AllowedOrigins), // React dev server
        handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
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


