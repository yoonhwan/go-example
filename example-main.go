package main

import (
	"time"

	"github.com/rs/zerolog/log"

	. "github.com/yoonhwan/go-example/datastructure/linkedlist"
	. "github.com/yoonhwan/go-example/datastructure"
)

func main() {
	logger()

	Singly()
	Doubly()
	test()
	
	// StartStackTest()

	go contextText()
	go test_signal()
	for {
		select {
		case <-time.After(time.Second * 1):
			log.Info().Msg("running..")
		}
	}

}
