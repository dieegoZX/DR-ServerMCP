package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

// createMCPHandler handles the creation of a new MCPContext.
func createMCPHandler(w http.ResponseWriter, r *http.Request) {
	var data any
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	lock.Lock()
	defer lock.Unlock()

	id := uuid.New().String()
	mcpCtx := MCPContext{
		ID:        id,
		CreatedAt: time.Now(),
		Data:      data,
	}
	db[id] = mcpCtx

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(mcpCtx)
}

// getAllMCPHandler handles retrieving all MCPContext objects.
func getAllMCPHandler(w http.ResponseWriter, r *http.Request) {
	lock.RLock()
	defer lock.RUnlock()

	contexts := make([]MCPContext, 0, len(db))
	for _, ctx := range db {
		contexts = append(contexts, ctx)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(contexts)
}

// getMCPHandler handles retrieving a single MCPContext by its ID.
func getMCPHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	lock.RLock()
	defer lock.RUnlock()

	ctx, ok := db[id]
	if !ok {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ctx)
}

// updateMCPHandler handles updating an existing MCPContext.
func updateMCPHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	lock.Lock()
	defer lock.Unlock()

	_, ok := db[id]
	if !ok {
		http.NotFound(w, r)
		return
	}

	var data any
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// For simplicity, we'll just update the Data field.
	updatedCtx := MCPContext{
		ID:        id,
		CreatedAt: db[id].CreatedAt, // Keep original creation time
		Data:      data,
	}
	db[id] = updatedCtx

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedCtx)
}

// deleteMCPHandler handles deleting an MCPContext.
func deleteMCPHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	lock.Lock()
	defer lock.Unlock()

	if _, ok := db[id]; !ok {
		http.NotFound(w, r)
		return
	}

	delete(db, id)

	w.WriteHeader(http.StatusNoContent)
}
