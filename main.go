package main

import (
	"fmt"
	"time"
)

type URL struct {
	ID          string    `json:"id"`
	OriginalURL string    `json:"original_url"`
	ShortUrl    string    `json:"short_url"`
	CreatedAt   time.Time `json:"created_at"`
}

func main() {
	fmt.Println("Setting Up Project.")
}
