package main

import (
	"time"

	"github.com/rs/zerolog/log"

	. "github.com/yoonhwan/go-example/datastructure"
	. "github.com/yoonhwan/go-example/datastructure/linkedlist"
)

func main() {
	logger()

	STContextMgr().MakeContext("TestOne")

	Singly()
	Doubly()
	test()

	StartStackTest()
	StartQueueTest()

	BackoffTest()

	go contextText()
	go test_signal()
	for {
		select {
		case <-time.After(time.Second * 1):
			log.Info().Msg("running..")
		}
	}

}
