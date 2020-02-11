package main

import (
	"time"

	"github.com/rs/zerolog/log"

	. "github.com/yoonhwan/go-example/datastructure"
	. "github.com/yoonhwan/go-example/datastructure/linkedlist"
)

func main() {
	logger()

	// for i := 0; i < 100; i++ {
		STContextMgr().MakeContext("TestOne")
	// }

	// STContextMgr().Clear("TestOne")


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

	/*
	b, cancel :=context.Context.NewContextWithCancel()
	go func(ctx context.Context, test int) {
	outer:
		for {
			select {
			case <- ctx.Done():
				break outer
			default:

				test++

				if test >= 100
					break outer
				break
			}
		}
	}(b, 0)
	cancel()

	func test1() {
		test := 1
		for i := 0; i < 100; i++ {
			test++	
		}
	}

	*/

}
