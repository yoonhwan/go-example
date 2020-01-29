package container

import (
	. "github.com/yoonhwan/go-example/datastructure/linkedlist"
)

type map_func func(int)
type iStack interface {
	push() *DNode
	pop() *DNode
	clear()
	flatmap(fn map_func)
}

type Stack struct {
	doubly_head *DNode
	doubly_tail *DNode
}

func StartStackTest() {
	// testStack := Stack{}
}
