package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load() // Carrega .env local

	connectToDB()

	r := mux.NewRouter()

	// Ping handler for health check
	r.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("pong"))
	}).Methods("GET")

	// API routes
	r.HandleFunc("/mcp", createMCPHandler).Methods("POST")
	r.HandleFunc("/mcp", getAllMCPHandler).Methods("GET")
	r.HandleFunc("/mcp/{id}", getMCPHandler).Methods("GET")
	r.HandleFunc("/mcp/{id}", updateMCPHandler).Methods("PUT")
	r.HandleFunc("/mcp/{id}", deleteMCPHandler).Methods("DELETE")

	// Serve static files
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./frontend"))))
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "frontend/index.html")
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port
	}

	log.Printf("Server starting on port %s", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal(err)
	}
}
