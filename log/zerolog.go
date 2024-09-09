package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	log.Debug().
		Int("StuID", 1001).
		Msg("Getting employee information")

	log.Debug().
		Str("Name", "John").
		Send()
	//log.Info().Msg("Hello, World!")
}
