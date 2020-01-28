package main

import (
	"time"

	"github.com/rs/zerolog/log"
)

func main() {
	logger()
	test()
	go contextText()
	go test_signal()
	for {
		select {
		case <-time.After(time.Second * 1):
			log.Info().Msg("running..")
		}
	}
}
