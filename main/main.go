package main

import (
	"os"

	"fiber-service/composition"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func init() {
	// init logger
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: "2006-01-02 15:04:05", NoColor: true}).Level(zerolog.InfoLevel)
}

func main() {
	log.Info().Msg("Starting..")

	application, err := composition.Application()
	if err != nil {
		log.Error().Err(err).Msg("Composition")
	}
	err = application.Start()
	if err != nil {
		log.Error().Err(err).Msg("Application")
	}
}
