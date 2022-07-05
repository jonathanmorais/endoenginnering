package main

import (
	"github.com/jonathanmorais/endoenginnering/rec-service/pkg/logger"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"net"
	"os"
)

func main() {

	logInternal := logger.Instance()

	zerolog.SetGlobalLevel(zerolog.DebugLevel)

	log.Logger = log.Output(
		zerolog.ConsoleWriter{
			Out:     os.Stderr,
			NoColor: false,
		},
	)

	_, err := net.Listen("tcp", ":30000")
	if err != nil {
		log.Error().
			Str("Error", err.Error()).
			Msg("Failed to listen")
	}

	logInternal.Info().
		Msg("Listener for carbohidra-service is created")

}
