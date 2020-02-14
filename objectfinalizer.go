package main

import (
	"runtime"
	"time"

	"github.com/rs/zerolog/log"
)

type person struct {
	Name string
	Age  int
}

// FinalizerTest : desc
func FinalizerTest() {
	log.Debug().Msg("start finalizerTest")
	go func() {
		p := &person{"Bob", 20}
		runtime.SetFinalizer(p, func(p2 *person) {
			log.Debug().Msgf("Finalizing %v", p2)
		})
		p = nil
		runtime.GC()
	}()

	time.Sleep(time.Second * 1)
	log.Debug().Msg("Done")
}
