package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
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

func CreateAndSave(originalUrl string) string {
	ShortUrl := GenerateShortUrl(originalUrl)

	id := ShortUrl
	newData := URL{
		ID:          id,
		OriginalURL: originalUrl,
		ShortUrl:    ShortUrl,
		CreatedAt:   time.Now(),
	}

	urlDB[ShortUrl] = newData
	return ShortUrl
}

func GetURL(id string) (URL, error) {

	url, err := urlDB[id]

	if !err {
		return URL{}, errors.New("URL NOT FOUND")
	}

	return url, nil
}

func ShortenURLController(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Url string `json:"url"`
	}

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}

	shortURL := CreateAndSave(data.Url)

	response := struct {
		ShrtUrl string `json:"shrtUrl"`
	}{ShrtUrl: shortURL}

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(response)

	fmt.Println(shortURL)
}

func main() {
	fmt.Println("Setting Up Project.")

	http.HandleFunc("/shorten", ShortenURLController)

	fmt.Println("Starting the server on PORT 3000")
	err := http.ListenAndServe(":3000", nil)

	if err != nil {
		log.Fatal(err)
		return
	}

}
