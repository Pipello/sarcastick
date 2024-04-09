package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"sarcastick/pkg/news"

	"golang.org/x/net/html"
)

func main() {
	response, err := http.Get("https://www.expats.cz")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		log.Fatalf("HTTP request failed with status code %d", response.StatusCode)
	}

	htmlBody, err := html.Parse(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	link := news.FindNewsLink(htmlBody)
	if link == "" {
		log.Fatal("No link found")
	}

	response, err = http.Get("https://www.expats.cz" + link)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		log.Fatalf("HTTP request failed with status code %d", response.StatusCode)
	}

	htmlBody, err = html.Parse(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	ct := news.ExtractContentWithTitle(htmlBody)
	file, err := os.Create("news.json")
	if err != nil {
		log.Fatal(err)
	}
	err = json.NewEncoder(file).Encode(ct)
	if err != nil {
		log.Fatal(err)
	}
}
