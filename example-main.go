package main

import (
	"flag"

	"github.com/rs/zerolog/log"

	"time"
	// . "github.com/yoonhwan/go-example/datastructure"
	// . "github.com/yoonhwan/go-example/datastructure/linkedlist"
)

var (
	env = flag.String("env", "dev1", "the Env")
)

// 1. 사용할 config type을 정한 뒤 config.go의 import 주소를 변경한다
// _ "github.com/crgimenes/goconfig/ini" <- json 타입 사용시 ini를 json으로 변경ㅎ나다
// 2. flag로 개발환경 변수를 받아서 config file을 로드 한다
// config.dev1.ini 형식으로 사용중
// 3. go run . -env=dev1

func main() {

	flag.Parse() // 명령줄 옵션의 내용을 각 자료형별로 분석

	config := confgTestFunc(*env)

	log.Info().Msg(config.ServerName)
	log.Info().Msg(config.Sqs.SendQueue)

	ch := make(chan int)
	go func() {
		detail, _ := STContextMgr().MakeLoopContext("main watcher")
		for {
			select {
			case <-time.After(time.Second * 1):
				// log.Info().Msg("r....")
			case <-detail.ctx.Done():
				// for context test
				ch <- 1
				return
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

	pubsubTest()
}
