package main

import (
	"sync"
	"context"
	"time"

	"github.com/rs/zerolog/log"
)

func contextText() {

	
	var wg sync.WaitGroup
	
	wg.Add(1)
	go func() {
		
		ctx := context.Background()
		ctx, cancel := context.WithCancel(ctx)
		ctx, cancel = context.WithCancel(ctx)
		defer func() {
			cancel()
			wg.Done()
		}()
		go testA(&ctx)
		go testB(&ctx)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		outer:
		for {
			select {
				case <-time.After(time.Second * 6):
					log.Info().Msg("Waitting.")
					break outer
			}
		}
	}()
	wg.Wait()
	wg.Add(2)
	go func() {
	
		ctx1, cancel1 := context.WithCancel(context.Background())
		ctx2, _ := context.WithCancel(context.Background())
		defer func() {
			cancel1()
			wg.Done()
		}()
		defer func() {
			// cancel2()
			wg.Done()
		}()
		go testA(&ctx1)
		go testB(&ctx2)
	}()
	wg.Wait()
}

func testA(ctx *context.Context) {
	outer:
	for {
		select {
		case <-(*ctx).Done():
			log.Info().Msg("testA stop here")
			break outer
		case <-time.After(time.Second * 3):
			log.Info().Msg("testA")
			break
		}
	}
}

func testB(ctx *context.Context) {
	outer:
	for {
		select {
		case <-(*ctx).Done():
			log.Info().Msg("testB stop here")
			break outer
		case <-time.After(time.Second * 2):
			log.Info().Msg("testB")
			break
		}
	}
}