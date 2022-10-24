package nslice

import (
	"github.com/foxclore/go-utils/everything"
	"sync"
)

// This package implements a nice-slice, with some nice built-in methods

type NSlice[underlying comparable] struct {
	S      []underlying
	synced bool
	mutex  *sync.Mutex
	mapper map[underlying]struct{}
}

func NiceSlice[underlying comparable](items []underlying) *NSlice[underlying] {
	n := &NSlice[underlying]{
		S:      items,
		synced: true,
		mutex:  &sync.Mutex{},
	}
	n.mutex.Lock()
	n.mapper = make(map[underlying]struct{})
	n.mutex.Unlock()
	n.Sync()
	return n
}

func (n *NSlice[underlying]) Sync() {
	n.mutex.Lock()
	n.mapper = make(map[underlying]struct{})
	for _, x := range n.S {
		n.mapper[x] = everything.Empty
	}
	n.synced = true
	n.mutex.Unlock()
}

func (n *NSlice[underlying]) Append(items ...underlying) *NSlice[underlying] {
	return NiceSlice(append(n.S, items...))
}

func (n *NSlice[underlying]) AppendAssign(items ...underlying) {
	n.S = append(n.S, items...)
	n.Sync()
}

func (n *NSlice[underlying]) Contains(item underlying) bool {
	n.mutex.Lock()
	_, ok := n.mapper[item]
	n.mutex.Unlock()
	return ok
}
