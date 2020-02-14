package main

// import "github.com/rs/zerolog"
import (
	// "errors"
	"flag"
	"fmt"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func init() {
	//https://github.com/rs/zerolog

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	//Pretty logging
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	log.Logger = log.With().Caller().Logger()

}

func loggerTest() {
	fmt.Println("logger")

	log.Print("hello world")

	//Contextual Logging
	log.Debug().
		Str("Scale", "833 cents").
		Float64("Interval", 833.09).
		Msg("Fibonacci is everywhere")

	log.Debug().
		Str("Name", "Tom").
		Send()
	//Leveled Logging
	log.Info().Msg("hello world")

	//Logging without Level or Message
	log.Log().
		Str("foo", "bar").
		Msg("")

	//Setting Global Log Level
	debug := flag.Bool("debug", false, "sets log level to debug")
	flag.Parse()

	// Default level for this example is info, unless debug flag is present
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if *debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	log.Debug().Msg("This message appears only when log level set to Debug")
	log.Info().Msg("This message appears when log level set to Debug or Info")

	if e := log.Debug(); e.Enabled() {
		// Compute log output only if enabled.
		value := "bar"
		e.Str("foo", value).Msg("some debug message")
	}

	//Logging Fatal Messages
	// err := errors.New("A repo man spends his life getting into tense situations")
	// service := "myservice"
	// log.Fatal().
	// 	Err(err).
	// 	Str("service", service).
	// 	Msgf("Cannot start %s", service)

}
