package main

import (
	"context"
	"errors"
	"fmt"
	"runtime"
	"sync"
	"time"

	"github.com/rs/zerolog/log"
)

type contextDetailItem struct {
	ctx      context.Context
	cancel   context.CancelFunc
	name     string
	stime    time.Time
	duration time.Duration
	callFunc string
}

func (a *contextDetailItem) set(n string, st time.Time, ca string, dur time.Duration) *contextDetailItem {
	a.name = n
	a.stime = st
	a.duration = dur
	a.callFunc = ca

	return a
}

func makeDetail(c context.Context, cc context.CancelFunc) *contextDetailItem {
	return &contextDetailItem{
		ctx:    c,
		cancel: cc,
	}
}

// ContextMgr : context manager
type ContextMgr struct {
	sync.RWMutex
	contextMap map[string]*contextDetailItem
}

// Set : desc
func (a *ContextMgr) Set(name string, detail *contextDetailItem) {
	a.Lock()
	defer a.Unlock()
	if _, ok := a.contextMap[name]; ok {
		panic(errors.New("duplicate"))
	} else {
		a.contextMap[name] = detail
		a.cancelWatcher(detail)
	}
}

// Get : desc
func (a *ContextMgr) Get(name string) *contextDetailItem {
	a.Lock()
	defer a.Unlock()
	if a.contextMap == nil {
		return nil
	} else if val, ok := a.contextMap[name]; ok {
		return val
	}
	return nil
}

// Cancel : desc
func (a *ContextMgr) Cancel(name string) {
	a.Lock()
	defer a.Unlock()
	if detail, ok := a.contextMap[name]; ok {
		detail.cancel()
	}
}

// Finish : desc
func (a *ContextMgr) Finish(detail *contextDetailItem) {
	a.Lock()
	defer a.Unlock()
	if _, ok := a.contextMap[detail.name]; ok {
		log.Info().Msgf("\"%v\" context finish %v working duration with (%v)", detail.name, time.Now().Sub(detail.stime), detail.ctx.Err())
		a.contextMap[detail.name] = nil
		delete(a.contextMap, detail.name)
	}
}

// cancelWatcher : desc
func (a *ContextMgr) cancelWatcher(detail *contextDetailItem) {
	go func() {
		if detail != nil {
			for {
				select {
				case <-detail.ctx.Done():
					a.Finish(detail)
					break
				}
			}
		}
	}()
}

// Iter : desc
func (a *ContextMgr) Iter() <-chan *contextDetailItem {
	c := make(chan *contextDetailItem, len(a.contextMap))

	f := func() {
		a.Lock()
		defer a.Unlock()
		for _, detail := range a.contextMap {
			c <- detail
		}
		close(c)
	}
	go f()
	return c
}

// MakeDurationContext : new context
func (a *ContextMgr) MakeDurationContext(name string, dur time.Duration) (detail *contextDetailItem, err error) {
	_, fn, line, _ := runtime.Caller(1)
	err = nil
	detail = makeDetail(context.WithCancel(context.Background())).set(name, time.Now(), fmt.Sprintf("%v::%v", fn, line), dur)
	a.Set(name, detail)

	log.Info().Msgf("%v::%v, Make context for %v", fn, line, name)
	return
}

// MakeLoopContext : new context
func (a *ContextMgr) MakeLoopContext(name string) (detail *contextDetailItem, err error) {
	_, fn, line, _ := runtime.Caller(1)
	err = nil
	detail = makeDetail(context.WithCancel(context.Background())).set(name, time.Now(), fmt.Sprintf("%v::%v", fn, line), -1)
	a.Set(name, detail)

	log.Info().Msgf("%v::%v, Make context for %v", fn, line, name)
	return
}

// Watch : watch monitor context manager
func (a *ContextMgr) Watch() {
	detail, _ := a.MakeLoopContext("ContextMgr Watcher")
	go func(ctx context.Context) {

		for {
			iter := a.Iter()
			for sub := range iter {
				if sub.duration == -1 {
					log.Info().Msgf("context watcher alive loop context :: %v", sub.name)
				} else {
					if time.Now().Sub(detail.stime) > sub.duration {
						log.Error().
							Str("fix:", "make cancel operation like \"defer contextDetailItem.cancel()\" inside single context method").
							Msgf("context watcher alive context error :: %v", sub.name)
					}
				}

			}
			time.Sleep(time.Second * 4)
		}
	}(detail.ctx)
}

// StartContextMgr : desc
func (a *ContextMgr) StartContextMgr() *ContextMgr {
	a.Lock()
	defer a.Unlock()
	if a.contextMap == nil {
		a.contextMap = make(map[string]*contextDetailItem)
	}

	return a
}

// FinishContextMgr : desc
func (a *ContextMgr) FinishContextMgr() {
	go func() {
		iter := a.Iter()
		for sub := range iter {
			a.Cancel(sub.name)

		}
	}()
}

// ContextText : desc
func ContextText() {
	// new context manager samples
	if detail, err := STContextMgr().MakeDurationContext("TestOne", time.Second*10); err == nil {
		log.Info().Msgf(longFuncWithCtx(detail))
	}
	// STContextMgr().Cancel("TestOne")

	if detail, err := STContextMgr().MakeLoopContext("TestTwo"); err == nil {
		loopFuncWithCtx(detail)
		<-time.After(time.Second * 3)
		STContextMgr().Cancel("TestTwo")
	}

	if detail, err := STContextMgr().MakeLoopContext("TestThree"); err == nil {
		go func() {
			for {
				select {
				case <-detail.ctx.Done():
					log.Info().Msg("TestThree1")
					return
				}
			}
		}()
		go func() {
			for {
				select {
				case <-detail.ctx.Done():
					log.Info().Msg("TestThree2")
					return
				}
			}
		}()

		loopFuncWithCtx(detail)
		<-time.After(time.Second * 3)
		STContextMgr().Cancel("TestThree")

		<-time.After(time.Second * 10)
		STContextMgr().FinishContextMgr()
	}
}

func longFunc() string {
	<-time.After(time.Second * 3) // long running job
	return "Success"
}

func longFuncWithCtx(detail *contextDetailItem) (string, error) {
	done := make(chan string)
	// defer detail.cancel()
	//context leak sample

	go func() {
		done <- longFunc()
	}()

	select {
	case result := <-done:
		return result, nil
	case <-detail.ctx.Done():
		return "Fail", detail.ctx.Err()
	}
}

func loopFuncWithCtx(detail *contextDetailItem) {
	go func() {
	outer:
		for {
			select {
			case <-detail.ctx.Done():
				log.Info().Msg("quit loopFuncWithCtx 1")
				break outer
				//or return
			case <-time.After(time.Second * 1):
				log.Info().Msg("hi i'm alive")
			}
		}
	}()
}
