package main

import "sync"

type singletonMgr struct {
	context *ContextMgr

}

var (
	once   sync.Once
	single *singletonMgr = &singletonMgr{}
)

//STContextMgr : get singleton context manager
func STContextMgr() *ContextMgr {
	once.Do(func() {
		single.context = new(ContextMgr)
	})

	return single.context
}

func init() {

}
