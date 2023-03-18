package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

type Config struct {
	HttpPort string
	HttpHost int

	RendererHost string
	RendererPort int

	ParserTimeoutSeconds int
}

func LoadConfig() (Config, error) {

	err := godotenv.Load()
	if err != nil {
		log.Printf("Error loading .env file")
		panic(err)
	}

	RendererPort, err := strconv.Atoi(os.Getenv("RENDERER_PORT"))
	if err != nil {
		return Config{}, fmt.Errorf("error parsing RENDERER_PORT: %w", err)
	}
	HttpPort, err := strconv.Atoi(os.Getenv("HTTP_PORT"))
	if err != nil {
		return Config{}, fmt.Errorf("error parsing HTTP_PORT: %w", err)
	}
	timeout, err := strconv.Atoi(os.Getenv("PARSER_TIMEOUT_SECONDS"))
	if err != nil {
		return Config{}, fmt.Errorf("error parsing PARSER_TIMEOUT_SECONDS: %w", err)
	}
	config := Config{
		HttpPort:             os.Getenv("HTTP_PORT"),
		HttpHost:             HttpPort,
		RendererHost:         os.Getenv("RENDER_HOST"),
		RendererPort:         RendererPort,
		ParserTimeoutSeconds: timeout,
	}

	log.Printf("config: %#v\n", config)
	return config, nil
}
