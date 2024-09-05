package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"time"
)

type URL struct {
	ID          string    `json:"id"`
	OriginalURL string    `json:"original_url"`
	ShortUrl    string    `json:"short_url"`
	CreatedAt   time.Time `json:"created_at"`
}

var urlDB = make(map[string]URL)

func GenerateShortUrl(OriginalURL string) string {
	hasher := md5.New()

	hasher.Write([]byte(OriginalURL))

	data := hasher.Sum(nil)

	hash := hex.EncodeToString(data)

	return hash[:9]
}

func main() {
	fmt.Println("Setting Up Project.")
	shrt := GenerateShortUrl("https://helloworld.com")
	fmt.Println(shrt)
}
