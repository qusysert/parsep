package main

import (
	_ "github.com/gocolly/colly"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"parser/internal/app/handler"
	"parser/internal/app/pkg"
	"parser/internal/app/pkg/config"
	"parser/internal/app/repository"
	"parser/internal/app/service"
	"parser/pkg/render_client"
	"time"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("cannot load cfg: ", err)
	}
	parser := pkg.New(cfg.ParserTimeoutSeconds)
	repo := repository.New(parser.Collector)
	renderer := render_client.New(cfg.RendererHost, cfg.RendererPort)
	srv := service.New(repo, renderer)
	hdl := handler.New(srv)

	router := mux.NewRouter()
	// Setting timeout for the server
	server := &http.Server{
		Addr:         ":" + cfg.HttpPort,
		ReadTimeout:  600 * time.Second,
		WriteTimeout: 600 * time.Second,
		Handler:      router,
	}
	// Linking addresses and handlers
	for _, rec := range [...]struct {
		route   string
		handler http.HandlerFunc
	}{
		{route: "/getTable", handler: hdl.GetTableHandler},
	} {
		router.HandleFunc(rec.route, rec.handler)
	}

	http.Handle("/", router)

	log.Printf("Server started on port %s \n", cfg.HttpPort)

	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err.Error())
	}
}
