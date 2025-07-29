package main

import (
	"time"
)

type MCPContext struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Data      any       `json:"data"` // você pode trocar por um tipo específico se quiser
}
