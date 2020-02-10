package main

//
// import "github.com/cenkalti/backoff"
import (
	"context"
	"errors"
	"fmt"
	"time"

	. "github.com/cenkalti/backoff"
	"github.com/rs/zerolog/log"
)

func backOff() BackOff {
	result := NewExponentialBackOff()
	result.InitialInterval = 1 * time.Second
	result.MaxInterval = 10 * time.Second
	result.MaxElapsedTime = 1 * time.Minute
	result.Reset()
	return result
}

func ClosureOperator(name string) func() error {
	x := 0
	f_name := name
	return func() error {
		x++
		if x >= 6 {
			log.Info().Msgf("%s :: work Closure op", f_name)
			return nil
		}
		return errors.New(fmt.Sprintf("error %v", f_name))
	}
}

func BackoffTest() {

	b := NewConstantBackOff(time.Millisecond)
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	cb := WithContext(b, ctx)

	if cb.Context() != ctx {
		log.Error().Msg("invalid context")
	}

	cancel()

	if cb.NextBackOff() != Stop {
		log.Error().Msg("invalid next back off")
	}

	// normal function backoff
	op := ClosureOperator("test1")
	notify := func(err error, t time.Duration) {
		log.Warn().Msgf("error: %v happened at time: %v", err, t)
	}

	eb := backOff()
	err := RetryNotify(op, eb, notify)
	// err := Retry(operation(), eb)
	if err != nil {
		// Handle error.
		log.Error().Msgf("error: %v happend", err)
		return
	}

	//customize backoff iteration
	customEb, _ := backOff().(*ExponentialBackOff)
	customEb.MaxElapsedTime = 3 * time.Second
	eb = WithContext(customEb, context.Background())
	op = ClosureOperator("test2")

	start := time.Now()
	for {
		//itrator
		error1 := op()

		if error1 == nil {
			break
		}

		notify(error1, time.Now().Sub(start))

		next := customEb.NextBackOff()
		if next == Stop {
			log.Error().Msg("error: time out woker")
			break
		}
		time.Sleep(next)
	}

	log.Info().Msg("Backoff test finish")
}
