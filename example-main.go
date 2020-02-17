package main

import (
	"time"

	"github.com/rs/zerolog/log"
	// . "github.com/yoonhwan/go-example/datastructure"
	// . "github.com/yoonhwan/go-example/datastructure/linkedlist"
)

func main() {
	ch := make(chan int)
	go func() {
		detail, _ := STContextMgr().MakeLoopContext("main watcher")
		for {
			select {
			case <-time.After(time.Second * 1):
				// log.Info().Msg("r....")
			case <-detail.ctx.Done():
				// for context test
				// ch <- 1
				// return
			}
		}

	}()
	go test_signal()

	// slack test
	SlackTest()

	// golang default system test
	// loggerTest()
	FinalizerTest()

	// datastructure test
	// Singly()
	// Doubly()
	// StartStackTest()
	// StartQueueTest()

	// context & backoff test
	// BackoffTest()
	ContextText()

	<-ch

	log.Info().Msg("bye bye..")
}
