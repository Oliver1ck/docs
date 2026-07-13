package main

import (
	"github.com/rs/zerolog/log"

	"github.com/Oliver1ck/docs/internal/config"
)

func main() {
	_, err := config.LoadConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to load config")
	}
}
