package main

import (
	"log"
	"sync"
)

// A simple in-memory database
var (
	db   map[string]MCPContext
	lock sync.RWMutex
)

// connectToDB initializes the in-memory database.
func connectToDB() {
	db = make(map[string]MCPContext)
	log.Println("In-memory database initialized")
}
