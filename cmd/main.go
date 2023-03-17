package main

import (
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly"
	_ "github.com/gocolly/colly"
	"log"
	"parser/internal/app/repository"
	"parser/internal/app/service"
	model "parser/internal/model"
	"parser/pkg/render_client"
	"time"
)

func main() {
	repo := repository.New()
	renderer := render_client.New("localhost", 3000)
	srv := service.New(repo, renderer)
	start := time.Now()

	collector := colly.NewCollector()
	collector.SetRequestTimeout(120 * time.Second)
	collector.OnRequest(func(r *colly.Request) {
		log.Println("Visiting", r.URL)
	})

	collector.OnResponse(func(r *colly.Response) {
		log.Println("Got a response from", r.Request.URL)
	})

	collector.OnError(func(r *colly.Response, e error) {
		log.Println("Got this error:", e)
	})

	materials, err := srv.FormTableContent(*collector, model.ParsePool)
	if err != nil {
		log.Fatalf("cant get table content: %v", err)
	}

	jsonString, _ := json.MarshalIndent(materials, "", "  ")

	// Print the JSON string with newlines
	fmt.Println(string(jsonString))
	fmt.Println(srv.GenTableHTML(materials))
	fmt.Printf("Scraped in %v\n", time.Since(start))
}
